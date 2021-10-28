package model

// User represents user entity
type User struct {
	ID       string
	Email    string `gorm:"unique"`
	Password string
}
