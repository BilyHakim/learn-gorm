package main

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"learn-gorm/database"
	"learn-gorm/models"
)

func main() {
	database.StartDB()

	//createUser("bilyhakim12@gmail.com")
	getUserById(1)
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

func getUserById(id uint) {
	db := database.GetDB()

	user := models.User{}

	err := db.First(&user, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("User data not found")
			return
		}
		print("Error finding user:", err)
	}
	fmt.Printf("User Data: %+v \n", user)
}
