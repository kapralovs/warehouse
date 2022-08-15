package http

import (
	"net/http"

	"github.com/kapralovs/warehouse/internal/tasks"
)

type Handler struct {
	usecase tasks.UseCase
}

func NewHandler(u tasks.UseCase) *Handler {
	return &Handler{
		usecase: u,
	}
}

func (h *Handler) createTask() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *Handler) editTask() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *Handler) deleteTask() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *Handler) getTasks() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *Handler) geTaskByID() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *Handler) startTask() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *Handler) completeTask() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
