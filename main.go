package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chetana/conv-chet/internal/controller"
	"github.com/chetana/conv-chet/internal/repository"
	"github.com/chetana/conv-chet/internal/service"

	"github.com/gorilla/mux"
)

func main() {
	// Initialisation du repository
	todoRepository := repository.NewTodoRepository()

	// Initialisation du service
	todoService := service.NewTodoService(todoRepository)

	// Initialisation du controller
	todoController := controller.NewTodoController(todoService)

	// Création du routeur
	router := mux.NewRouter()

	// Définition des routes
	router.HandleFunc("/todos", todoController.GetAllTodos).Methods("GET")                   // GET /todos
	router.HandleFunc("/todos", todoController.CreateTodo).Methods("POST")                   // POST /todos
	router.HandleFunc("/todos", todoController.UpdateTodo).Methods("PUT")                    // PUT /todos
	router.HandleFunc("/todos", todoController.DeleteTodo).Methods("DELETE")                 // DELETE /todos
	router.HandleFunc("/todos", todoController.GetTodo).Methods("GET").Queries("id", "{id}") // GET /todos?id=123

	// Démarrage du serveur
	port := 8080
	fmt.Printf("Server listening on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
