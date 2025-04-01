package jsonManager

import (
	"Sinekod/models"
	"Sinekod/storage"
	"encoding/json"
	"fmt"
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
	data, err := json.Marshal(storage.Books)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

func Get_json_books_id(id int) ([]byte, string) { //GET /books/{id}
	value, ok := storage.Books[id]
	if !ok {
		return nil, "404"
	}
	data, err := json.Marshal(value)
	if err != nil {
		fmt.Println(err)
	}
	return data, "200"
}

func Post_json_users(id int, new_user models.User) ([]byte, string) {
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
	storage.Users[id] = new_user
	return Get_json_id(id), "201"
}

func Post_json_books(id int, new_book models.Book) ([]byte, string) {
	storage.Books[id] = new_book
	return Get_json_id(id), "201"
}
