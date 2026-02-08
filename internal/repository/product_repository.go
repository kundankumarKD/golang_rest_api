package repository

import (
	"product-api/internal/domain"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (r *ProductRepository) CreateProduct(product *domain.Product) error {
	return r.DB.Create(product).Error
}

func (r *ProductRepository) GetAllProducts() ([]domain.Product, error) {
	var products []domain.Product
	err := r.DB.Find(&products).Error
	return products, err
}

func (r *ProductRepository) GetProductByID(id uint) (*domain.Product, error) {
	var product domain.Product
	err := r.DB.First(&product, id).Error
	return &product, err
}

func (r *ProductRepository) UpdateProduct(product *domain.Product) error {
	return r.DB.Save(product).Error
}

func (r *ProductRepository) DeleteProduct(id uint) error {
	return r.DB.Delete(&domain.Product{}, id).Error
}
