package service

import (
	"github.com/MinhDuy27/Ecommerce-App/domain"
	"github.com/MinhDuy27/Ecommerce-App/internal/dto"
	"github.com/MinhDuy27/Ecommerce-App/internal/repository"
)

type ProductService struct {
	Repo repository.ProductRepository
}


func (p *ProductService) Create(input dto.CreateProductDto) error {
	if err := p.Repo.CreateProduct(input); err != nil {
		return err
	}
	return nil
}
func (p *ProductService) Find(id uint) (domain.Product, error) {
	value, err := p.Repo.FindProduct(id)
	if err != nil {
		return domain.Product{}, err
	}
	return value, nil
}

func (p *ProductService) Update(id uint, input domain.Product) error {
	if err := p.Repo.UpdateProduct(id, input); err != nil {
		return err
	}
	return nil
}

func (p *ProductService) Delete(id uint) error {
	if err := p.Repo.DeleteProduct(id); err != nil {
		return err
	}
	return nil
}

func (p *ProductService) GetAll(amount int) ([]domain.Product, error) {
	value, err := p.Repo.GetAllProduct(amount)
	if err != nil {
		return nil, err
	}
	return value, nil
}