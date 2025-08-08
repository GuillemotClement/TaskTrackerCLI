package task

import "testing"

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
