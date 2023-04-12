package business

import (
	"context"
	"go-rest-api-todo-list/module/item/model"
)

type DeleteItemStorage interface {
	GetItem(
		ctx context.Context,
		condition map[string]interface{},
	) (*model.ToDoItem, error)

	DeleteItem(
		ctx context.Context,
		condition map[string]interface{},
	) error
}

type deleteItemBiz struct {
	store DeleteItemStorage
}

func NewDeleteItemBiz(store DeleteItemStorage) *deleteItemBiz {
	return &deleteItemBiz{store}
}

func (biz *deleteItemBiz) DeleteItem(
	ctx context.Context,
	condition map[string]interface{},
) error {
	// step 1: Find item by conditions
	_, err := biz.store.GetItem(ctx, condition)

	if err != nil {
		return err
	}

	// Step 2: call to storage to update item
	if err := biz.store.DeleteItem(ctx, condition); err != nil {
		return err
	}

	return nil
}
