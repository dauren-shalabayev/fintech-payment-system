package main

type Status string

const (
	StatusToRead  Status = "to_read"
	StatusReading Status = "reading"
	StatusRead    Status = "read"
)

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"-"` // prevent marshaling to JSON
}

type Book struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Status Status `json:"status" gorm:"default:to_read"`
	Author string `json:"author"`
	Year   int    `json:"year"`
	UserID uint   `json:"user_id"`
}
