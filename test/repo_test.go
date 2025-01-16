// test/repository_test.go
package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/username/go_rest_api_crud/model"
	"github.com/username/go_rest_api_crud/repo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
    // Gunakan database khusus untuk testing
    dsn := "root:@tcp(localhost:3306)/go_rest?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to test database: " + err.Error())
    }
    
    // Migrate table
    db.AutoMigrate(&models.Product{})
    
    return db
}

func cleanupTestData(db *gorm.DB) {
    // Bersihkan data test setelah selesai
    db.Exec("DELETE FROM products")
}

func create(db *gorm.DB){
	db.Exec("INSERT INTO products (name, price) VALUES ('test', 10000)")
}



func TestProductRepository_FindAll(t *testing.T) {
    // Setup
    db := setupTestDB()
    repo := repo.NewProductRepository(db)
    
    // // Cleanup setelah test selesai
    t.Cleanup(func() {
        cleanupTestData(db)
    })

    // Insert test data
    testProducts := []models.Product{
        {
            Name:  "Product 1",
            Price: 10000,
        },
    }

    for _, product := range testProducts {
        err := db.Create(&product).Error
        assert.NoError(t, err)
    }

    // Execute FindAll
    products, err := repo.FindAll()

    // Assertions
    assert.NoError(t, err)
    assert.NotNil(t, products)
    assert.Equal(t, 1, len(products))
    assert.Equal(t, "Product 1", products[0].Name)
    assert.Equal(t, float64(10000), products[0].Price)
}

func TestProductRepository_Create(t *testing.T) {
    // Setup database
    db := setupTestDB()
    repo := repo.NewProductRepository(db)
    
    // Cleanup setelah test selesai
    t.Cleanup(func() {
        cleanupTestData(db)
    })

    // Test case 1: Create Success
    t.Run("Create Success", func(t *testing.T) {
        product := &models.Product{
            Name:  "tetsd",
            Price: 15000,
			Description: "Description",
        }

        // Execute Create
        err := repo.Create(product)

        // Assertions
        assert.NoError(t, err)
        assert.NotZero(t, product.ID) // memastikan ID terisi
        
        // Verifikasi data tersimpan di database
        var savedProduct models.Product
        result := db.First(&savedProduct, product.ID)
        assert.NoError(t, result.Error)
        assert.Equal(t, product.Name, savedProduct.Name)
        assert.Equal(t, product.Price, savedProduct.Price)
    })

    // // Test case 2: Create with Empty Name (Error case)
    // t.Run("Create with Empty Name", func(t *testing.T) {
    //     product := &models.Product{
    //         Name:  "", // empty name
    //         Price: 15000,
    //     }

    //     // Execute Create
    //     err := repo.Create(product)

    //     // Assertions
    //     assert.Error(t, err) // harusnya error karena nama kosong
    // })

}


func TestDeleteProduct(t *testing.T) {
	db := setupTestDB()
    repo := repo.NewProductRepository(db)

	// Cleanup setelah test selesai
	t.Cleanup(func() {
		cleanupTestData(db)
	})
    
	t.Run("Delete Success", func(t *testing.T) {
		product := &models.Product{
			Name:  "test",
			Price: 15000,
			Description: "Description",
		}

		err := repo.Create(product)
		assert.NoError(t, err)
		assert.NotZero(t, product.ID)

		
		err = repo.Delete(product.ID)
		assert.NoError(t, err)



		// ini biar  ngecek dah ke hapus belum data nya
		deleteProduct, err := repo.FindByID(product.ID)
		fmt.Println(deleteProduct)
		assert.Error(t, err)
		assert.Nil(t , deleteProduct) 
	})
}



 func TestProductGetById(t *testing.T) {
    db := setupTestDB()
    repo := repo.NewProductRepository(db)

    t.Cleanup(func() {
        cleanupTestData(db)
    })

    t.Run("GetById Success", func(t *testing.T) {
        product := &models.Product{
            Name:  "test",
            Price: 15000,
            Description: "Description",
        }


        err := repo.Create(product)
        assert.NoError(t, err)
        assert.NotZero(t, product.ID)


        product, err = repo.FindByID(product.ID)
        assert.NoError(t, err)
        assert.Equal(t, product.Name, "test")
    })
 }