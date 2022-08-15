package http

import (
	"net/http"

	"github.com/kapralovs/warehouse/internal/products"
)

type Handler struct {
	usecase products.UseCase
}

func NewHandler(uc products.UseCase) *Handler {
	return &Handler{
		usecase: uc,
	}
}

func (s *server) createProduct() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *server) editProduct() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *server) deleteProduct() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *server) getProducts() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *server) getProductByID() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
