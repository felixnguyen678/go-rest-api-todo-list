package storage

import (
	"context"
	"go-rest-api-todo-list/module/item/model"
)

func (s *mysqlStorage) UpdateItem(
	ctx context.Context,
	condition map[string]interface{},
	dataUpdate *model.ToDoItemUpdate,
) error {
	if err := s.db.Where(condition).Updates(dataUpdate).Error; err != nil {
		return err
	}

	return nil
}
