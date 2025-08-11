package main

import (
	"amarhrs/ecommerce/configs"
	"amarhrs/ecommerce/handlers"
	"amarhrs/ecommerce/middlewares"
	"amarhrs/ecommerce/migrations"
	"amarhrs/ecommerce/seeders"
	"net/http"

	_ "net/http/pprof"

	"github.com/gin-gonic/gin"
)

func main() {
	// Koneksi Database
	db, err := configs.InitDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	migrations.Migrate(db)
	seeders.Seed(db)

	router := gin.Default()

	router.GET("/products", handlers.ListProducts(db))
	router.GET("/products/:id", handlers.GetProduct(db))

	router.POST("/products", middlewares.AuthMiddleware(), handlers.CreateProduct(db))
	router.PUT("/products/:id", middlewares.AuthMiddleware(), handlers.UpdateProduct(db))
	router.DELETE("/products/:id", middlewares.AuthMiddleware(), handlers.DeleteProduct(db))

	router.PATCH("/categories/:id/status", handlers.UpdateProductCategoryStatus(db))

	router.POST("/login", handlers.Login(db))
	router.POST("/register", handlers.Register(db))

	router.GET("/debug/pprof/*pprof", gin.WrapH(http.DefaultServeMux))

	router.Run(":8080")
}
