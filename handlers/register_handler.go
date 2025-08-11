package handlers

import (
	"amarhrs/ecommerce/helpers"
	"amarhrs/ecommerce/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func Register(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input models.User
		if err := ctx.ShouldBindJSON(&input); err != nil {
			helpers.Error(ctx, http.StatusBadRequest, "Invalid input")
			return
		}

		// Validasi semua field
		if ok, errs := helpers.ValidateRegisterInput(input); !ok {
			helpers.Error(ctx, http.StatusBadRequest, errs)
			return
		}

		// Check duplicate username
		var existingUser models.User
		if err := db.Where("username = ?", input.Username).First(&existingUser).Error; err == nil {
			helpers.Error(ctx, http.StatusConflict, "Username already exists")
			return
		}

		// Check duplicate email
		if err := db.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
			helpers.Error(ctx, http.StatusConflict, "Email already exists")
			return
		}

		// Hash password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			helpers.Error(ctx, http.StatusInternalServerError, "Failed to hash password")
			return
		}

		// Create new user
		newUser := models.User{
			Username: input.Username,
			Email:    input.Email,
			Password: string(hashedPassword),
		}
		if err := db.Create(&newUser).Error; err != nil {
			helpers.Error(ctx, http.StatusInternalServerError, "Failed to create user")
			return
		}

		// Respon sukses
		helpers.Success(ctx, http.StatusCreated, "User registered successfully", gin.H{
			"user": gin.H{
				"id":       newUser.ID,
				"username": newUser.Username,
			},
		})
	}
}
