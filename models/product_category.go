package models

type ProductCategory struct {
	ID       uint      `gorm:"primaryKey" json:"id"`
	Name     string    `json:"name"`
	Products []Product `gorm:"foreignKey:CategoryID" json:"products,omitempty"`
	IsActive bool      `gorm:"default:true" json:"is_active"` // Untuk validasi kategori aktif/tidak
}
