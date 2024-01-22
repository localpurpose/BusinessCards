package controller

import (
	"encoding/base64"
	db "github.com/buscard/config"
	models "github.com/buscard/models"
	"github.com/buscard/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"os"
	"strconv"
	"strings"
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
		Description:  data["desc"],
		Organization: data["org"],
		PhoneNumber:  data["phone_number"],
		Email:        data["email"],
		WebSite:      data["website"],
		TelegramName: data["tg_name"],
		WhatsApp:     data["whatsApp"],
		Theme:        data["theme"],
		CreatedAt:    time.Now().String(),
		LastUpdate:   time.Now().String(),
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
	user.LastUpdate = time.Now().String()
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

func InsertTheme(c *fiber.Ctx) error {

	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Invalid data",
			})
	}

	if data["user_id"] == "" {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Invalid data. User ID is required.",
			})
	}
	if data["theme"] == "" {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Invalid data. User Theme is required.",
			})
	}

	user_id, err := strconv.ParseInt(data["user_id"], 10, 64)
	if err != nil {
		log.Info(err)
	}

	theme := models.User{
		Id:    uint(user_id),
		Theme: data["theme"],
	}

	db.DB.Create(&theme)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Theme Details",
		"data":    theme,
	})
}

func RenderUserProfile(c *fiber.Ctx) error {
	userId := c.Params("userid")
	var user models.User
	var path string

	db.DB.Select("id, name, telegram_name, organization, phone_number, email, web_site, whats_app, description, qr_path, theme").Where("id = ?", userId).First(&user)
	log.Info(user.Id)
	if user.Id == 0 {
		err := c.Redirect("/1")
		if err != nil {
			log.Info(err)
		}
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "User not found",
			"error":   map[string]interface{}{},
		})
	}

	path = "usersData/" + userId
	if !utils.PathExists(path) {
		err := os.Mkdir(path, 0755)
		utils.QrGenerate(user.Theme, "http://localhost/"+strconv.FormatUint(uint64(user.Id), 10), path)
		if err != nil {
			log.Info(err)
		}

		user.QrPath = strings.Split(path, "/")[1]
		db.DB.Save(&user)
	} else {
		os.Remove(path + "/qr.png")
		utils.QrGenerate(user.Theme, "http://localhost/"+strconv.FormatUint(uint64(user.Id), 10), path)

		user.QrPath = strings.Split(path, "/")[1]
		db.DB.Save(&user)
	}
	log.Info(user.QrPath)
	userDataMap := fiber.Map{
		"Title":        user.Name,
		"Name":         user.Name,
		"Desk":         user.Description,
		"Org":          user.Organization,
		"PhoneNumber":  user.PhoneNumber,
		"EmailAddress": user.Email,
		"WebSite":      user.WebSite,
		"TgName":       user.TelegramName,
		"WhatsApp":     user.WhatsApp,
		"ImageDir":     user.QrPath,
	}

	if utils.PathExists(path + "/image.png") {
		userDataMap["ImageDir1"] = user.QrPath + "/image.png"
	}

	if user.Theme == "brown" {
		return c.Render("brownView/index", userDataMap)
	}
	if user.Theme == "blue" {
		return c.Render("blueView/index", userDataMap)
	}
	if user.Theme == "orange" {
		return c.Render("orangeView/index", userDataMap)
	}
	if user.Theme == "pink" {
		return c.Render("pinkView/index", userDataMap)
	}

	return nil
}

func UploadImage(c *fiber.Ctx) error {

	var data map[string]string
	var userPath string

	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Invalid data",
			})
	}

	if data["user_id"] == "" {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Invalid data. User id is required.",
			})
	} else {
		userPath = "./usersData/" + data["user_id"]
	}
	if data["image"] == "" {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Invalid data. Image required.",
			})
	}

	// Save b64 from data and making image
	decoded, err := base64.StdEncoding.DecodeString(data["image"])
	if err != nil {
		log.Fatal(err)
	}

	// Write the image to a file
	//Check if file exists is not necessary, if file exists it will be overwritten!!!
	file, err := os.Create(userPath + "/image.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.Write(decoded)
	if err != nil {
		log.Fatal(err)
	}

	return c.Status(404).JSON(fiber.Map{
		"success": true,
		"message": "Image upload successfully!",
		"error":   map[string]interface{}{},
	})
}
