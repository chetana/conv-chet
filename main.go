package main

import (
	"fmt"
	"log"
	"net/http"
	"os" // Import du package os
	"strconv"

	"github.com/chetana/conv-chet/internal/app"        // Correction ici
	"github.com/chetana/conv-chet/internal/controller" // Correction ici
	"github.com/chetana/conv-chet/internal/repository" // Correction ici
	"github.com/chetana/conv-chet/internal/service"    // Correction ici
	"github.com/gorilla/mux"
)

func main() {
	// Initialisation de Firestore
	app.InitializeFirestore()

	// Initialisation du repository
	todoRepository := repository.NewTodoRepository()

	// Initialisation du service
	todoService := service.NewTodoService(todoRepository)

	// Initialisation du controller
	todoController := controller.NewTodoController(todoService)

	// Création du routeur
	router := mux.NewRouter()

	// Définition des routes
	router.HandleFunc("/todos", todoController.GetAllTodos).Methods("GET")
	router.HandleFunc("/todos", todoController.CreateTodo).Methods("POST")
	router.HandleFunc("/todos", todoController.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todos", todoController.DeleteTodo).Methods("DELETE")
	router.HandleFunc("/todos", todoController.GetTodo).Methods("GET").Queries("id", "{id}")

	// Récupérer le port depuis la variable d'environnement (pour Cloud Run)
	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = "8080" // Valeur par défaut si la variable n'est pas définie
	}

	// Conversion du port en entier
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Erreur lors de la conversion du port en entier: %v\n", err)
	}

	// Démarrage du serveur
	fmt.Printf("Server listening on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
