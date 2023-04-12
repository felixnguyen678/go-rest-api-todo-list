package storage

import (
	"context"
	"go-rest-api-todo-list/module/item/model"
)

func (s *mysqlStorage) GetItem(ctx context.Context, condition map[string]interface{}) (*model.ToDoItem, error) {
	var data model.ToDoItem

	if err := s.db.Where(condition).First(&data).Error; err != nil {
		return nil, model.ErrItemNotFound
	}
	return &data, nil
}
