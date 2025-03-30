package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"Sinekod/models"

	"Sinekod/storage"

	"Sinekod/jsonManager"
)

// Старовая страница
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

// GET для пользователей
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	response := jsonManager.Get_json_id(id)
	w.Write(response)

}

func GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	response, code := jsonManager.Get_json_books_id(id)
	if code == "200" {
		w.Write(response)
	} else {
		fmt.Fprintf(w, code)
	}
}

func main() {
	storage.Books["0"] = models.Book{Title: "Первая книга"}
	storage.Books["1"] = models.Book{Title: "Вторая книга"}
	storage.Books["2"] = models.Book{Title: "Третья книга"}
	storage.Users["0"] = models.User{Name: "Николай", Email: "gracevnikolaj220@gmail.com"}
	storage.Users["1"] = models.User{Name: "Тимофей", Email: "TIMAK435@gmail.com"}

	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler)                              // OK
	r.HandleFunc("/users/{id}", GetUsersHandler).Methods("GET") // OK
	r.HandleFunc("/users", GetUsersHandler).Methods("POST")
	r.HandleFunc("/books", GetUsersHandler).Methods("GET")
	// r.HandleFunc("/books", GetUsersHandler).Methods("POST")
	r.HandleFunc("/books/{id}", GetBooksHandler).Methods("GET") // OK

	fmt.Println("Server listening...")
	http.ListenAndServe(":8080", r)

}
