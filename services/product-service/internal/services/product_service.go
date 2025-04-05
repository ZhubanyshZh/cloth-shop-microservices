package services

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/copier"
	"strconv"
	"time"

	"github.com/ZhubanyshZh/go-project-service/internal/cache"
	"github.com/ZhubanyshZh/go-project-service/internal/models"
	"github.com/ZhubanyshZh/go-project-service/internal/repositories"
)

type ProductService struct {
	Repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{Repo: repo}
}

func (s *ProductService) GetProducts() ([]models.Product, error) {
	return s.Repo.GetAll()
}

func (s *ProductService) GetProduct(id uint) (*models.Product, error) {
	cacheKey := fmt.Sprintf("product:%d", id)

	if cached, err := cache.GetCache(cacheKey); err == nil {
		var product models.Product
		json.Unmarshal([]byte(cached), &product)
		fmt.Println("🔄 Cache Hit!")
		return &product, nil
	}

	product, err := s.FindProductById(id)
	if err != nil {
		return nil, err
	}

	productJSON, _ := json.Marshal(product)
	cache.SetCache(cacheKey, string(productJSON), 10*time.Minute)

	return product, nil
}

func (s *ProductService) CreateProduct(createProduct *models.ProductCreate) error {
	product := &models.Product{}
	copier.Copy(product, createProduct)
	fmt.Println(product, createProduct)
	createErr := s.Repo.Create(product)
	if createErr != nil {
		return createErr
	}
	productJSON, _ := json.Marshal(product)
	cacheKey := fmt.Sprintf("product:%d", product.ID)
	cache.SetCache(cacheKey, string(productJSON), 10*time.Minute)
	return nil
}

func (s *ProductService) UpdateProduct(productUpdate *models.ProductUpdate) error {
	product, err := s.FindProductById(productUpdate.ID)
	if err != nil {
		return err
	}
	copier.Copy(product, productUpdate)
	updErr := s.Repo.Update(product)
	if updErr != nil {
		return updErr
	}
	productJSON, _ := json.Marshal(product)
	cacheKey := fmt.Sprintf("product:%d", product.ID)
	cache.UpdateCache(cacheKey, string(productJSON), 10*time.Minute)
	return nil
}

func (s *ProductService) DeleteProduct(id uint) error {
	_, err := s.FindProductById(id)
	if err != nil {
		return err
	}
	delErr := s.Repo.Delete(id)
	if delErr != nil {
		return delErr
	}
	cache.DeleteCache(strconv.Itoa(int(id)))
	return nil
}

func (s *ProductService) FindProductById(id uint) (*models.Product, error) {
	product, err := s.Repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}
