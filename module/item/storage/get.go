package storage

import (
	"context"
	"go-rest-api-todo-list/common"
	"go-rest-api-todo-list/module/item/model"
)

func (s *mysqlStorage) GetItem(ctx context.Context, condition map[string]interface{}) (*model.ToDoItem, error) {
	var data model.ToDoItem

	if err := s.db.Where(condition).First(&data).Error; err != nil {
		return nil, model.ErrItemNotFound
	}
	return &data, nil
}

func (s *mysqlStorage) ListItems(
	ctx context.Context,
	condition map[string]interface{},
	paging *common.Paging,
) ([]model.ToDoItem, error) {

	offset := (paging.Page - 1) * paging.Limit
	var result []model.ToDoItem

	if err := s.db.Table(model.ToDoItem{}.TableName()).
		Count(&paging.Total).
		Offset(offset).
		Order("id desc").
		Find(&result).Error; err != nil {

		return nil, model.ErrItemNotFound
	}

	return result, nil

}
