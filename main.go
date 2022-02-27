package main

import (
	"github.com/gin-gonic/gin"
	routers "rohith.com/Routers"
	"rohith.com/Routers/database"
)

func main() {

	r := gin.Default()

	r.GET("/users", routers.GetUsers)
	r.POST("/register", routers.PostUsers)
	r.GET("/users/:id", routers.GetUsersById)
	database.Database()
	r.Run(":8080")
}
