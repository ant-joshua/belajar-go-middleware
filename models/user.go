package models

import (
	"belajar-middleware/helpers"
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Fullname string    `gorm:"not null" json:"fullname" valid:"required~Your fullname is required"`
	Email    string    `gorm:"not null;unique" json:"email" valid:"required~Your email is required,email~Your email is not valid"`
	Password string    `gorm:"not null" json:"password" valid:"required~Your password is required,,minstringlength(6)~Your password must be at least 6 characters"`
	Products []Product `gorm:"foreignkey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Role     string    `gorm:"default:user" json:"role" valid:"required~Your role is required"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(user)
	if errCreate != nil {
		err = errCreate
		return
	}

	// hash password
	hashedPassword, err := helpers.HashPass(user.Password)
	if err != nil {
		return
	}

	user.Password = hashedPassword
	return nil
}
