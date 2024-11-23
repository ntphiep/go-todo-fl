package data

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
	"time"
)

// -------- status

type ItemStatus int

const (
	ItemStatusTodo ItemStatus = iota
	ItemStatusDoing
	ItemStatusDone
	ItemStatusDeleted
)

var allStatus = []string{"todo", "doing", "done", "deleted"}

func (i *ItemStatus) String() string {
	return allStatus[*i]
}

func parseStatus(status string) (ItemStatus, error) {
	for i := range allStatus {
		if allStatus[i] == status {
			return ItemStatus(i), nil
		}
	}
	return ItemStatus(0), errors.New("invalid status")
}

func (i *ItemStatus) Scan(value any) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("invalid type")
	}

	v, err := parseStatus(string(bytes))
	if err != nil {
		return errors.New("invalid status")
	}

	*i = v
	return nil
}

func (i *ItemStatus) Value() (driver.Value, error) {
	if i == nil {
		return nil, nil
	}
	return i.String(), nil
}

func (i *ItemStatus) MarshalJSON() ([]byte, error) {
	if i == nil {
		return []byte("null"), nil
	}

	return []byte(fmt.Sprintf(`"%s"`, i.String())), nil
}

func (i *ItemStatus) UnmarshalJSON(data []byte) error {
	str := strings.ReplaceAll(string(data), `\"`, "")

	v, err := parseStatus(str)
	if err != nil {
		return err
	}

	*i = v
	return nil
}

// -------- paging
type DataPaging struct {
	Page  int   `json:"page" form:"page"`
	Limit int   `json:"limit" form:"limit"`
	Total int64 `json:"total" form:"-"`
}

func (p *DataPaging) Process() {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.Limit <= 0 || p.Limit > 100 {
		p.Limit = 10
	}

}

// -------- todo

type ToDoItem struct {
	Id          int         `json:"id" gorm:"column:id;"`
	Title       string      `json:"title" gorm:"column:title;"`
	Description string      `json:"description" gorm:"column:description;"`
	Status      *ItemStatus `json:"status" gorm:"column:status;"`
	CreatedAt   *time.Time  `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt   *time.Time  `json:"updated_at,omitempty" gorm:"column:updated_at;"`
}

func (ToDoItem) TableName() string {
	return "todo_items"
}

type TodoItemCreate struct {
	Id          int    `json:"-" gorm:"column:id"` // auto increment
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
	Status      *ItemStatus `json:"status" gorm:"column:status"`
}

func (TodoItemCreate) TableName() string {
	return ToDoItem{}.TableName()
}

type TodoItemEdit struct {
	Title       *string `json:"title" gorm:"column:title"`
	Description *string `json:"description" gorm:"column:description"`
	Status      *string `json:"status" gorm:"column:status"`
}

func (TodoItemEdit) TableName() string {
	return ToDoItem{}.TableName()
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
