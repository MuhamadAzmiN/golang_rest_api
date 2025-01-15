package repo

import (
    "gorm.io/gorm"
    "github.com/username/go_rest_api_crud/model"
)

type ProductRepository struct {
    db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
    return &ProductRepository{db: db}
}

/*************  ✨ Codeium Command ⭐  *************/
// Create adds a new product to the database.
/******  05242946-8c44-45c2-baac-e27114721d66  *******/
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
    return &product, err
}

func (r *ProductRepository) Update(product *models.Product) error {
    return r.db.Save(product).Error
}

func (r *ProductRepository) Delete(id uint) error {
    return r.db.Delete(&models.Product{}, id).Error
}
