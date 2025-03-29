package repository

import (
	"errors"

	"github.com/chetana/conv-chet/internal/model"
)

// UserRepository est une interface pour accéder aux données des utilisateurs.
type UserRepository interface {
	GetUserByID(id string) (*model.User, error) // Retourne un pointeur vers un User et une erreur
	CreateUser(user *model.User) error
}

// userRepositoryImpl est une implémentation "en mémoire" de UserRepository.
type userRepositoryImpl struct {
	users map[string]*model.User // Stocke les utilisateurs dans une map (ID -> User)
}

// NewUserRepository crée une nouvelle instance de userRepositoryImpl.
func NewUserRepository() UserRepository {
	return &userRepositoryImpl{
		users: make(map[string]*model.User), // Initialise la map
	}
}

// GetUserByID récupère un utilisateur par son ID.
func (r *userRepositoryImpl) GetUserByID(id string) (*model.User, error) {
	user, ok := r.users[id] // Vérifie si l'utilisateur existe dans la map
	if !ok {
		return nil, errors.New("user not found") // Retourne une erreur si l'utilisateur n'existe pas
	}
	return user, nil // Retourne l'utilisateur et nil (pas d'erreur)
}

// CreateUser crée un nouvel utilisateur.
func (r *userRepositoryImpl) CreateUser(user *model.User) error {
	if _, ok := r.users[user.ID]; ok {
		return errors.New("user already exists") // Retourne une erreur si l'utilisateur existe déjà
	}
	r.users[user.ID] = user // Ajoute l'utilisateur à la map
	return nil              // Retourne nil (pas d'erreur)
}
