package repository_test

import (
	"testing"
	"todo/internal/repository"
)

func TestInMemoryRepo_Create(t *testing.T) {
	repo := repository.NewInMemoryRepo()
	task := &repository.Task{Title: "Test task", Status: "pending"}
	created := repo.Create(task)

	if created.ID == 0 {
		t.Errorf("expected ID > 0, got %d", created.ID)
	}
	if created.Title != task.Title {
		t.Error("title mismatch")
	}
}

func TestInMemoryRepo_GetNotFound(t *testing.T) {
	repo := repository.NewInMemoryRepo()
	_, ok := repo.Get(999)
	if ok {
		t.Error("expected task not found")
	}
}
