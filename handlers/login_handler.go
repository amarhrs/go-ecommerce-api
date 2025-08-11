package handlers

import (
	"amarhrs/ecommerce/helpers"
	"amarhrs/ecommerce/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func Login(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input models.User
		if err := ctx.ShouldBindJSON(&input); err != nil {
			helpers.Error(ctx, http.StatusBadRequest, "Invalid input")
			return
		}

		// Validasi login
		if ok, errs := helpers.ValidateLoginInput(input); !ok {
			helpers.Error(ctx, http.StatusBadRequest, errs)
			return
		}

		var user models.User
		if err := db.Where("username = ?", input.Username).First(&user).Error; err != nil {
			helpers.Error(ctx, http.StatusUnauthorized, "Invalid username or password")
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
			helpers.Error(ctx, http.StatusUnauthorized, "Invalid username or password")
			return
		}

		// Bikin Access Token & Refresh Token
		accessToken, err := CreateAccessToken(user.ID)
		if err != nil {
			helpers.Error(ctx, http.StatusInternalServerError, "Failed to generate access token")
			return
		}

		refreshToken, err := CreateRefreshToken(user.ID)
		if err != nil {
			helpers.Error(ctx, http.StatusInternalServerError, "Failed to generate refresh token")
			return
		}

		helpers.Success(ctx, http.StatusOK, "Login successful", gin.H{
			"user": gin.H{
				"id":       user.ID,
				"username": user.Username,
			},
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		})
	}
}
