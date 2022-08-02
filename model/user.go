package model

// User represents user entity
type User struct {
	ID       string `faker:"uuid_hyphenated"`
	Email    string `gorm:"unique" faker:"email"`
	Password string `faker:"password"`
}
