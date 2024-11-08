package data

import (
	"fmt"
	"time"
)

// -------- todo
type ToDoItem struct {
	Id        int        `json:"id" gorm:"column:id;"`
	Title     string     `json:"title" gorm:"column:title;"`
	Status    string     `json:"status" gorm:"column:status;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}
 
func (ToDoItem) TableName() string {
	return "todo_items"
}

func clm() {
  fmt.Println("Hola mundo")
}



// -------- user
type User struct {
  Id int `json:"id" gorm:"column:id;"`
  Name string `json:"name" gorm:"column:name;"`
  Age int `json:"age" gorm:"column:age;"`
  Email string `json:"email" gorm:"column:email;`
}


func (u *User) ConfigUserEmail(newEmail string) {
  u.Email = newEmail;
}
