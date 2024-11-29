package models

type Wishlist struct{
	ID uint `gorm:"primarykey"`
	UserID int `gorm:"not null"`
	ProductID int `gorm:"not null"`
}