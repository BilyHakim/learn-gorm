package main

import (
	"fmt"
	"learn-gorm/database"
	"learn-gorm/models"
)

func main() {
	database.StartDB()

	createUser("bilyhakim12@gmail.com")
}

func createUser(email string) {
	db := database.GetDB()

	User := models.User{
		Email: email,
	}

	err := db.Create(&User).Error

	if err != nil {
		fmt.Println("Error creating user data:", err)
		return
	}

	fmt.Println("New User Data:", User)
}
