package models

import (
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint32    `gorm:"primarykey;autoincremenet" json:"id"`
	Nickname  string    `gorm:"not null;unique;size:20" json:"nickname"`
	Email     string    `gorm:"size:50;not null;unique" json:"email"`
	Password  string    `gorm:"size:60;not null" json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GenerateHashedPassword ..
func GenerateHashedPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword ..
func VerifyPassword(hashedPassword, passwordFromUser string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(passwordFromUser))
}

// BeforeSave ..
func (u *User) BeforeSave(*gorm.DB) error {
	if u.Password == "" {
		return fmt.Errorf("Can't save user as the password is invalid, %s", u.Password)
	}

	b, err := GenerateHashedPassword(u.Password)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return err
	}
	u.Password = string(b)
	return nil
}
