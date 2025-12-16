package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/pz3-http/internal/storage"
)

func TestCreateTask(t *testing.T) {
	store := storage.NewMemoryStore()
	h := NewHandlers(store)

	body := bytes.NewBufferString(`{"title":"test task"}`)
	req := httptest.NewRequest(http.MethodPost, "/tasks", body)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	h.CreateTask(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", w.Code)
	}
}

func TestGetTaskNotFound(t *testing.T) {
	store := storage.NewMemoryStore()
	h := NewHandlers(store)

	req := httptest.NewRequest(http.MethodGet, "/tasks/999", nil)
	w := httptest.NewRecorder()

	h.GetTask(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", w.Code)
	}
}
