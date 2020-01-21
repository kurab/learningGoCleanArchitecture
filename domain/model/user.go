package model

type User struct {
    Id    int    `gorm:"primary_key" json:"id"`
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Email string `json:"email"`
}

type Users []User
