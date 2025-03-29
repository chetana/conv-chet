package repository

import (
	"errors"

	"github.com/chetana/conv-chet/internal/model"
)

// TodoRepository est une interface pour accéder aux données des tâches.
type TodoRepository interface {
	GetTodoByID(id string) (*model.Todo, error)
	GetAllTodos() ([]*model.Todo, error)
	CreateTodo(todo *model.Todo) error
	UpdateTodo(todo *model.Todo) error
	DeleteTodo(id string) error
}

// todoRepositoryImpl est une implémentation "en mémoire" de TodoRepository.
type todoRepositoryImpl struct {
	todos map[string]*model.Todo // Stocke les tâches dans une map (ID -> Todo)
}

// NewTodoRepository crée une nouvelle instance de todoRepositoryImpl.
func NewTodoRepository() TodoRepository {
	return &todoRepositoryImpl{
		todos: make(map[string]*model.Todo), // Initialise la map
	}
}

// GetTodoByID récupère une tâche par son ID.
func (r *todoRepositoryImpl) GetTodoByID(id string) (*model.Todo, error) {
	todo, ok := r.todos[id] // Vérifie si la tâche existe dans la map
	if !ok {
		return nil, errors.New("todo not found") // Retourne une erreur si la tâche n'existe pas
	}
	return todo, nil // Retourne la tâche et nil (pas d'erreur)
}

// GetAllTodos récupère toutes les tâches.
func (r *todoRepositoryImpl) GetAllTodos() ([]*model.Todo, error) {
	todos := []*model.Todo{} // Crée une slice vide pour stocker les tâches
	for _, todo := range r.todos {
		todos = append(todos, todo) // Ajoute chaque tâche à la slice
	}
	return todos, nil // Retourne la slice de tâches et nil (pas d'erreur)
}

// CreateTodo crée une nouvelle tâche.
func (r *todoRepositoryImpl) CreateTodo(todo *model.Todo) error {
	if _, ok := r.todos[todo.ID]; ok {
		return errors.New("todo already exists") // Retourne une erreur si la tâche existe déjà
	}
	r.todos[todo.ID] = todo // Ajoute la tâche à la map
	return nil              // Retourne nil (pas d'erreur)
}

// UpdateTodo met à jour une tâche existante.
func (r *todoRepositoryImpl) UpdateTodo(todo *model.Todo) error {
	if _, ok := r.todos[todo.ID]; !ok {
		return errors.New("todo not found") // Retourne une erreur si la tâche n'existe pas
	}
	r.todos[todo.ID] = todo // Met à jour la tâche dans la map
	return nil              // Retourne nil (pas d'erreur)
}

// DeleteTodo supprime une tâche par son ID.
func (r *todoRepositoryImpl) DeleteTodo(id string) error {
	if _, ok := r.todos[id]; !ok {
		return errors.New("todo not found") // Retourne une erreur si la tâche n'existe pas
	}
	delete(r.todos, id) // Supprime la tâche de la map
	return nil          // Retourne nil (pas d'erreur)
}
