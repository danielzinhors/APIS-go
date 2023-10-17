package main

import (
	"net/http"

	"github.com/danielzinhors/APIS-go/configs"
	"github.com/danielzinhors/APIS-go/internal/entity"
	"github.com/danielzinhors/APIS-go/internal/infra/database"
	"github.com/danielzinhors/APIS-go/internal/infra/webservers/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	configs, err := configs.LoadConfig(".")
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
	userHandler := handlers.NewUserHanbler(userDb, configs.TokenAuth, configs.JwtExperesIn)
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(configs.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Post("/users", userHandler.CreateUser)
	r.Post("/users/generate_token", userHandler.GetJWT)
	http.ListenAndServe(":8000", r)
}
