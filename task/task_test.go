package task

import (
	"errors"
	"testing"
)

func TestCreateTask(t *testing.T) {
	t.Run("should create a new task", func(t *testing.T) {
		// on ne mock pas la struct complete car les champs de date sont dynamique
		// on viens tester uniquement que la valeur est bien set (update et created)
		// on test que les valeurs pouvant être définis (status et description)
		// id est générer lors du stockage, donc on check si il est égale à 0 uniquement

		got, err := NewTask("Ma premiere task")

		// test qu'on récupère pas une erreur
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got.ID != 0 {
			t.Errorf("expected ID : 0, got '%d'", got.ID)
		}
		// test qu'on récupère bien la string passer en argument
		if got.Description != "Ma premiere task" {
			t.Errorf("expected description : 'Ma premiere task', got '%s'", got.Description)
		}
		// test qu'on récupère bien une todo
		if got.Status != "todo" {
			t.Errorf("expected status 'todo', got '%s'", got.Status)
		}
		// test qu'on as bien une date de set
		if got.CreatedAt.IsZero() {
			t.Errorf("Expected CreatedAt to be set, got '%v'", got.CreatedAt)
		}
		// test qu'on as bien une date de set
		if got.UpdatedAt.IsZero() {
			t.Errorf("Expected UpdatedAt to be set, got '%v'", got.UpdatedAt)
		}
	})

	t.Run("Should throw error if description is empty", func(t *testing.T) {
		_, err := NewTask("")

		if err == nil {
			t.Fatal("expected error for empty description, got nil")
		}
	})
}

func TestUpdateTask(t *testing.T) {
	t.Run("should update a task", func(t *testing.T) {
		// on commence par créer la nouvelle task
		task, _ := NewTask("ma premiere task")
		// on récupère l'ancienne date de mise à jour pour tester si la modif à bien été effectuer
		oldUpdatedAt := task.UpdatedAt
		// on viens faire l'update de la task en pointant sur la task
		err := task.UpdateDescription("New description")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if task.Description != "New description" {
			t.Errorf("expected description 'New description', got %q", task.Description)
		}

		// Renvoie true si la date sur laquelle tu l’appelles est postérieure à t.
		// Renvoie false si elle est égale ou antérieure.
		// permet de vérifier que la date à bien changer
		if !task.UpdatedAt.After(oldUpdatedAt) {
			t.Errorf("exected UpdatedAt to change")
		}

	})

	t.Run("should throw error if description is empty", func(t *testing.T) {
		task, _ := NewTask("Old description")

		err := task.UpdateDescription("   ")
		if err == nil {
			t.Fatal("expected error, got nil")
		}

		// Renvoie true si la date est la valeur zéro de time.Time (0001-01-01 00:00:00 UTC).
		// En gros, ça veut dire que la date n’a pas été définie.
		if !errors.Is(err, ErrEmptyDescription) {
			t.Errorf("expected ErrEmptyDescription, got %v", err)
		}
	})
}

func TestMarkIsDone(t *testing.T) {
	t.Run("Should set status Done", func(t *testing.T) {
		task, _ := NewTask("Ma super task")

		oldUpdatedAt := task.UpdatedAt

		err := task.MarkIsDone()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if task.Status != "done" {
			t.Errorf("expected status for the task 'done', got : %q", task.Status)
		}

		if !task.UpdatedAt.After(oldUpdatedAt) {
			t.Errorf("exected UpdatedAt to change")
		}
	})
}

func TestMarkIsProgress(t *testing.T) {
	t.Run("Should set status Done", func(t *testing.T) {
		task, _ := NewTask("Ma super task")

		oldUpdatedAt := task.UpdatedAt

		err := task.MarkIsProgress()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if task.Status != "in_progress" {
			t.Errorf("expected status for the task 'in_progress', got : %q", task.Status)
		}

		if !task.UpdatedAt.After(oldUpdatedAt) {
			t.Errorf("exected UpdatedAt to change")
		}
	})
}

// Add, Update, and Delete tasks
// Mark a task as in progress or done
// List all tasks
// List all tasks that are done
// List all tasks that are not done
// List all tasks that are in progress
