package controller

import(
	"Sinekod/service"
	"net/http"
	"fmt"
)


type Controller struct{
	Service *service.Service
}

func NewController (service *service.Service) *Controller{
	return &Controller{
		Service: service,
	}
}

func (c Controller) HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
	w.WriteHeader(http.StatusOK)
}

func (c Controller) PostUsers(w http.ResponseWriter, r *http.Request){
	array, code := c.Service.Post_json_users(r)
	if code == "201" {
		w.Write(array)
		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (c Controller) PostBooks(w http.ResponseWriter, r *http.Request){
	array, code := c.Service.Post_json_books(r)
	if code == "201" {
		w.Write(array)
		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}