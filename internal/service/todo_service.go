package service

import (
	"github.com/chetana/conv-chet/internal/model"
	"github.com/chetana/conv-chet/internal/repository"
)

// TodoService est une interface pour la logique métier des tâches.
type TodoService interface {
	GetTodoByID(id string) (*model.Todo, error)
	GetAllTodos() ([]*model.Todo, error)
	CreateTodo(todo *model.Todo) error
	UpdateTodo(todo *model.Todo) error
	DeleteTodo(id string) error
}

// todoServiceImpl est une implémentation de TodoService.
type todoServiceImpl struct {
	todoRepository repository.TodoRepository // Dépend du TodoRepository
}

// NewTodoService crée une nouvelle instance de todoServiceImpl.
func NewTodoService(todoRepository repository.TodoRepository) TodoService {
	return &todoServiceImpl{todoRepository: todoRepository}
}

// GetTodoByID récupère une tâche par son ID.
func (s *todoServiceImpl) GetTodoByID(id string) (*model.Todo, error) {
	return s.todoRepository.GetTodoByID(id) // Délègue l'appel au TodoRepository
}

// GetAllTodos récupère toutes les tâches.
func (s *todoServiceImpl) GetAllTodos() ([]*model.Todo, error) {
	return s.todoRepository.GetAllTodos() // Délègue l'appel au TodoRepository
}

// CreateTodo crée une nouvelle tâche.
func (s *todoServiceImpl) CreateTodo(todo *model.Todo) error {
	// Ici, tu peux ajouter de la logique métier (validation, etc.) avant de créer la tâche
	return s.todoRepository.CreateTodo(todo) // Délègue l'appel au TodoRepository
}

// UpdateTodo met à jour une tâche existante.
func (s *todoServiceImpl) UpdateTodo(todo *model.Todo) error {
	// Ici, tu peux ajouter de la logique métier (validation, etc.) avant de mettre à jour la tâche
	return s.todoRepository.UpdateTodo(todo) // Délègue l'appel au TodoRepository
}

// DeleteTodo supprime une tâche par son ID.
func (s *todoServiceImpl) DeleteTodo(id string) error {
	// Ici, tu peux ajouter de la logique métier (vérification des autorisations, etc.) avant de supprimer la tâche
	return s.todoRepository.DeleteTodo(id) // Délègue l'appel au TodoRepository
}
