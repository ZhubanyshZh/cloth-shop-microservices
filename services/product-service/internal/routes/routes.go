package routes

import (
	"github.com/ZhubanyshZh/go-project-service/internal/models"
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
		middlewares.ValidateProductMiddleware(&models.ProductCreate{})(
			http.HandlerFunc(handler.UpdateProduct),
		)).Methods(http.MethodPost)

	r.Handle(
		"/products/{id:[0-9]+}",
		http.HandlerFunc(handler.GetProduct)).Methods(http.MethodGet)

	r.Handle(
		"/products/{id:[0-9]+}",
		http.HandlerFunc(handler.DeleteProduct)).Methods(http.MethodDelete)

	r.Handle(
		"/products",
		middlewares.ValidateProductMiddleware(&models.ProductUpdate{})(
			http.HandlerFunc(handler.UpdateProduct),
		)).Methods(http.MethodPut)
	return r
}
