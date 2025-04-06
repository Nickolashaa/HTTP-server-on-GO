package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"Sinekod/repository"
)

type Service struct {
	repository *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		repository: repo,
	}
}

func (srv Service) Get_json_id(id int) []byte { //любой вывод json id
	var choto = map[string]int{"id": id}
	data, err := json.Marshal(choto)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

func (srv Service) Post_json_users(r *http.Request) ([]byte, string) {
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

	id := srv.repository.Post_users(new_user)

	return srv.Get_json_id(id), "201"
}

func (srv Service) Post_json_books(r *http.Request) ([]byte, string) {
	var temp struct {
		Title string
	}
	err := json.NewDecoder(r.Body).Decode(&temp)
	if err != nil {
		return nil, "400"
	}

	id := srv.repository.Post_json_books(temp)

	return srv.Get_json_id(id), "201"
}

func (srv Service) GetAllBooks() []byte {
	data := srv.repository.GetAllBooks()
	array, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return array
}

func (srv Service) GetBookById(id int) ([]byte, string) {
	data := srv.repository.GetBookById(id)
	if data.Id == -1 {
		return nil, "404"
	}
	array, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return array, "200"
}
