package repositories

import (
	"github.com/salasi00/go-market/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user models.User) (models.User, error)
	GetUser(ID int) (models.User, error)
}

func NewUserRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateUser(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

// func (r *repository) FindUser()

func (r *repository) GetUser(ID int) (models.User, error) {
	var user models.User
	err := r.db.Preload("Products").First(&user, ID).Error

	return user, err
}
