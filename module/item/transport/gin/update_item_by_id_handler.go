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

func UpdateItemById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var dataItem model.ToDoItemUpdate

		if err := c.ShouldBind(&dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := storage.NewMySQLStorage(db)
		biz := business.NewUpdateItemBiz(store)

		err = biz.UpdateItem(c.Request.Context(), map[string]interface{}{"id": id}, &dataItem)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
