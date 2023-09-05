package models

type Product struct {
	Id          int    `gorm:"primary_key" json:"id"`
	ProductName string `gorm:"type:varchar(300)" json:"product_name"`
	Description string `gorm:"type:text" json:"description"`
}
