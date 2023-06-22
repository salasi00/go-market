package services

import (
	"github.com/salasi00/go-market/internal/models"
	"github.com/salasi00/go-market/internal/repositories"
)

type ProductService struct {
	productRepository repositories.ProductRepository
}

func NewProductService(productRepository repositories.ProductRepository) *ProductService {
	return &ProductService{productRepository}
}

func (s *ProductService) CreateProduct(product models.Product) (models.Product, error) {
	createdProduct, err := s.productRepository.CreateProduct(product)
	if err != nil {
		return models.Product{}, err
	}

	return createdProduct, nil
}

func (s *ProductService) FindProduct() ([]models.Product, error) {
	products, err := s.productRepository.FindProduct()
	if err != nil {
		return []models.Product{}, err
	}

	return products, nil
}

func (s *ProductService) GetProduct(ID int) (models.Product, error) {
	user, err := s.productRepository.GetProduct(ID)
	if err != nil {
		return models.Product{}, err
	}

	return user, nil
}
