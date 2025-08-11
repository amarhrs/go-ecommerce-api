package models

type Product struct {
	ID          uint            `gorm:"primaryKey" json:"id"`
	Name        string          `json:"name"`
	CategoryID  *uint           `json:"category_id"`
	Category    ProductCategory `gorm:"foreignKey:CategoryID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"category,omitempty"`
	Price       float64         `json:"price"`
	Stock       int             `json:"stock"`
	Description string          `json:"description"`
	Images      []ProductImage  `gorm:"foreignKey:ProductID" json:"images,omitempty"` // Validasi format & minimal jumlah
}
