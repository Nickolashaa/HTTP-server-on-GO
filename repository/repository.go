package repository

import (
	"database/sql"

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
