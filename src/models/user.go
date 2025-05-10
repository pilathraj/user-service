// models/user.go
package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"primaryKey autoIncrement"`
	Username  string    `gorm:"unique;not null" json:"username"`
	Email     string    `gorm:"unique;not null" json:"email"`
	Password  string    `gorm:"not null" json:"-"`
	CreatedAt time.Time `gorm:"autoCreateTime default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"autoUpdateTime default:CURRENT_TIMESTAMP"`
}

type UserRole struct {
	ID        uint      `gorm:"primaryKey autoIncrement"`
	UserID    uint      `json:"userId"`
	Role      string    `json:"role"`
	CreatedAt time.Time `gorm:"autoCreateTime default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"autoUpdateTime default:CURRENT_TIMESTAMP"`
}

type UserPreference struct {
	ID        uint      `gorm:"primaryKey autoIncrement"`
	UserID    uint      `json:"userId"`
	Key       string    `json:"key"`
	Value     string    `json:"value"`
	CreatedAt time.Time `gorm:"autoCreateTime default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"autoUpdateTime default:CURRENT_TIMESTAMP"`
}

type ContactInformation struct {
	ID        uint      `gorm:"primaryKey autoIncrement"`
	UserID    uint      `json:"userId"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	CreatedAt time.Time `gorm:"autoCreateTime default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"autoUpdateTime default:CURRENT_TIMESTAMP"`
}

// Migrate the schema
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{}, &UserRole{}, &UserPreference{}, &ContactInformation{})
}
