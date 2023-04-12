package business

import (
	"context"
	"go-rest-api-todo-list/common"
	"go-rest-api-todo-list/module/item/model"
)

type ListTodoItemStorage interface {
	ListItems(
		ctx context.Context,
		condition map[string]interface{},
		paging *common.Paging,
	) ([]model.ToDoItem, error)
}

type listBiz struct {
	store ListTodoItemStorage
}

func NewListToDoItemBiz(store ListTodoItemStorage) *listBiz {
	return &listBiz{store: store}
}

func (biz *listBiz) ListItems(ctx context.Context,
	condition map[string]interface{},
	paging *common.Paging,
) ([]model.ToDoItem, error) {
	result, err := biz.store.ListItems(ctx, condition, paging)

	if err != nil {
		return nil, err
	}

	return result, err
}
