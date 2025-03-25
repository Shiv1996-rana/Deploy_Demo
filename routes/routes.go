package routes

import (
	Controllers "deploy_demo/Controller"
	"fmt"

	"github.com/gin-gonic/gin"
)

func SetHandlers() *gin.Engine {
	router := gin.Default()
	router.GET("test", Controllers.TestCode)
	router.POST("/newData", Controllers.CreateData)
	router.GET("/allGet", Controllers.FetchAllData)
	router.GET("/allGet/:id", Controllers.GetAllById)
	router.DELETE("/delete/:id", Controllers.DeleteDataById)
	router.PUT("/update/:id", Controllers.UpdateData)
	fmt.Println("runnning on server....")
	router.Run(":1003")
	return router
}
