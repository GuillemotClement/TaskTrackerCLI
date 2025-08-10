package task

// on définis le nom de package selon le dossier dans lequel le fichier se trouve.
// le package main est utiliser uniquement pour le point d'entrée de l'application -> main.go

// importation du package pour le time
import (
	"errors"
	"strings"
	"time"
)

// déclare une var globale pour le message d'erreur
var ErrEmptyDescription = errors.New("description cannot be empty")

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

// on utilise un pointeur car on veux modifer par la suite et ne pas copier l'intégralité.
// amélioration des perf car on fait référence à un emplacement mémoire, et pas une copie de la struct
func NewTask(description string) (*Task, error) {
	//TrimSpace est plus adapter car il gère automatiquement les espaces et les tabs
	desc := strings.TrimSpace(description)
	// on test que la description n'est pas vide
	if len(desc) == 0 {
		// on retourne une task vide et le message d'erreur déclarer dans la variable globale
		// on peut retourner une valeur vide car on utilise un pointeur, on empêche ainsi la création de Task nul
		return nil, ErrEmptyDescription
	}
	// on déclare la date qui sera utiliser dans la struct pour les deux champs
	// sinon on peut potentiellement avoir des valeurs différentes dans la struct pour les champs de date
	now := time.Now()

	// par convention, on utilise un autre nom de variable que la struct
	// on fait référence au pointeur, permet de manipuler l'originale
	t := &Task{
		ID:          0,
		Description: desc,
		Status:      "todo",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	return t, nil
}

// on utilsie le receiver pour permettre d'accéder à la Task et ses fileds.
func (t *Task) UpdateDescription(description string) error {
	desc := strings.TrimSpace(description)
	// on fait le check de la longeur après le trim
	if len(desc) == 0 {
		return ErrEmptyDescription
	}

	// on accéde à la struct et on modifie les champs.
	t.Description = desc
	t.UpdatedAt = time.Now()
	return nil
}

func (t *Task) MarkIsDone() error {
	t.Status = "done"
	t.UpdatedAt = time.Now()
	return nil
}

func (t *Task) MarkIsProgress() error {
	t.Status = "in_progress"
	t.UpdatedAt = time.Now()
	return nil
}
