package main

import (
	"net/http"

	"github.com/danielzinhors/APIS-go/configs"
	"github.com/danielzinhors/APIS-go/internal/entity"
	"github.com/danielzinhors/APIS-go/internal/infra/database"
	"github.com/danielzinhors/APIS-go/internal/infra/webservers/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDb := database.NewProduct(db)
	productHandler := handlers.NewProductHanbler(productDb)
	http.HandleFunc("/products", productHandler.CreateProduct)
	http.ListenAndServe(":8000", nil)
}
