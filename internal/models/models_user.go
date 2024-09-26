package models

type User struct {
	Id           uint   `gorm:"primary_key"`
	Username     string `gorm:"unique;not null"`
	Email        string `gorm:"unique;not null"`
	HashPassword string `gorm:"unique;not null"`
}
