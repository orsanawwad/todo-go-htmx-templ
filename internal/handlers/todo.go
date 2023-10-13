package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/orsanawwad/htmxdemo/internal/templates"
	"github.com/orsanawwad/htmxdemo/internal/services"
)

type TodoService interface {
	List() ([]services.Todo, error)
	Create(pos int, val string) error
	// Update(pos int, val string, )
	Delete(pos int) error
}

type TodosHandler struct {
	TodoService TodoService
}

func NewTodosHandler(ts TodoService) *TodosHandler {
	return &TodosHandler{ts}
}

func (h *TodosHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		todos, err := h.TodoService.List()
		if err != nil {
			http.Error(w, "Failed to get todos", http.StatusInternalServerError)
		}
		templates.Todos(todos).Render(r.Context(), w)
	})

	r.Get("/list", func(w http.ResponseWriter, r *http.Request) {
		todos, err := h.TodoService.List()
		if err != nil {
			http.Error(w, "Failed to get todos", http.StatusInternalServerError)
		}
		templates.TodosList(todos).Render(r.Context(), w)
	})

	r.Get("/edit", func(w http.ResponseWriter, r *http.Request) {
		todos, err := h.TodoService.List()
		if err != nil {
			http.Error(w, "Failed to get todos", http.StatusInternalServerError)
		}
		templates.TodosEdit(todos).Render(r.Context(), w)
	})

	r.Put("/edit", func(w http.ResponseWriter, r *http.Request) {
		// todoVal := r.FormValue("todo")
		// if (todoVal != "") {
		// 	h.TodoService.Create(0, todoVal)
		// }
		// w.Header().Set("HX-Trigger", "ReloadTodos, ReloadTodosEdit")
		// templates.TodosForm().Render(r.Context(), w)
	})

	r.Delete("/delete/{todoPos}", func(w http.ResponseWriter, r *http.Request) {
		todoPosStr := chi.URLParam(r, "todoPos")
		todoPos, err := strconv.Atoi(todoPosStr)
		if err != nil {
			http.Error(w, "Failed to delete todo", http.StatusInternalServerError)
		}
		h.TodoService.Delete(todoPos)
		todos, err := h.TodoService.List()
		if err != nil {
			http.Error(w, "Failed to get todos", http.StatusInternalServerError)
		}
		templates.TodosEdit(todos).Render(r.Context(), w)
	})

	r.Put("/", func(w http.ResponseWriter, r *http.Request) {
		todoVal := r.FormValue("todo")
		if (todoVal != "") {
			h.TodoService.Create(0, todoVal)
		}
		w.Header().Set("HX-Trigger", "ReloadTodos, ReloadTodosEdit")
		templates.TodosForm().Render(r.Context(), w)
	})

	return r
}
