package handler

import (
	"example.com/db_connection/pkg/database"
	"example.com/db_connection/pkg/model"
	"github.com/gin-gonic/gin"
)

func TestFunc(ctx *gin.Context) {
	ctx.IndentedJSON(201, "Hello World")
}

func CreateTaskFunc(ctx *gin.Context) {
	var newTask model.Task
	err := ctx.BindJSON(&newTask)
	if err != nil {
		return
	}
	// arr = append(arr, newTask)
	reId, err := database.CreateTaskInDB(newTask)
	if err != nil {

		ctx.IndentedJSON(400, err.Error())
		return
	}

	ctx.IndentedJSON(201, reId)
}

func GetTaskFunc(ctx *gin.Context) {

	list, err := database.GetTasksInDB()
	if err != nil {
		ctx.IndentedJSON(400, err.Error())
		return
	}

	ctx.IndentedJSON(200, &list)
}

func UpdateTaskFunc(ctx *gin.Context) {

	var id string = ctx.Query("id")

	var newTask model.Task
	err := ctx.BindJSON(&newTask)
	if err != nil {
		return
	}

	lastId, err := database.UpdateTaskInDB(id, newTask)
	if err != nil {
		ctx.IndentedJSON(400, err.Error())
	}

	ctx.IndentedJSON(203, lastId)
}

func DeleteTaskFunc(ctx *gin.Context) {
	var id string = ctx.Query("id")

	lastId, err := database.DeleteTaskInDB(id)
	if err != nil {
		ctx.IndentedJSON(400, err.Error())
		return
	}

	ctx.IndentedJSON(202, lastId)
}
