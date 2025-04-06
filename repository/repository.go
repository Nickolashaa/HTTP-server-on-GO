package repository

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func NewDB() *sql.DB {
	db, err := sql.Open("sqlite", "db/database.db")
	if err != nil {
		panic(err)
	}
	return db
}

func (r Repository) Post_users(new_user struct {
	Name  string
	Email string
}) int {
	result, err := r.DB.Exec("INSERT INTO Users (Name, Email) VALUES ($1, $2)", new_user.Name, new_user.Email)
	if err != nil {
		panic(err)
	}
	id, _ := result.LastInsertId()

	return int(id)
}

func (r Repository) Post_json_books(temp struct {
	Title string
}) int {
	result, err := r.DB.Exec("INSERT INTO Books (Title) VALUES ($1)", temp.Title)
	if err != nil {
		panic(err)
	}
	id, _ := result.LastInsertId()
	return int(id)
}

func (r Repository) GetAllBooks() []struct {
	Id    int
	Title string
} {
	rows, err := r.DB.Query("select * from Books")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var Temp struct {
		Id    int
		Title string
	}

	var result []struct {
		Id    int
		Title string
	}

	for rows.Next() {
		rows.Scan(&Temp.Id, &Temp.Title)
		result = append(result, Temp)
	}

	return result
}

func (r Repository) GetBookById(id int) struct {
	Id    int
	Title string
} {
	rows, err := r.DB.Query("SELECT * from Books")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var temp struct {
		Id    int
		Title string
	}
	for rows.Next() {
		err := rows.Scan(&temp.Id, &temp.Title)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if temp.Id == id {
			return temp
		}
	}
	return struct {
		Id    int
		Title string
	}{Id: -1, Title: "None"}
}
