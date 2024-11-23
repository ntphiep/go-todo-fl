package data

import (
	"time"
)

// -------- todo
type ToDoItem struct {
	Id          int        `json:"id" gorm:"column:id;"`
	Title       string     `json:"title" gorm:"column:title;"`
	Description string     `json:"description" gorm:"column:description;"`
	Status      string     `json:"status" gorm:"column:status;"`
	CreatedAt   *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt   *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func (ToDoItem) TableName() string {
	return "todo_items"
}

type TodoItemCreation struct {
	Id          int    `json:"-" gorm:"column:id"` // auto increment
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
	Status      string `json:"status" gorm:"column:status"`
}

func (TodoItemCreation) TableName() string {
	return "todo_items"
}

// -------- user
type User struct {
	Id    int    `json:"id" gorm:"column:id;"`
	Name  string `json:"name" gorm:"column:name;"`
	Email string `json:"email" gorm:"column:email"`
}

// func (u *User) ConfigUserEmail(newEmail string) {
// 	u.Email = newEmail
// }
