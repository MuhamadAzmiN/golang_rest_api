package services

import (
    "github.com/username/go_rest_api_crud/model"
    "github.com/username/go_rest_api_crud/repo"
)

type ProductService struct {
    repo *repo.ProductRepository
}

func NewProductService(repo *repo.ProductRepository) *ProductService {
    return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(product *models.Product) error {
    return s.repo.Create(product)
}

func (s *ProductService) GetAllProducts() ([]models.Product, error) {
    return s.repo.FindAll()
}

func (s *ProductService) GetProduct(id uint) (*models.Product, error) {
    return s.repo.FindByID(id)
}

func (s *ProductService) UpdateProduct(product *models.Product) error {
    return s.repo.Update(product)
}

func (s *ProductService) DeleteProduct(id uint) error {
    return s.repo.Delete(id)
}
