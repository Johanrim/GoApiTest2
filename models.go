package main

import (
	"time"
)

// type User struct {
// 	ID        uint           `gorm:"primaryKey" json:"id"`
// 	Account   string         `gorm:"not null; unique;" json:"account"`
// 	Password  string         `gorm:"not null" json:"password"`
// 	Name      string         `gorm:"not null" json:"name"`
// 	Age       uint           `gorm:"not null" json:"age"`
// 	Phone     string         `gorm:"not null" json:"phone"`
// 	Address   string         `gorm:"not null" json:"address"`
// 	Email     string         `gorm:"not null" json:"email"`
// 	BookID    uint           `gorm:"not null" json:"book"`
// 	Book      []Book         `gorm:"ForeignKey:BookID;" json:"-"`
// 	CreatedAt time.Time      `json:"created_at"`
// 	UpdatedAt time.Time      `json:"updated_at"`
// 	DeletedAt gorm.DeletedAt `gorm:"index"`
// }

// func (u *User) TableName() string {
// 	return "user"
// }

type Book struct {
	ID         uint      `gorm:"primaryKey"`
	Name       string    `gorm:"not null" json:"name"`
	CategoryID uint      `gzorm:"not null" json:"category"`
	Category   Category  `gorm:"ForeignKey:CategoryID;" json:"-"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	// DeletedAt  gorm.DeletedAt `gorm:"index"`
	// AuthorID   uint           `gorm:"not null" json:"author"`
	// Author     Author         `gorm:"ForeignKey:AuthorID;"json:"-"`
}

func (b *Book) TableName() string {
	return "book"
}

type Category struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"not null" json:"name"`
}

func (c *Category) TableName() string {
	return "category"
}

// type Author struct {
// 	ID    uint   `gorm:"primaryKey" json:"id"`
// 	Name  string `gorm:"not null" json:"name"`
// 	Age   uint   `gorm:"not null" json:"age"`
// 	Phone string `gorm:"not null" json:"phone"`
// }

// func (a *Author) TableName() string {
// 	return "author"
// }
