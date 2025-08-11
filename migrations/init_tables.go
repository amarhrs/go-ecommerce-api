package migrations

import (
	"amarhrs/ecommerce/models"

	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.Product{},
		&models.ProductCategory{},
		&models.ProductImage{},
		&models.User{},
	)
}
