package routers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"rohith.com/Routers/database"
)

type user struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Department string `json:"department"`
}

var users = []user{
	{ID: "1", Name: "Rohith", Email: "vanterurohith@gmail.com", Department: "CSE"},
	{ID: "2", Name: "Rahul", Email: "rahul@gmail.com", Department: "IT"},
	{ID: "3", Name: "Sujith", Email: "sujith@gmail.com", Department: "CSE"},
	{ID: "4", Name: "Kaushik", Email: "kaushik@gmail.com", Department: "ECE"},
}

func GetUsers(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
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
