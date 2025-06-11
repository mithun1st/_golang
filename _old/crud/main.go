/*
	>> go mod init example.com/project_name1/
	>> go get "github.com/gin-gonic/gin"
*/

package main
import (
	handler "example.com/crud/pkg/http"
	"github.com/gin-gonic/gin"
)

func main() {
	
	router := gin.Default()

	router.POST("/create-task", handler.CreateTaskFunc)
	router.GET("/get-tasks", handler.GetTaskFunc)
	router.PUT("/update-task", handler.UpdateTaskFunc)
	router.DELETE("/delete-task", handler.DeleteTaskFunc)

	router.Run("127.0.0.1:2222")
}
