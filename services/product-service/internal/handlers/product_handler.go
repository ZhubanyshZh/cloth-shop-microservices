package handlers

import (
	"encoding/json"
	"github.com/ZhubanyshZh/go-project-service/internal/models"
	"github.com/ZhubanyshZh/go-project-service/internal/services"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	Service *services.ProductService
}

func getIdFromRequest(r *http.Request, w http.ResponseWriter) (int, error) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		handleError(w, err, "❌ Invalid product ID", http.StatusBadRequest)
		return 0, err
	}
	return id, nil
}

func handleError(w http.ResponseWriter, err error, message string, statusCode int) {
	log.Println("❌", message, err)
	http.Error(w, message, statusCode)
}

func decodeJSONRequest[T any](w http.ResponseWriter, r *http.Request, dst *T) bool {
	if err := json.NewDecoder(r.Body).Decode(dst); err != nil {
		handleError(w, err, "❌ Invalid JSON format", http.StatusBadRequest)
		return false
	}
	return true
}

func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r, w)
	if err != nil {
		return
	}

	product, err := h.Service.GetProduct(uint(id))
	if err != nil {
		handleError(w, err, "❌ Product not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.ProductCreate
	if !decodeJSONRequest(w, r, &product) {
		return
	}

	if err := h.Service.CreateProduct(&product); err != nil {
		handleError(w, err, "❌ Error creating product", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.ProductUpdate
	if !decodeJSONRequest(w, r, &product) {
		return
	}

	if err := h.Service.UpdateProduct(&product); err != nil {
		handleError(w, err, "❌ Error updating product", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r, w)
	if err == nil {
		return
	}

	if err := h.Service.DeleteProduct(uint(id)); err != nil {
		handleError(w, err, "❌ Product not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.Service.GetProducts()
	if err != nil {
		handleError(w, err, "❌ Error fetching products", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
