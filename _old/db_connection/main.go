package main

import (
	"fmt"

	"example.com/db_connection/pkg/database"
	"example.com/db_connection/pkg/handler"
	"github.com/gin-gonic/gin"
)

func main() {

	//init db
	fmt.Println(database.InitDB())

	//define handler
	router := gin.Default()
	router.GET("/", handler.TestFunc)
	router.POST("/create-task", handler.CreateTaskFunc)
	router.GET("/get-tasks", handler.GetTaskFunc)
	router.PUT("/update-task", handler.UpdateTaskFunc)
	router.DELETE("/delete-task", handler.DeleteTaskFunc)

	//init server
	router.Run("0.0.0.0:2222")
}

/*

POST:
127.0.0.1:2222/create-task
{
    "id": "101",
    "title": "morning coffee",
    "isDone" : false
}

GET:
127.0.0.1:2222/get-tasks

UPDATE:
127.0.0.1:2222/update-task?id=101
{
    "id": "101",
    "title": "morning coffee",
    "isDone" : true
}

DELETE:
127.0.0.1:2222/delete-task?id=101


DB	"mysql", "root:admin@tcp(127.0.0.1:8080)/mydb"
table name- task

*/
