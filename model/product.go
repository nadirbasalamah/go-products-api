package model

// Product represents product entity
type Product struct {
	ID          string  `faker:"uuid_hyphenated"`
	Name        string  `faker:"name"`
	Description string  `faker:"word"`
	Price       float64 `faker:"oneof: 15, 27, 61"`
	Stock       int     `faker:"oneof: 15, 27, 61"`
}
