package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User represents a user in the system
type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username" binding:"required"`
	Password string `gorm:"not null" json:"-"` // never JSON-export the hash
}

// HashPassword hashes the user's password using bcrypt.
func (u *User) HashPassword() error {
	hp, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hp)
	return nil
}

// CheckPassword compares a plaintext pw against the stored hash.
func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

