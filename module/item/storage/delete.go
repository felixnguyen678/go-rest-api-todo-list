package storage

import (
	"context"
	"go-rest-api-todo-list/module/item/model"
)

func (s *mysqlStorage) DeleteItem(
	ctx context.Context,
	condition map[string]interface{},
) error {
	if err := s.db.Table(model.
		ToDoItem{}.TableName()).
		Where(condition).
		Delete(nil).Error; err != nil {
		return err
	}

	return nil
}
