package http

import (
	"net/http"

	"github.com/kapralovs/warehouse/internal/warehouse"
)

type Handler struct {
	usecase warehouse.UseCase
}

func NewHandler(uc warehouse.UseCase) *Handler {
	return &Handler{
		usecase: uc,
	}
}

func (h *Handler) createCell() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *Handler) editCell() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *Handler) deleteCell() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *Handler) getCells() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *Handler) getCellByID() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
