package models

// Todo represents a task with a title, description, and completion status.
// It includes an auto-incrementing primary key ID, a Title, a Description, and a Completed flag.
type Todo struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}
