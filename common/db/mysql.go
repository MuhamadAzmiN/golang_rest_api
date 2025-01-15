package db

import (
    "fmt"
    "log"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "github.com/username/go_rest_api_crud/model"
)

type Config struct {
    Host     string
    Port     string
    User     string
    Password string
    DBName   string
}

func NewConnection(config *Config) *gorm.DB {
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        config.User,
        config.Password,
        config.Host,
        config.Port,
        config.DBName,
    )

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    log.Println("Database connected successfully")
    db.AutoMigrate(&models.Product{})

    return db
}