package task

// on définis le nom de package selon le dossier dans lequel le fichier se trouve.
// le package main est utiliser uniquement pour le point d'entrée de l'application -> main.go

// importation du package pour le time
import "time"

// définition de la struct de la task
// utiliser Time.time pour les champs de type date pour garder les méthodes lié aux dates
// par convention, ID est en majuscule
type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
