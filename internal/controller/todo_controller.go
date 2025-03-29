package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/chetana/conv-chet/internal/model"
	"github.com/chetana/conv-chet/internal/service"
	"github.com/google/uuid" // Pour générer des UUIDs
)

// TodoController gère les requêtes HTTP liées aux tâches.
type TodoController struct {
	todoService service.TodoService // Dépend du TodoService
}

// NewTodoController crée une nouvelle instance de TodoController.
func NewTodoController(todoService service.TodoService) *TodoController {
	return &TodoController{todoService: todoService}
}

// GetTodo récupère une tâche par son ID.
func (c *TodoController) GetTodo(w http.ResponseWriter, r *http.Request) {
	// Récupérer l'ID de la tâche depuis les paramètres de la requête (à implémenter)
	todoID := r.URL.Query().Get("id") // Exemple : /todos?id=123

	if todoID == "" {
		http.Error(w, "Missing todo ID", http.StatusBadRequest)
		return
	}

	todo, err := c.todoService.GetTodoByID(todoID) // Appelle le TodoService
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting todo: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	if todo == nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	// Sérialiser la tâche en JSON et l'envoyer dans la réponse
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

// GetAllTodos récupère toutes les tâches.
func (c *TodoController) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := c.todoService.GetAllTodos() // Appelle le TodoService
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting all todos: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	// Sérialiser les tâches en JSON et les envoyer dans la réponse
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

// CreateTodo crée une nouvelle tâche.
func (c *TodoController) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo model.Todo
	err := json.NewDecoder(r.Body).Decode(&todo) // Décode le corps de la requête JSON
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding request body: %s", err.Error()), http.StatusBadRequest)
		return
	}

	// Générer un ID unique pour la tâche
	todo.ID = uuid.New().String()

	err = c.todoService.CreateTodo(&todo) // Appelle le TodoService
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating todo: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated) // Envoie un code de statut 201 (Created)
	fmt.Fprintf(w, "Todo created successfully")
}

// UpdateTodo met à jour une tâche existante.
func (c *TodoController) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	var todo model.Todo
	err := json.NewDecoder(r.Body).Decode(&todo) // Décode le corps de la requête JSON
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding request body: %s", err.Error()), http.StatusBadRequest)
		return
	}

	err = c.todoService.UpdateTodo(&todo) // Appelle le TodoService
	if err != nil {
		http.Error(w, fmt.Sprintf("Error updating todo: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK) // Envoie un code de statut 200 (OK)
	fmt.Fprintf(w, "Todo updated successfully")
}

// DeleteTodo supprime une tâche existante.
func (c *TodoController) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	// Récupérer l'ID de la tâche depuis les paramètres de la requête (à implémenter)
	todoID := r.URL.Query().Get("id") // Exemple : /todos?id=123

	if todoID == "" {
		http.Error(w, "Missing todo ID", http.StatusBadRequest)
		return
	}

	err := c.todoService.DeleteTodo(todoID) // Appelle le TodoService
	if err != nil {
		http.Error(w, fmt.Sprintf("Error deleting todo: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK) // Envoie un code de statut 200 (OK)
	fmt.Fprintf(w, "Todo deleted successfully")
}
