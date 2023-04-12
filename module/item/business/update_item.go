package business

import (
	"context"
	"go-rest-api-todo-list/module/item"
	"go-rest-api-todo-list/module/item/model"
)

type UpdateItemStorage interface {
	GetItem(
		ctx context.Context,
		condition map[string]interface{},
	) (*model.ToDoItem, error)

	UpdateItem(
		ctx context.Context,
		condition map[string]interface{},
		dataUpdate *model.ToDoItemUpdate,
	) error
}

type updateItemBiz struct {
	store UpdateItemStorage
}

func NewUpdateItemBiz(store UpdateItemStorage) *updateItemBiz {
	return &updateItemBiz{store: store}
}

func (biz *updateItemBiz) UpdateItem(
	ctx context.Context,
	condition map[string]interface{},
	dataUpdate *model.ToDoItemUpdate,
) error {
	// step 1: Find item by conditions
	oldItem, err := biz.store.GetItem(ctx, condition)

	if err != nil {
		return err
	}

	// just a demo in case we don't allow update Finished item
	if *(oldItem.Status) == item.ItemStatusDone {
		return model.ErrCannotUpdateFinishedItem
	}

	// Step 2: call to storage to update item
	if err := biz.store.UpdateItem(ctx, condition, dataUpdate); err != nil {
		return err
	}

	return nil
}
