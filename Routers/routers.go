package routers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"rohith.com/Routers/database"
)

type user struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

var users = []user{
	{ID: "1", FirstName: "Rohith", LastName: "Vanteru", Email: "vanterurohith@gmail.com"},
	{ID: "2", FirstName: "Rahul", LastName: "TK", Email: "rahul@gmail.com"},
	{ID: "3", FirstName: "Sujith", LastName: "V", Email: "sujith@gmail.com"},
	{ID: "4", FirstName: "Kaushik", LastName: "T", Email: "kaushik@gmail.com"},
}

func GetUsers(c *gin.Context) {

	var newuserInfo []user
	database.DB.Find(&newuserInfo)
	c.IndentedJSON(http.StatusOK, newuserInfo)
}

func PostUsers(c *gin.Context) {

	var newUser user

	if err := c.BindJSON(&newUser); err == nil {
		fmt.Printf("obj", newUser)
	} else {
		fmt.Printf("error here", err)
	}
	database.DB.Create(&newUser)
	c.IndentedJSON(http.StatusCreated, newUser)

}

func GetUsersById(c *gin.Context) {
	id := c.Param("id")

	var newUserId user
	idresult := database.DB.Where("id=?", id).First(&newUserId)
	if idresult.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
	} else {
		c.IndentedJSON(http.StatusOK, newUserId)
	}

}
