package models

type Product struct {
	GormModel
	Title       string `gorm:"not null" json:"title" valid:"required~Your title is required"`
	Description string `gorm:"not null" json:"description" valid:"required~Your description is required"`
	UserID      uint
	User        *User
}
