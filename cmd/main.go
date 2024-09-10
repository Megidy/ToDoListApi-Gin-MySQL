package main

import (
	"github.com/Megidy/To-Do-List-Api/pkj/config"
	"github.com/Megidy/To-Do-List-Api/pkj/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	config.Connect()
	routes.InitRouter(router)

	router.Run(":8080")
}
