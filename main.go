package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"go.uber.org/dig"

	"Sinekod/controller"
	"Sinekod/repository"
	"Sinekod/service"
)

func main() {
	container := dig.New()

	_ = container.Provide(controller.NewController)
	_ = container.Provide(service.NewService)
	_ = container.Provide(repository.NewRepository)
	_ = container.Provide(repository.NewDB)

	container.Invoke(func(controller *controller.Controller) {
		r := mux.NewRouter()

		r.HandleFunc("/", controller.HomeHandler)
		r.HandleFunc("/users/{id}", controller.GetUsersId).Methods("GET")
		r.HandleFunc("/users", controller.PostUsers).Methods("POST")
		r.HandleFunc("/books", controller.GetAllBooks).Methods("GET")
		r.HandleFunc("/books", controller.PostBooks).Methods("POST")
		r.HandleFunc("/books/{id}", controller.GetBookById).Methods("GET")

		fmt.Println("Server listening...")
		http.ListenAndServe(":8080", r)
	})
}
