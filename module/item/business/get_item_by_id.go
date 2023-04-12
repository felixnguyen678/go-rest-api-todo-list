package business

import (
	"context"
	"go-rest-api-todo-list/module/item/model"
)

type GetItemStorage interface {
	GetItem(ctx context.Context, condition map[string]interface{}) (*model.ToDoItem, error)
}

type getItemBiz struct {
	store GetItemStorage
}

func NewGetItemBiz(store GetItemStorage) *getItemBiz {
	return &getItemBiz{store: store}
}

func (biz *getItemBiz) GetItemById(ctx context.Context, id int) (*model.ToDoItem, error) {
	_condition := map[string]interface{}{"id": id}

	data, err := biz.store.GetItem(ctx, _condition)
	if err != nil {
		return nil, err
	}
	return data, nil
}
