package model

type Todo struct {
	ID        string `json:"id"`        // Identifiant unique de la tâche
	Title     string `json:"title"`     // Titre de la tâche
	Completed bool   `json:"completed"` // Indique si la tâche est terminée ou non
}
