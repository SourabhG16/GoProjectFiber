package routes

import (
	"errors"
	"fmt"
	"time"

	"github.com/SourabhG16/goProject1/database"
	"github.com/SourabhG16/goProject1/database/models"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Firstname string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateResponseUser(userModel models.User) User {
	return User{ID: userModel.ID, Firstname: userModel.Firstname, LastName: userModel.LastName}
}

func CreateUser(c *fiber.Ctx) error {
	startTime := time.Now()
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Create(&user)
	responseUser := CreateResponseUser(user)
	fmt.Printf("Time elapsed in CreateUser Function is %v", time.Since(startTime))
	return c.Status(200).JSON(responseUser)
}

func GetUsers(c *fiber.Ctx) error {
	startTime := time.Now()
	users := []models.User{}
	responseUsers := []User{}
	database.Database.Db.Find(&users)
	for _, user := range users {
		responseUsers = append(responseUsers, CreateResponseUser(user))
	}
	fmt.Printf("Time elapsed in GetUsers Function is %v", time.Since(startTime))
	return c.Status(200).JSON(responseUsers)
}

func fetchUser(id int, user *models.User) error {
	database.Database.Db.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return errors.New("User does not exist")
	}
	return nil
}

func FindUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(404).JSON("Please enter the id")
	}
	var user models.User
	if err := fetchUser(id, &user); err != nil {
		return c.Status(404).JSON("No user found")
	}
	responseUser := CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)

}

func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(404).JSON("Please enter the id")
	}
	var user models.User
	if err := fetchUser(id, &user); err != nil {
		return c.Status(404).JSON("No user found")
	}

	type UpdateUser struct {
		Firstname string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	var updateData UpdateUser
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(404).JSON("Enter correct information")
	}
	user.Firstname = updateData.Firstname
	user.LastName = updateData.LastName
	database.Database.Db.Save(&user)

	responseUser := CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)

}
