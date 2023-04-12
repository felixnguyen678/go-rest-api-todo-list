package storage

import (
	"context"
	"go-rest-api-todo-list/module/item/model"
)

func (s *mysqlStorage) CreateItem(ctx context.Context, data *model.ToDoItemCreation) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}

	return nil
}
