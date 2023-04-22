package gin_item

import (
	"github.com/gin-gonic/gin"
	"go-rest-api-todo-list/common"
	"go-rest-api-todo-list/module/item/business"
	"go-rest-api-todo-list/module/item/storage"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetItemById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := storage.NewMySQLStorage(db)
		biz := business.NewGetItemBiz(store)
		dataItem, err := biz.GetItemById(c.Request.Context(), id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(dataItem))
		return
	}
}
