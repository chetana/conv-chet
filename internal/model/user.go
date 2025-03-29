package model

type User struct {
	ID    string `json:"id"`    // `json:"id"` indique le nom du champ JSON
	Name  string `json:"name"`  // pour la sérialisation/désérialisation
	Email string `json:"email"`
}