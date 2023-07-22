package util

import (
	"time"
)

type User struct {
	ID        string    `gorm:"column:id;primaryKey"`
	FirstName string    `gorm:"column:first_name"`
	LastName  string    `gorm:"column:last_name"`
	Email     string    `gorm:"column:email;unique"`
	Password  string    `gorm:"column:password"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

type Config struct {
	Host string
	Port string
	User string
	Pass string
	DB   string
}
