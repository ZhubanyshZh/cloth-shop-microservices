package routes

import (
	"github.com/ZhubanyshZh/go-project-service/internal/dto"
	"net/http"
	"os"

	"github.com/ZhubanyshZh/go-project-service/internal/handlers"
	"github.com/ZhubanyshZh/go-project-service/internal/middlewares"

	"github.com/gorilla/mux"
)

func RegisterRoutes(handler *handlers.ProductHandler) *mux.Router {
	var baseUrl = "/api/" + os.Getenv("API_VERSION") + "/products"
	r := mux.NewRouter()
	r.Handle(
		baseUrl,
		http.HandlerFunc(handler.GetProducts)).Methods(http.MethodGet)

	r.Handle(
		baseUrl,
		http.HandlerFunc(handler.CreateProduct),
	).Methods(http.MethodPost)

	r.Handle(
		baseUrl+"/{id:[0-9]+}",
		http.HandlerFunc(handler.GetProduct)).Methods(http.MethodGet)

	r.Handle(
		baseUrl+"/{id:[0-9]+}",
		http.HandlerFunc(handler.DeleteProduct)).Methods(http.MethodDelete)

	r.Handle(
		baseUrl,
		middlewares.ValidateProductMiddleware(&dto.ProductUpdate{})(
			http.HandlerFunc(handler.UpdateProduct),
		)).Methods(http.MethodPut)
	return r
}
