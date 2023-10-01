package handler

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

func CreateTaskFunc(ctx *gin.Context) {
	var newTask Task
	err := ctx.BindJSON(&newTask)
	if err != nil {
		return
	}
	arr = append(arr, newTask)
	ctx.IndentedJSON(201, true)
}

func GetTaskFunc(ctx *gin.Context) {
	ctx.IndentedJSON(200, &arr)
}

func UpdateTaskFunc(ctx *gin.Context) {

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

func DeleteTaskFunc(ctx *gin.Context) {
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
