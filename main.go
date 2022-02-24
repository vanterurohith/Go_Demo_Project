package main

import (
	"github.com/gin-gonic/gin"
	"rohith.com/router"
)

func main() {

	r := gin.Default()

	r.GET("/users", router.GetUsers)
	r.POST("/register", router.PostUsers)
	r.GET("/users/:id", router.GetUsersById)
	r.Run(":8080")
}
