package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/orsanawwad/htmxdemo/internal/templates"
)

type MainHandler struct {
}

func NewMainHandler() *MainHandler {
	return &MainHandler{}
}

func (h *MainHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		h.View(w, r)
	})
	return r
}

func (h *MainHandler) View(w http.ResponseWriter, r *http.Request) {
	templates.Index().Render(r.Context(), w)
}
