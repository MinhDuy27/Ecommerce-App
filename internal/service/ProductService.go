package service

import (
	"github.com/MinhDuy27/Ecommerce-App/domain"
	"github.com/MinhDuy27/Ecommerce-App/internal/dto"
	"github.com/MinhDuy27/Ecommerce-App/internal/repository"
)

type ProductService struct {
	Rp repository.ProductRepository
}


func (p *ProductService) Create(input dto.CreateProductDto) error {
	if err := p.Rp.CreateProduct(input); err != nil {
		return err
	}
	return nil
}
func (p *ProductService) Find(id uint) (domain.Product, error) {
	value, err := p.Rp.FindProduct(id)
	if err != nil {
		return domain.Product{}, err
	}
	return value, nil
}

func (p *ProductService) Update(id uint, input domain.Product) error {
	if err := p.Rp.UpdateProduct(id, input); err != nil {
		return err
	}
	return nil
}

func (p *ProductService) Delete(id uint) error {
	if err := p.Rp.DeleteProduct(id); err != nil {
		return err
	}
	return nil
}

func (p *ProductService) GetAll(amount int) ([]domain.Product, error) {
	value, err := p.Rp.GetAllProduct(amount)
	if err != nil {
		return nil, err
	}
	return value, nil
}