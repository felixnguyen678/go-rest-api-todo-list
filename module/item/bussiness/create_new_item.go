package bussiness

import (
	"context"
	"go-rest-api-todo-list/module/item/model"
)

type CreateTodoItemStorage interface {
	CreateItem(ctx context.Context, data *model.ToDoItemCreation) error
}

type createBiz struct {
	store CreateTodoItemStorage
}

func NewCreateToDoItemBiz(store CreateTodoItemStorage) *createBiz {
	return &createBiz{store: store}
}

func (biz *createBiz) CreateNewItem(ctx context.Context, data *model.ToDoItemCreation) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.store.CreateItem(ctx, data); err != nil {
		return err
	}

	return nil
}
