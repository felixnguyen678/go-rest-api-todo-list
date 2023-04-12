package gin_item

import (
	"github.com/gin-gonic/gin"
	"go-rest-api-todo-list/common"
	"go-rest-api-todo-list/module/item/business"
	"go-rest-api-todo-list/module/item/model"
	"go-rest-api-todo-list/module/item/storage"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetItemById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dataItem model.ToDoItemCreation

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Where("id = ?", id).First(&dataItem).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		store := storage.NewMySQLStorage(db)
		biz := business.NewCreateToDoItemBiz(store)

		if err := biz.CreateNewItem(c.Request.Context(), &dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(dataItem))
	}
}
