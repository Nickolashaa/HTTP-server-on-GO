package jsonManager

import (
	"Sinekod/models"
	"Sinekod/storage"
	"encoding/json"
	"fmt"
)

func Get_json_id(id string) []byte { //любой вывод json id
	var choto = map[string]string{"id": id}
	data, err := json.Marshal(choto)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

func Get_json_books() []byte { //GET /books
	data, err := json.Marshal(storage.Books)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

func Get_json_books_id(id string) ([]byte, string) { //GET /books/{id}
	value, ok := storage.Books[id]
	if !ok {
		return nil, "404 not found"
	}
	data, err := json.Marshal(value)
	if err != nil {
		fmt.Println(err)
	}
	return data, "200"
}

func Post_json_users(id string, data []byte) ([]byte, string) {
	new_user := models.User{}
	err := json.Unmarshal(data, &new_user)
	if err != nil {
		fmt.Println(err)
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
		fmt.Println("Некорректный email")
	}
	storage.Users[id] = new_user
	return Get_json_id(id), "201 Created"
}

func Post_json_books(id string, data []byte) ([]byte, string) {
	new_book := models.Book{}
	err := json.Unmarshal(data, &new_book)
	if err != nil {
		fmt.Println(err)
	}
	storage.Books[id] = new_book
	return Get_json_id(id), "201 Created"
}
