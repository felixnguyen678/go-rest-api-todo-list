package gin_item

import (
	"github.com/gin-gonic/gin"
	"go-rest-api-todo-list/common"
	"go-rest-api-todo-list/module/item/business"
	"go-rest-api-todo-list/module/item/storage"
	"gorm.io/gorm"
	"net/http"
)

func GetListOfItems(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		paging.Process()

		store := storage.NewMySQLStorage(db)
		biz := business.NewListToDoItemBiz(store)

		result, err := biz.ListItems(c.Request.Context(), nil, &paging)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(
			result, paging, nil,
		))
		return
	}
}
