package controller

import (
	"Sinekod/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Controller struct {
	Service *service.Service
}

func NewController(service *service.Service) *Controller {
	return &Controller{
		Service: service,
	}
}

func (c Controller) HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
	w.WriteHeader(http.StatusOK)
}

func (c Controller) PostUsers(w http.ResponseWriter, r *http.Request) {
	array, code := c.Service.Post_json_users(r)
	if code == "201" {
		w.Write(array)
		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (c Controller) PostBooks(w http.ResponseWriter, r *http.Request) {
	array, code := c.Service.Post_json_books(r)
	if code == "201" {
		w.Write(array)
		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (c Controller) GetUsersId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	response := c.Service.Get_json_id(idInt)
	w.Write(response)
	w.WriteHeader(http.StatusOK)
}

func (c Controller) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	response := c.Service.GetAllBooks()
	w.Write(response)
	w.WriteHeader(http.StatusOK)
}

func (c Controller) GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	response, code := c.Service.GetBookById(idInt)
	if code == "200" {
		w.Write(response)
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func (c Controller) DeleteBookId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	code := c.Service.DeleteBookId(idInt)
	if code == "200" {
		fmt.Fprintf(w, "Element was successfully deleted")
	} else {
		fmt.Fprintf(w, "There is no element with this id")
	}
}
