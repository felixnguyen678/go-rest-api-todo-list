package main

import (
	"github.com/gin-gonic/gin"
	gin_item "go-rest-api-todo-list/module/item/transport/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

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
		v1.POST("/items", gin_item.CreateItem(db))           // create item
		v1.GET("/items", gin_item.GetListOfItems(db))        // list items
		v1.GET("/items/:id", gin_item.GetItemById(db))       // get an item by ID
		v1.PUT("/items/:id", gin_item.UpdateItemById(db))    // edit an item by ID
		v1.DELETE("/items/:id", gin_item.DeleteItemById(db)) // delete an item by ID
	}

	router.Run()
}
