package seeders

import (
	"amarhrs/ecommerce/models"
	"log"

	"golang.org/x/crypto/bcrypt"

	"github.com/jinzhu/gorm"
)

func Seed(db *gorm.DB) {
	seedUsers(db)
	seedCategories(db)
	seedProducts(db)
}

func seedUsers(db *gorm.DB) {
	var count int
	db.Model(&models.User{}).Count(&count)
	if count > 0 {
		log.Println("Seeder User dilewati (sudah ada data)")
		return
	}

	password, _ := bcrypt.GenerateFromPassword([]byte("Password123"), bcrypt.DefaultCost)

	users := []models.User{
		{Username: "user1", Email: "user1@example.com", Password: string(password)},
		{Username: "user2", Email: "user2@example.com", Password: string(password)},
	}

	for _, user := range users {
		db.Create(&user)
	}

	log.Println("Seeder User berhasil dijalankan")
}

func seedCategories(db *gorm.DB) {
	var count int
	db.Model(&models.ProductCategory{}).Count(&count)
	if count > 0 {
		log.Println("Seeder Category dilewati (sudah ada data)")
		return
	}

	categories := []models.ProductCategory{
		{Name: "Elektronik", IsActive: false},
		{Name: "Fashion", IsActive: true},
		{Name: "Makanan", IsActive: true},
	}

	for _, category := range categories {
		db.Create(&category)
	}

	log.Println("Seeder Category berhasil dijalankan")
}

func seedProducts(db *gorm.DB) {
	var count int
	db.Model(&models.Product{}).Count(&count)
	if count > 0 {
		log.Println("Seeder Product dilewati (sudah ada data)")
		return
	}

	var elektronik, fashion, makanan models.ProductCategory
	db.Where("name = ?", "Elektronik").First(&elektronik)
	db.Where("name = ?", "Fashion").First(&fashion)
	db.Where("name = ?", "Makanan").First(&makanan)

	products := []models.Product{
		{
			Name:        "Laptop XYZ",
			CategoryID:  &elektronik.ID,
			Price:       10000000,
			Stock:       5,
			Description: "Laptop performa tinggi untuk pekerjaan berat",
			Images: []models.ProductImage{
				{URL: "/uploads/products/laptop1.jpg"},
				{URL: "/uploads/products/laptop2.jpg"},
			},
		},
		{
			Name:        "Kaos Polos",
			CategoryID:  &fashion.ID,
			Price:       75000,
			Stock:       50,
			Description: "Kaos polos bahan katun adem",
			Images: []models.ProductImage{
				{URL: "/uploads/products/kaos1.jpg"},
			},
		},
		{
			Name:        "Roti Coklat",
			CategoryID:  &makanan.ID,
			Price:       15000,
			Stock:       100,
			Description: "Roti lembut isi coklat premium",
			Images: []models.ProductImage{
				{URL: "/uploads/products/roti1.jpg"},
				{URL: "/uploads/products/roti2.jpg"},
			},
		},
	}

	for _, product := range products {
		db.Create(&product)
	}

	log.Println("Seeder Product berhasil dijalankan")
}
