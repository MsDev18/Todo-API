package entity

type User struct {
	ID          uint
	Name        string
	PhoneNumber string

	// // Refrence to cateogry entity
	// CategoryIDs []uint
	// // Refrence to todo entity
	// TodoIDs     []uint
}
