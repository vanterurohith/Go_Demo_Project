package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	gorm.Model

	ID        string
	FirstName string
	LastName  string
	Email     string
}

func setup(db *gorm.DB) {
	db.AutoMigrate(&User{})
	seed(db)
}

func seed(db *gorm.DB) {
	users := []User{
		{ID: "1", FirstName: "rohith", LastName: "Vanteru", Email: "vanterurohith@gmail.com"},
		{ID: "2", FirstName: "rahul", LastName: "TK", Email: "rahultk@gmail.com"},
		{ID: "3", FirstName: "sujith", LastName: "V", Email: "sujithv@gmail.com"},
	}
	for _, u := range users {
		db.Create(&u)
	}
}
func Database() {
	dsn := "root:rohith@tcp(127.0.0.1:3306)/studentdata?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("unable to connect to database")
	}
	DB = db
	setup(db)
}
