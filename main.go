package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"Sinekod/jsonManager"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
	w.WriteHeader(http.StatusOK)
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	response := jsonManager.Get_json_id(idInt)
	w.Write(response)
	w.WriteHeader(http.StatusOK)

}

func GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	response, code := jsonManager.Get_json_books_id(idInt)
	if code == "200" {
		w.Write(response)
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func GetAllBooksHandler(w http.ResponseWriter, r *http.Request) {
	response := jsonManager.Get_json_books()
	w.Write(response)
	w.WriteHeader(http.StatusOK)
}

func PostUsers(w http.ResponseWriter, r *http.Request) {
	array, code := jsonManager.Post_json_users(r)
	if code == "201" {
		w.Write(array)
		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

}

func PostBooks(w http.ResponseWriter, r *http.Request) {
	
	array, code := jsonManager.Post_json_books(r)
	if code == "201" {
		w.Write(array)
		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler)                              // OK
	r.HandleFunc("/users/{id}", GetUsersHandler).Methods("GET") // OK
	r.HandleFunc("/users", PostUsers).Methods("POST")
	r.HandleFunc("/books", GetAllBooksHandler).Methods("GET") // OK
	r.HandleFunc("/books", PostBooks).Methods("POST")
	r.HandleFunc("/books/{id}", GetBooksHandler).Methods("GET") // OK

	fmt.Println("Server listening...")
	http.ListenAndServe(":8080", r)
}
