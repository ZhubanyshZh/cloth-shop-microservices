package routes

import (
	"net/http"

	"github.com/ZhubanyshZh/go-project-service/internal/handlers"
	"github.com/ZhubanyshZh/go-project-service/internal/middlewares"

	"github.com/gorilla/mux"
)

func RegisterRoutes(handler *handlers.ProductHandler) *mux.Router {
	r := mux.NewRouter()
	r.Handle(
		"/products",
		http.HandlerFunc(handler.GetProducts)).Methods(http.MethodGet)

	r.Handle(
		"/products",
		middlewares.ValidateProductMiddleware(
			http.HandlerFunc(handler.CreateProduct))).Methods(http.MethodPost)

	r.Handle(
		"/products/{id:[0-9]+}",
		http.HandlerFunc(handler.GetProduct)).Methods(http.MethodGet)

	r.Handle(
		"/products/{id:[0-9]+}",
		http.HandlerFunc(handler.DeleteProduct)).Methods("DELETE")

	r.Handle(
		"/products",
		middlewares.ValidateProductMiddleware(
			http.HandlerFunc(handler.UpdateProduct))).Methods(http.MethodPut)
	return r
}
