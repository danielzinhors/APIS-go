package main

import (
	"net/http"

	"github.com/danielzinhors/APIS-go/configs"
	"github.com/danielzinhors/APIS-go/internal/entity"
	"github.com/danielzinhors/APIS-go/internal/infra/database"
	"github.com/danielzinhors/APIS-go/internal/infra/webservers/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
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
	//http.HandleFunc("/products", productHandler.CreateProduct) router go interno
	userDb := database.NewUser(db)
	userHandler := handlers.NewUserHanbler(userDb)
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/products", productHandler.CreateProduct)
	r.Get("/products", productHandler.GetProducts)
	r.Get("/products/{id}", productHandler.GetProduct)
	r.Put("/products/{id}", productHandler.UpdateProduct)
	r.Delete("/products/{id}", productHandler.DeleteProduct)

	r.Post("/users", userHandler.CreateUser)

	http.ListenAndServe(":8000", r)
}
