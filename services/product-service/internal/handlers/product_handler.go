package handlers

import (
	"net/http"

	"github.com/ZhubanyshZh/go-project-service/internal/models"
	"github.com/ZhubanyshZh/go-project-service/internal/services"
	"github.com/ZhubanyshZh/go-project-service/internal/utils"
)

type ProductHandler struct {
	Service *services.ProductService
}

func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetIDFromRequest(r)
	if err != nil {
		utils.HandleError(w, err, "❌ Invalid product_cache ID", http.StatusBadRequest)
		return
	}

	product, err := h.Service.GetProduct(uint(id))
	if err != nil {
		utils.HandleError(w, err, "❌ Product not found", http.StatusNotFound)
		return
	}

	utils.WriteJSON(w, http.StatusOK, product)
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.Service.GetProducts()
	if err != nil {
		utils.HandleError(w, err, "❌ Error fetching products", http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusOK, products)
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.ProductCreate
	if !utils.DecodeJSONRequest(w, r, &product) {
		return
	}

	if err := h.Service.CreateProduct(&product); err != nil {
		utils.HandleError(w, err, "❌ Error creating product_cache", http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, product)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.ProductUpdate
	if !utils.DecodeJSONRequest(w, r, &product) {
		return
	}

	if err := h.Service.UpdateProduct(&product); err != nil {
		utils.HandleError(w, err, "❌ Error updating product_cache", http.StatusBadRequest)
		return
	}

	utils.WriteJSON(w, http.StatusOK, product)
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetIDFromRequest(r)
	if err != nil {
		utils.HandleError(w, err, "❌ Invalid product_cache ID", http.StatusBadRequest)
		return
	}

	if err := h.Service.DeleteProduct(uint(id)); err != nil {
		utils.HandleError(w, err, "❌ Product not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
