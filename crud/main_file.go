/*
	>> go mod init example.com/project_name1/
	>> go get "github.com/gin-gonic/gin"

*/

package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Task struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	IsDone bool   `json:"isDone"`
}

var arr []Task = []Task{
	{Id: 1, Title: "Wakeup Early", IsDone: true},
	{Id: 2, Title: "Make Coffee", IsDone: true},
}

func main() {

	router := gin.Default()

	router.POST("/create-task", createTaskFunc)
	router.GET("/get-tasks", getTaskFunc)
	router.PUT("/update-task", updateTaskFunc)
	router.DELETE("/delete-task", deleteTaskFunc)

	router.Run("127.0.0.1:2020")
}

//function

func createTaskFunc(ctx *gin.Context) {
	var newTask Task
	err := ctx.BindJSON(&newTask)
	if err != nil {
		return
	}
	arr = append(arr, newTask)
	ctx.IndentedJSON(201, true)
}

func getTaskFunc(ctx *gin.Context) {
	ctx.IndentedJSON(200, &arr)
}

func updateTaskFunc(ctx *gin.Context) {

	var id string = ctx.Query("id")
	var isExist bool = false
	for i, v := range arr {
		if fmt.Sprint(v.Id) == id {
			arr = append(arr[:i], arr[i+1:]...)
			isExist = true
		}
	}
	if !isExist {
		ctx.IndentedJSON(404, "no id found")
		return
	}

	var newTask Task
	err := ctx.BindJSON(&newTask)
	if err != nil {
		return
	}
	arr = append(arr, newTask)
	ctx.IndentedJSON(203, true)
}

func deleteTaskFunc(ctx *gin.Context) {
	var id string = ctx.Query("id")
	for i, v := range arr {
		if fmt.Sprint(v.Id) == id {
			arr = append(arr[:i], arr[i+1:]...)
			ctx.IndentedJSON(203, true)
			return
		}
	}
	ctx.IndentedJSON(404, "no id found for delete")
}
