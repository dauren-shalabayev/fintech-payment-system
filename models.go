package main

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"-"` // prevet masrshaling to JSON
}

type Book struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
}
