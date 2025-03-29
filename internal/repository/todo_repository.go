package repository

import (
	"context"
	"log"

	"github.com/chetana/conv-chet/internal/app"
	"github.com/chetana/conv-chet/internal/model"
)

const collectionName = "todos"

// TodoRepository est une interface pour accéder aux données des tâches.
type TodoRepository interface {
	GetTodoByID(id string) (*model.Todo, error)
	GetAllTodos() ([]*model.Todo, error)
	CreateTodo(todo *model.Todo) error
	UpdateTodo(todo *model.Todo) error
	DeleteTodo(id string) error
}

// todoRepositoryImpl est une implémentation Firestore de TodoRepository.
type todoRepositoryImpl struct{}

// NewTodoRepository crée une nouvelle instance de todoRepositoryImpl.
func NewTodoRepository() TodoRepository {
	return &todoRepositoryImpl{}
}

// GetTodoByID récupère une tâche par son ID.
func (r *todoRepositoryImpl) GetTodoByID(id string) (*model.Todo, error) {
	ctx := context.Background()
	doc, err := app.FirestoreClient.Collection(collectionName).Doc(id).Get(ctx)
	if err != nil {
		log.Printf("Failed to get todo: %v", err)
		return nil, err
	}
	var todo model.Todo
	err = doc.DataTo(&todo)
	if err != nil {
		log.Printf("Failed to convert data to todo: %v", err)
		return nil, err
	}
	todo.ID = doc.Ref.ID // Récupérer l'ID du document
	return &todo, nil
}

// GetAllTodos récupère toutes les tâches.
func (r *todoRepositoryImpl) GetAllTodos() ([]*model.Todo, error) {
	ctx := context.Background()
	iter := app.FirestoreClient.Collection(collectionName).Documents(ctx)
	defer iter.Stop()

	var todos []*model.Todo
	for {
		doc, err := iter.Next()
		if err != nil {
			log.Printf("Failed to iterate: %v", err)
			break
		}
		if doc == nil {
			break
		}

		var todo model.Todo
		err = doc.DataTo(&todo)
		if err != nil {
			log.Printf("Failed to convert data to todo: %v", err)
			continue
		}

		todo.ID = doc.Ref.ID // Récupérer l'ID du document
		todos = append(todos, &todo)
	}
	return todos, nil
}

// CreateTodo crée une nouvelle tâche.
func (r *todoRepositoryImpl) CreateTodo(todo *model.Todo) error {
	ctx := context.Background()
	_, err := app.FirestoreClient.Collection(collectionName).Doc(todo.ID).Set(ctx, todo)
	if err != nil {
		log.Printf("Failed to create todo: %v", err)
		return err
	}
	return nil
}

// UpdateTodo met à jour une tâche existante.
func (r *todoRepositoryImpl) UpdateTodo(todo *model.Todo) error {
	ctx := context.Background()
	_, err := app.FirestoreClient.Collection(collectionName).Doc(todo.ID).Set(ctx, todo)
	if err != nil {
		log.Printf("Failed to update todo: %v", err)
		return err
	}
	return nil
}

// DeleteTodo supprime une tâche par son ID.
func (r *todoRepositoryImpl) DeleteTodo(id string) error {
	ctx := context.Background()
	_, err := app.FirestoreClient.Collection(collectionName).Doc(id).Delete(ctx)
	if err != nil {
		log.Printf("Failed to delete todo: %v", err)
		return err
	}
	return nil
}
