package model

import (
	"errors"
	"go-rest-api-todo-list/common"
	"go-rest-api-todo-list/module/item"
	"strings"
)

var (
	ErrTitleCannotBeBlank       = errors.New("title can not be blank")
	ErrItemNotFound             = errors.New("item not found")
	ErrCannotUpdateFinishedItem = errors.New("can not update finished item")
	ErrTitleCannotBeEmpty       = errors.New("title cannot be empty")
)

type ToDoItem struct {
	common.SQLModel
	Title  string           `json:"title" gorm:"column:title;"`
	Status *item.ItemStatus `json:"status" gorm:"column:status;"`
}

func (ToDoItem) TableName() string { return "todo_items" }

type ToDoItemCreation struct {
	Id    int    `json:"-" gorm:"column:id;"` // return id of todoItem after creating
	Title string `json:"title" gorm:"column:title;"`
}

func (t *ToDoItemCreation) Validate() error {
	t.Title = strings.TrimSpace(t.Title)

	if t.Title == "" {
		return ErrTitleCannotBeEmpty
	}
	return nil
}

func (ToDoItemCreation) TableName() string { return ToDoItem{}.TableName() }

type ToDoItemUpdate struct {
	Title  *string `json:"title" gorm:"column:title;"`
	Status *string `json:"status" gorm:"column:status;"`
}

func (ToDoItemUpdate) TableName() string { return ToDoItem{}.TableName() }
