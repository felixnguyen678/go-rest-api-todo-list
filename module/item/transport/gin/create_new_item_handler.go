package gin_item

import (
	"github.com/gin-gonic/gin"
	"go-rest-api-todo-list/common"
	"go-rest-api-todo-list/module/item/business"
	"go-rest-api-todo-list/module/item/model"
	"go-rest-api-todo-list/module/item/storage"
	"gorm.io/gorm"
	"net/http"
)

func CreateItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dataItem model.ToDoItemCreation

		if err := c.ShouldBind(&dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := storage.NewMySQLStorage(db)
		biz := business.NewCreateToDoItemBiz(store)

		if error := biz.CreateNewItem(c.Request.Context(), &dataItem); error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(dataItem.Id))
		return
	}
}
