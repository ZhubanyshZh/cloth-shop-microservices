package routes

import (
	"fmt"
	"github.com/ZhubanyshZh/go-project-service/internal/middlewares"
	"net/http"
	"os"

	"github.com/ZhubanyshZh/go-project-service/internal/handlers"
	"github.com/gorilla/mux"
)

func RegisterRoutes(handler *handlers.ProductHandler) *mux.Router {
	apiVersion := os.Getenv("API_VERSION")
	baseURL := fmt.Sprintf("/api/%s/products", apiVersion)

	r := mux.NewRouter()
	productRouter := r.PathPrefix(baseURL).Subrouter()

	productRouter.HandleFunc("", handler.GetProducts).Methods(http.MethodGet)
	productRouter.HandleFunc("/{id:[0-9]+}", handler.GetProduct).Methods(http.MethodGet)

	productRouter.HandleFunc("/{id:[0-9]+}", handler.DeleteProduct).Methods(http.MethodDelete)

	createSub := productRouter.PathPrefix("").Subrouter()
	createSub.Use(middlewares.ProductCreateMiddleware)
	createSub.HandleFunc("", handler.CreateProduct).Methods(http.MethodPost)

	updateSub := productRouter.PathPrefix("").Subrouter()
	updateSub.Use(middlewares.ProductUpdateMiddleware)
	updateSub.HandleFunc("", handler.UpdateProduct).Methods(http.MethodPut)

	return r
}
