package repo

import (
	"errors"

	"github.com/username/go_rest_api_crud/model"
	"gorm.io/gorm"
)

type ProductRepository struct {
    db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
    return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(product *models.Product) error {
    return r.db.Create(product).Error
}

func (r *ProductRepository) FindAll() ([]models.Product, error) {
    var products []models.Product
    err := r.db.Find(&products).Error
    return products, err
}

func (r *ProductRepository) FindByID(id uint) (*models.Product, error) {
    var product models.Product
    err := r.db.First(&product, id).Error

    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, errors.New("product not found")
    }

    return &product, err
}

func (r *ProductRepository) Update(product *models.Product) error {
    return r.db.Save(product).Error
}

func (r *ProductRepository) Delete(id uint) error {
    return r.db.Delete(&models.Product{}, id).Error
}
