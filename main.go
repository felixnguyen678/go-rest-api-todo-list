package main

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-rest-api-todo-list/common"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

//enum area

type ItemStatus int

const (
	ItemStatusDoing ItemStatus = iota
	ItemStatusDone
	ItemStatusDeleted
)

var itemStatusStringList = [3]string{"doing", "done", "deleted"}

func (item ItemStatus) String() string {
	return fmt.Sprintf(itemStatusStringList[item])
}

func parseStr2ItemStatus(s string) (ItemStatus, error) {
	for i := range itemStatusStringList {
		if itemStatusStringList[i] == s {
			return ItemStatus(i), nil
		}
	}
	return ItemStatus(0), errors.New(fmt.Sprintf("Invalid status string"))
}

func (item *ItemStatus) Scan(value interface{}) error {
	bytes, ok := value.([]byte)

	if !ok {
		return errors.New(fmt.Sprintf("There is an error at %s", value))
	}

	strValue := string(bytes)
	v, error := parseStr2ItemStatus(strValue)

	if error != nil {
		return errors.New(fmt.Sprintf("There is an error at %s", value))
	}
	*item = v

	return nil
}

func (item *ItemStatus) Value() (driver.Value, error) {
	if item == nil {
		return nil, nil
	}

	return item.String(), nil
}

func (item *ItemStatus) MarshalJSON() ([]byte, error) {
	if item == nil {
		return nil, nil
	}
	return []byte(fmt.Sprintf("\"%s\"", item.String())), nil
}

func (item *ItemStatus) UnMarshalJSON(data []byte) error {
	str := strings.ReplaceAll(string(data), "\"", "")
	itemValue, err := parseStr2ItemStatus(str)

	if err != nil {
		return err
	}

	*item = itemValue

	return nil
}

//end enum area

type ToDoItem struct {
	common.SQLModel
	Title  string      `json:"title" gorm:"column:title;"`
	Status *ItemStatus `json:"status" gorm:"column:status;"`
}

func (ToDoItem) TableName() string { return "todo_items" }

type ToDoItemCreation struct {
	Id    int    `json:"-" gorm:"column:id;"` // return id of todoItem after creating
	Title string `json:"title" gorm:"column:title;"`
}

func (ToDoItemCreation) TableName() string { return ToDoItem{}.TableName() }

type ToDoItemUpdate struct {
	Title  *string `json:"title" gorm:"column:title;"`
	Status *string `json:"status" gorm:"column:status;"`
}

func (ToDoItemUpdate) TableName() string { return ToDoItem{}.TableName() }

func main() {
	dsn := os.Getenv("DB_STRING")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Cannot connect to MySQL:", err)
	}

	log.Println("Connected to MySQL:", db)

	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.POST("/items", createItem(db))           // create item
		v1.GET("/items", getListOfItems(db))        // list items
		v1.GET("/items/:id", readItemById(db))      // get an item by ID
		v1.PUT("/items/:id", editItemById(db))      // edit an item by ID
		v1.DELETE("/items/:id", deleteItemById(db)) // delete an item by ID
	}

	router.Run()
}

func createItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dataItem ToDoItemCreation

		if err := c.ShouldBind(&dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//// preprocess title - trim all spaces
		//dataItem.Title = strings.TrimSpace(dataItem.Title)
		//
		//if dataItem.Title == "" {
		//	c.JSON(http.StatusBadRequest, gin.H{"error": "title cannot be blank"})
		//	return
		//}
		//
		//// do not allow "finished" status when creating a new task
		//dataItem.Status = "Doing" // set to default
		//

		if err := db.Create(&dataItem).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(dataItem.Id))
	}
}

func readItemById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dataItem ToDoItem

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Where("id = ?", id).First(&dataItem).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(dataItem))
	}
}

func getListOfItems(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		paging.Process()

		offset := (paging.Page - 1) * paging.Limit

		var result []ToDoItem

		if err := db.Table(ToDoItem{}.TableName()).
			Count(&paging.Total).
			Offset(offset).
			Order("id desc").
			Find(&result).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(
			result, paging, nil,
		))
	}
}

func editItemById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var dataItem ToDoItemUpdate

		if err := c.ShouldBind(&dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Where("id = ?", id).Updates(&dataItem).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

func deleteItemById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Table(ToDoItem{}.TableName()).
			Where("id = ?", id).
			Delete(nil).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
