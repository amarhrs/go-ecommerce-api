package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Kirim respon error JSON
func Error(ctx *gin.Context, statusCode int, messages interface{}) {
	resp := gin.H{
		"error": gin.H{
			"code": statusCode,
		},
	}

	switch v := messages.(type) {
	case string:
		resp["error"].(gin.H)["message"] = v
	case []string:
		resp["error"].(gin.H)["messages"] = v
	default:
		resp["error"].(gin.H)["message"] = "An unexpected error occurred"
		statusCode = http.StatusInternalServerError
	}

	ctx.JSON(statusCode, resp)
	ctx.Abort()
}

// Kirim respon sukses JSON
func Success(ctx *gin.Context, statusCode int, message string, data gin.H) {
	resp := gin.H{"message": message}

	// Tambahkan data tambahan ke respon
	for k, v := range data {
		if k != "message" { // cegah overwrite message
			resp[k] = v
		}
	}

	ctx.JSON(statusCode, resp)
}
