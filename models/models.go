package models

import "gorm.io/gorm"

type Message struct {
	Description string `json:"description"`
}

type Blog struct {
	gorm.Model
	Content  string    `json:"content"`
	Author   string    `json:"author"`
	Comments []Comment `json:"comments" gorm:foreignkey:BlogID"`
}

type Comment struct {
	gorm.Model
	Content string `json:"content"`
	Author  string `json:"author"`
	BlogID  int    `json:"blogid"`
}
