package routes

import (
	"github.com/Megidy/To-Do-List-Api/pkj/controllers"
	"github.com/Megidy/To-Do-List-Api/pkj/middleware"
	"github.com/gin-gonic/gin"
)

var InitRouter = func(router gin.IRouter) {
	//user
	router.POST("/signup", controllers.SignUp)
	router.POST("/login", controllers.LogIn)
	//tasks
	router.POST("/todos", middleware.RequierAuth, controllers.CreateTask)
	router.GET("/todos", middleware.RequierAuth, controllers.GetAllTasks)
	router.GET("/todos/:taskId", middleware.RequierAuth, controllers.GetTaskById)
	router.DELETE("/todos/:taskId", middleware.RequierAuth, controllers.DeleteTask)
	router.PUT("/todos/:taskId", middleware.RequierAuth, controllers.UpdateTask)
}
