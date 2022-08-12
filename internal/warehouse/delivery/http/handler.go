package http

import "net/http"

type Handler struct {
	useCase warehouse.UseCase
}

func NewHandler(uc warehouse.UseCase) *Handler {
	return &Handler{
		useCase: uc,
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
