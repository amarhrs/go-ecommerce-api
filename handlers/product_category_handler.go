package handlers

import (
	"amarhrs/ecommerce/helpers"
	"amarhrs/ecommerce/models"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func ListProductCategories(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var categories []models.ProductCategory
		if err := db.Preload("Products").
			Where("is_active = ?", true).
			Find(&categories).Error; err != nil {
			helpers.Error(ctx, http.StatusInternalServerError, "Failed to fetch categories")
			return
		}
		helpers.Success(ctx, http.StatusOK, "Categories fetched successfully", gin.H{"data": categories})
	}
}

func UpdateProductCategoryStatus(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		var category models.ProductCategory

		if err := db.First(&category, id).Error; err != nil {
			helpers.Error(ctx, http.StatusNotFound, "Category not found")
			return
		}

		// Struktur input JSON, hanya menerima is_active
		var input struct {
			IsActive *bool `json:"is_active"`
		}

		if err := ctx.ShouldBindJSON(&input); err != nil {
			helpers.Error(ctx, http.StatusBadRequest, "Invalid input")
			return
		}

		if input.IsActive == nil {
			helpers.Error(ctx, http.StatusBadRequest, "is_active is required")
			return
		}

		if err := db.Model(&category).Update("is_active", *input.IsActive).Error; err != nil {
			helpers.Error(ctx, http.StatusInternalServerError, "Failed to update category status")
			return
		}

		helpers.Success(ctx, http.StatusOK, "Category status updated successfully", gin.H{"data": category})
	}
}

func GetProductCategory(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		var category models.ProductCategory
		if err := db.Preload("Products").First(&category, id).Error; err != nil {
			helpers.Error(ctx, http.StatusNotFound, "Category not found")
			return
		}
		helpers.Success(ctx, http.StatusOK, "Category fetched successfully", gin.H{"data": category})
	}
}

func CreateProductCategory(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input models.ProductCategory
		if err := ctx.ShouldBindJSON(&input); err != nil {
			helpers.Error(ctx, http.StatusBadRequest, "Invalid input")
			return
		}

		input.Name = strings.TrimSpace(input.Name)
		if input.Name == "" {
			helpers.Error(ctx, http.StatusBadRequest, "Name cannot be empty")
			return
		}

		if err := db.Create(&input).Error; err != nil {
			helpers.Error(ctx, http.StatusInternalServerError, "Failed to create category")
			return
		}
		helpers.Success(ctx, http.StatusCreated, "Category created successfully", gin.H{"data": input})
	}
}

func UpdateProductCategory(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		var category models.ProductCategory
		if err := db.First(&category, id).Error; err != nil {
			helpers.Error(ctx, http.StatusNotFound, "Category not found")
			return
		}

		var input models.ProductCategory
		if err := ctx.ShouldBindJSON(&input); err != nil {
			helpers.Error(ctx, http.StatusBadRequest, "Invalid input")
			return
		}

		input.Name = strings.TrimSpace(input.Name)
		if input.Name == "" {
			helpers.Error(ctx, http.StatusBadRequest, "Name cannot be empty")
			return
		}

		if err := db.Model(&category).Update("name", input.Name).Error; err != nil {
			helpers.Error(ctx, http.StatusInternalServerError, "Failed to update category")
			return
		}
		helpers.Success(ctx, http.StatusOK, "Category updated successfully", gin.H{"data": category})
	}
}

func DeleteProductCategory(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		var category models.ProductCategory
		if err := db.First(&category, id).Error; err != nil {
			helpers.Error(ctx, http.StatusNotFound, "Category not found")
			return
		}

		// Set CategoryID = NULL untuk semua produk yang punya kategori ini
		db.Model(&models.Product{}).Where("category_id = ?", category.ID).Update("category_id", nil)

		if err := db.Delete(&category).Error; err != nil {
			helpers.Error(ctx, http.StatusInternalServerError, "Failed to delete category")
			return
		}

		helpers.Success(ctx, http.StatusOK, "Category deleted successfully", gin.H{
			"id":         category.ID,
			"name":       category.Name,
			"deleted_at": time.Now().Format(time.RFC3339),
		})
	}
}
