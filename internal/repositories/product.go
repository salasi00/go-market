package repositories

import (
	"github.com/salasi00/go-market/internal/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(product models.Product) (models.Product, error)
	FindProduct() ([]models.Product, error)
	GetProduct(ID int) (models.Product, error)
}

func NewProductRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateProduct(product models.Product) (models.Product, error) {
	err := r.db.Create(&product).Error

	return product, err
}

func (r *repository) FindProduct() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Preload("User").Find(&products).Error

	return products, err
}

func (r *repository) GetProduct(ID int) (models.Product, error) {
	var product models.Product
	err := r.db.First(&product, ID).Error

	return product, err
}
