package controller

import (
	db "github.com/buscard/config"
	models "github.com/buscard/models"
	"github.com/gofiber/fiber/v2"
	"time"
)

func CreateUser(c *fiber.Ctx) error {

	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Invalid data",
			})
	}

	if data["name"] == "" {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Invalid data. User name is required.",
			})
	}
	if data["tg_name"] == "" {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Invalid data. User Telegram Name is required.",
			})
	}
	// Save User in the Database
	user := models.User{
		Name:         data["name"],
		TelegramName: data["tg_name"],
		CreatedAt:    time.Now(),
		LastUpdate:   time.Now(),
	}

	db.DB.Create(&user)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "User created!",
		"data":    user,
	})

}
func EditUser(c *fiber.Ctx) error {

	userId := c.Params("userid")
	var user models.User
	db.DB.Find(&user, "id = ?", userId)

	if user.Name == "" {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "User not found",
		})
	}

	var updateUser models.User
	err := c.BodyParser(&updateUser)
	if err != nil {
		return err
	}

	if updateUser.Name == "" {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "User name would not be empty!",
		})
	}

	user.Name = updateUser.Name
	user.LastUpdate = time.Now()
	db.DB.Save(&updateUser)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "User details updated!",
		"data":    user,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	userId := c.Params("userid")
	var user models.User
	db.DB.Where("id = ?", userId).First(&user)

	if user.Id == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "User not found!",
		})
	}
	db.DB.Where("id = ?", userId).Delete(&user)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "User deleted successfully!",
	})
}
func GetDetails(c *fiber.Ctx) error {

	userId := c.Params("userid")
	var user models.User

	db.DB.Select("id, name, created_at, last_update, telegram_name").Where("id = ?", userId).First(&user)

	userData := make(map[string]interface{})
	userData["id"] = user.Id
	userData["name"] = user.Name
	userData["telegram_name"] = user.TelegramName
	userData["created_at"] = user.CreatedAt
	userData["last_update"] = user.LastUpdate

	if user.Id == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "User not found",
			"error":   map[string]interface{}{},
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "User Details",
		"data":    userData,
	})
}
