package main

import (
	"github.com/diegoparra/check-birthday/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/get-user", controllers.GetUsers)

	r.Run(":8089")
}
