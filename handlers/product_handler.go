package handlers

import (
	"amarhrs/ecommerce/helpers"
	"amarhrs/ecommerce/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func ListProducts(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var products []models.Product
		query := db.Preload("Category").Preload("Images").
			Joins("JOIN product_categories ON products.category_id = product_categories.id").
			Where("product_categories.is_active = ?", true)

		if err := query.Find(&products).Error; err != nil {
			helpers.Error(ctx, http.StatusInternalServerError, "Failed to fetch products")
			return
		}

		// if err := db.Preload("Category").Preload("Images").Find(&products).Error; err != nil {
		// 	helpers.Error(ctx, http.StatusInternalServerError, "Failed to fetch products")
		// 	return
		// }
		helpers.Success(ctx, http.StatusOK, "Products fetched successfully", gin.H{"data": products})
	}
}

func GetProduct(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		var product models.Product

		err := db.Preload("Category").
			Preload("Images").
			Joins("JOIN product_categories ON products.category_id = product_categories.id").
			Where("products.id = ? AND product_categories.is_active = ?", id, true).
			First(&product).Error

		if err != nil {
			helpers.Error(ctx, http.StatusNotFound, "Product not found or category inactive")
			return
		}

		// if err := db.Preload("Category").Preload("Images").First(&product, id).Error; err != nil {
		// 	helpers.Error(ctx, http.StatusNotFound, "Product not found")
		// 	return
		// }

		helpers.Success(ctx, http.StatusOK, "Product fetched successfully", gin.H{"data": product})
	}
}

func CreateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input models.Product
		if err := ctx.ShouldBindJSON(&input); err != nil {
			helpers.Error(ctx, http.StatusBadRequest, "Invalid input")
			return
		}

		// Kumpulkan semua error validasi
		var validationErrors []string

		// Validasi umum (nama, harga, stok, deskripsi, gambar)
		if ok, errs := helpers.ValidateProductInput(input); !ok {
			validationErrors = append(validationErrors, errs...)
		}

		// Validasi kategori
		if input.CategoryID == nil {
			validationErrors = append(validationErrors, "Category is required")
		} else {
			var category models.ProductCategory
			if err := db.First(&category, *input.CategoryID).Error; err != nil {
				validationErrors = append(validationErrors, "Invalid CategoryID")
			} else if !category.IsActive {
				validationErrors = append(validationErrors, "Category is inactive")
			}
		}

		// Cek duplikasi nama per kategori
		if input.CategoryID != nil {
			var existing models.Product
			if err := db.Where("LOWER(name) = ?", strings.ToLower(input.Name)).
				Where("category_id = ?", *input.CategoryID).
				First(&existing).Error; err == nil {
				validationErrors = append(validationErrors, "Product with the same name already exists in this category")
			}
		}

		// Jika ada error validasi, kirim semua
		if len(validationErrors) > 0 {
			helpers.Error(ctx, http.StatusBadRequest, validationErrors)
			return
		}

		// Insert data
		if err := db.Create(&input).Error; err != nil {
			helpers.Error(ctx, http.StatusInternalServerError, "Failed to create product")
			return
		}

		// Ambil ulang data dengan relasi Category
		if err := db.Preload("Category").Preload("Images").First(&input, input.ID).Error; err != nil {
			helpers.Error(ctx, http.StatusInternalServerError, "Failed to load product data")
			return
		}

		helpers.Success(ctx, http.StatusCreated, "Product created successfully", gin.H{
			"product": input,
		})
	}
}

func UpdateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		var product models.Product
		if err := db.First(&product, id).Error; err != nil {
			helpers.Error(ctx, http.StatusNotFound, "Product not found")
			return
		}

		var input models.Product
		if err := ctx.ShouldBindJSON(&input); err != nil {
			helpers.Error(ctx, http.StatusBadRequest, "Invalid input")
			return
		}

		var validationErrors []string

		// Validasi umum dari helper
		if ok, errs := helpers.ValidateProductInput(input); !ok {
			validationErrors = append(validationErrors, errs...)
		}

		// Cek duplikasi nama produk (kecuali produk yang sedang di-update)
		if input.CategoryID != nil {
			var existing models.Product
			if err := db.Where("LOWER(name) = ?", strings.ToLower(input.Name)).
				Where("category_id = ?", *input.CategoryID).
				Where("id <> ?", product.ID).
				First(&existing).Error; err == nil {
				validationErrors = append(validationErrors, "Product with the same name already exists in this category")
			}
		} else {
			validationErrors = append(validationErrors, "Category is required")
		}

		// Validasi kategori aktif
		if input.CategoryID != nil {
			var category models.ProductCategory
			if err := db.First(&category, *input.CategoryID).Error; err != nil {
				validationErrors = append(validationErrors, "Invalid CategoryID")
			} else if !category.IsActive {
				validationErrors = append(validationErrors, "Category is inactive")
			}
		}

		// Jika ada error validasi, kirim semua
		if len(validationErrors) > 0 {
			helpers.Error(ctx, http.StatusBadRequest, validationErrors)
			return
		}

		if err := db.Model(&product).Updates(input).Error; err != nil {
			helpers.Error(ctx, http.StatusInternalServerError, "Failed to update product")
			return
		}

		// Hapus dulu semua images lama yg terkait product
		db.Where("product_id = ?", product.ID).Delete(&models.ProductImage{})

		// Insert ulang images baru (kalau ada)
		for _, img := range input.Images {
			img.ProductID = product.ID // pastikan product_id diset
			db.Create(&img)
		}

		// Reload product dengan kategori
		db.Preload("Category").Preload("Images").First(&product, product.ID)

		helpers.Success(ctx, http.StatusOK, "Product updated successfully", gin.H{"data": product})
	}
}

func DeleteProduct(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		var product models.Product
		if err := db.First(&product, id).Error; err != nil {
			helpers.Error(ctx, http.StatusNotFound, "Product not found")
			return
		}

		if err := db.Delete(&product).Error; err != nil {
			helpers.Error(ctx, http.StatusInternalServerError, "Failed to delete product")
			return
		}
		helpers.Success(ctx, http.StatusOK, "Product deleted successfully", gin.H{"id": product.ID})
	}
}
