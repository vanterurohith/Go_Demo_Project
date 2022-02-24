package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	c.IndentedJSON(http.StatusOK, users)
}

func PostUsers(c *gin.Context) {

	var newUser user

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)

}

func GetUsersById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range users {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
