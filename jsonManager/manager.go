package jsonManager

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "modernc.org/sqlite"
)

func Get_json_id(id int) []byte { //любой вывод json id
	var choto = map[string]int{"id": id}
	data, err := json.Marshal(choto)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

func Get_json_books() []byte { //GET /books
	var temp []struct {
		Id    int
		Title string
	}
	db, err := sql.Open("sqlite", "db/database.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM Books")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var b struct {
			Id    int
			Title string
		}
		err := rows.Scan(&b.Id, &b.Title)
		if err != nil {
			fmt.Println(err)
			continue
		}
		temp = append(temp, b)
	}
	data, err := json.Marshal(temp)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

func Get_json_books_id(id int) ([]byte, string) { //GET /books/{id}
	db, err := sql.Open("sqlite", "db/database.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM Books")
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
			data, err := json.Marshal(temp)
			if err != nil {
				fmt.Println(err)
			}
			return data, "200"
		}
	}
	return nil, "404"
}

func Post_json_users(r *http.Request) ([]byte, string) {
	var new_user struct {
		Name  string
		Email string
	}
	err := json.NewDecoder(r.Body).Decode(&new_user)
	if err != nil {
		return nil, "400"
	}
	s := new_user.Email
	flag := false
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "@" {
			flag = true
			break
		}
	}
	if !flag {
		return nil, "400"
	}

	db, err := sql.Open("sqlite", "db/database.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	result, err := db.Exec("INSERT INTO Users (Name, Email) VALUES ($1, $2)", new_user.Name, new_user.Email)
	if err != nil {
		panic(err)
	}

	id, _ := result.LastInsertId()

	return Get_json_id(int(id)), "201"

}

func Post_json_books(r *http.Request) ([]byte, string) {
	var temp struct {
		Title string
	}
	err := json.NewDecoder(r.Body).Decode(&temp)
	if err != nil {
		return nil, "400"
	}

	db, err := sql.Open("sqlite", "db/database.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	result, err := db.Exec("INSERT INTO Books (Title) VALUES ($1)", temp.Title)
	if err != nil {
		panic(err)
	}

	id, _ := result.LastInsertId()

	return Get_json_id(int(id)), "201"
}
