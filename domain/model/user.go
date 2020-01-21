package model

type User struct {
    Id    int    `gorm:"primary_key" json:"id"`
    Name  string `json:"name" validate:"required"`
    Age   int    `json:"age" validate:"gte=0,lt=256"`
    Email string `json:"email" validate:"required,email"`
}

type Users []User
