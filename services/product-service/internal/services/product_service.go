package services

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/copier"
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

func (s *ProductService) CreateProduct(productEdit *models.ProductEdit) error {
	product := &models.Product{}
	copier.Copy(product, productEdit)
	fmt.Println(product, productEdit)
	return s.Repo.Create(product)
}

func (s *ProductService) UpdateProduct(productEdit *models.ProductEdit) error {
	product := &models.Product{}
	copier.Copy(product, productEdit)
	return s.Repo.Update(product)
}

func (s *ProductService) DeleteProduct(id uint) error {
	_, err := s.FindProductById(id)
	if err != nil {
		return err
	}
	return s.Repo.Delete(id)
}

func (s *ProductService) FindProductById(id uint) (*models.Product, error) {
	product, err := s.Repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}
