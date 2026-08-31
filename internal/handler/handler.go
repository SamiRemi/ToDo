package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"todo/internal/repository"
	"todo/internal/service"
)

type TaskHandler struct {
	service *service.TaskService
}

func NewTaskHandler(s *service.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

// listTasks обрабатывает GET /tasks (список всех задач)
func (h *TaskHandler) listTasks(w http.ResponseWriter, r *http.Request) {
	tasks := h.service.ListTasks()
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

// createTask обрабатывает POST /tasks (создание задачи)
func (h *TaskHandler) createTask(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title  string `json:"title"`
		Status string `json:"status"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}
	if input.Title == "" {
		http.Error(w, "title is required", http.StatusBadRequest)
		return
	}

	task := &repository.Task{
		Title:  input.Title,
		Status: input.Status,
	}

	created := h.service.CreateTask(task)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(created); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

// getTaskByID обрабатывает GET /tasks/{id}
func (h *TaskHandler) getTaskByID(w http.ResponseWriter, r *http.Request, id int) {
	task, ok := h.service.GetTask(id)
	if !ok {
		http.Error(w, "task not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

// updateTaskByID обрабатывает PATCH /tasks/{id}
func (h *TaskHandler) updateTaskByID(w http.ResponseWriter, r *http.Request, id int) {
	var input struct {
		Title  *string `json:"title"`
		Status *string `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	updated, ok := h.service.UpdateTask(id, input.Title, input.Status)
	if !ok {
		http.Error(w, "task not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(updated); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

// deleteTaskByID обрабатывает DELETE /tasks/{id}
func (h *TaskHandler) deleteTaskByID(w http.ResponseWriter, r *http.Request, id int) {
	ok := h.service.DeleteTask(id)
	if !ok {
		http.Error(w, "task not found", http.StatusNotFound)
		return
	}
	// Возвращаем 204 No Content без тела ответа
	w.WriteHeader(http.StatusNoContent)
}

// RegisterRoutes настраивает все маршруты
func (h *TaskHandler) RegisterRoutes(mux *http.ServeMux) {
	// Маршрут для списка задач и создания новой задачи: /tasks
	mux.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.listTasks(w, r)
		case http.MethodPost:
			h.createTask(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Маршрут для операций с конкретной задачей: /tasks/{id}
	// net/http ServeMux позволяет использовать префикс "/tasks/"
	mux.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		// Путь будет выглядеть как "/tasks/123"
		path := r.URL.Path

		// Убираем префикс "/tasks/" и получаем ID
		idStr := strings.TrimPrefix(path, "/tasks/")

		// Если после префикса ничего нет или там не число — ошибка
		if idStr == "" {
			http.Error(w, "missing task ID", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil || id <= 0 {
			http.Error(w, "invalid task ID", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			h.getTaskByID(w, r, id)
		case http.MethodPatch:
			h.updateTaskByID(w, r, id)
		case http.MethodDelete:
			h.deleteTaskByID(w, r, id)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
