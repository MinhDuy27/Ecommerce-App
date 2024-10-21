package repository

import (
	"github.com/MinhDuy27/Ecommerce-App/domain"
	"github.com/MinhDuy27/Ecommerce-App/internal/dto"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductRepository interface {
	CreateProduct(p dto.CreateProductDto) error
	FindProduct(id uint) (domain.Product, error)
	UpdateProduct(id uint, p domain.Product) error
	DeleteProduct(id uint) error
	GetAllProduct(amount int) ([]domain.Product, error)
}

func GetProductImage(db *gorm.DB) productRepository {
	return productRepository{
		Db: db,
	}
}

type productRepository struct {
	Db *gorm.DB
}

func (pr productRepository) CreateProduct(p dto.CreateProductDto) error {
	product := domain.Product{
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Quantity:    p.Quantity,
		Image_url:   p.Image_url,
	}
	if err := pr.Db.Create(&product).Error; err != nil {
		return err
	}
	return nil
}
func (pr productRepository) FindProduct(id uint) (domain.Product, error) {
	var product domain.Product
	if err := pr.Db.First(&product, id).Error; err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (pr productRepository) UpdateProduct(id uint, p domain.Product) error {
	var product domain.Product
	product_update := map [string]interface{} {
		"name": p.Name,
		"description": p.Description,
		"price": p.Price,
		"quantity": p.Quantity,
		"image_url": p.Image_url,
	}
	err := pr.Db.Model(&product).Clauses(clause.Locking{Strength: "UPDATE"}).Where("ID=?", id).Updates(product_update).Error
	if err != nil {
		return err
	}
	return nil
}

func (pr productRepository) DeleteProduct(id uint) error {
	var product domain.Product
	err := pr.Db.Delete(&product, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (pr productRepository) GetAllProduct(amount int) ([]domain.Product, error) {
	var products []domain.Product
	if err := pr.Db.Limit(amount).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}