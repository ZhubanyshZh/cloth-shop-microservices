package services

import (
	"fmt"
	"github.com/ZhubanyshZh/go-project-service/internal/cache/product_cache"
	"github.com/ZhubanyshZh/go-project-service/internal/models"
	"github.com/ZhubanyshZh/go-project-service/internal/repositories"
	"github.com/jinzhu/copier"
	"log"
)

type ProductService struct {
	Repo  *repositories.ProductRepository
	Cache *product_cache.ProductCache
}

func NewProductService(repo *repositories.ProductRepository, cache *product_cache.ProductCache) *ProductService {
	return &ProductService{Repo: repo, Cache: cache}
}

func (s *ProductService) GetProducts() ([]models.Product, error) {
	return s.Repo.GetAll()
}

func (s *ProductService) GetProduct(id uint) (*models.Product, error) {
	cacheKey := s.Cache.BuildCacheKey(id)

	if product, found := s.Cache.GetFromCache(cacheKey); found {
		log.Println("🟢 Cache hit for product_cache:", id)
		return product, nil
	}

	productFromDB, err := s.FindProductById(id)
	if err != nil {
		return nil, err
	}

	s.Cache.SetToCache(cacheKey, productFromDB)
	return productFromDB, nil
}

func (s *ProductService) CreateProduct(productCreate *models.ProductCreate) error {
	product := &models.Product{}
	if err := copier.Copy(product, productCreate); err != nil {
		return fmt.Errorf("failed to copy product_cache data: %w", err)
	}

	if err := s.Repo.Create(product); err != nil {
		return err
	}

	s.Cache.SetToCache(s.Cache.BuildCacheKey(product.ID), product)
	return nil
}

func (s *ProductService) UpdateProduct(productUpdate *models.ProductUpdate) error {
	existingProduct, err := s.FindProductById(productUpdate.ID)
	if err != nil {
		return err
	}

	if err := copier.Copy(existingProduct, productUpdate); err != nil {
		return fmt.Errorf("failed to copy updated data: %w", err)
	}

	if err := s.Repo.Update(existingProduct); err != nil {
		return err
	}

	s.Cache.SetToCache(s.Cache.BuildCacheKey(existingProduct.ID), existingProduct)
	return nil
}

func (s *ProductService) DeleteProduct(id uint) error {
	if _, err := s.FindProductById(id); err != nil {
		return err
	}

	if err := s.Repo.Delete(id); err != nil {
		return err
	}

	s.Cache.DeleteFromCache(s.Cache.BuildCacheKey(id))
	return nil
}

func (s *ProductService) FindProductById(id uint) (*models.Product, error) {
	return s.Repo.GetByID(id)
}
