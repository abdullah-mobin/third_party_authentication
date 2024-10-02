package handler

import (
	"Google_sign_option/database"
	"Google_sign_option/middleware"
	"Google_sign_option/model"
	"context"
	"database/sql"
	"encoding/json"
	"os"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type UserData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {

	var provided, registered UserData

	err := c.BodyParser(&provided)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "invalid input formate",
			"message": "Input Username & password in correct way ",
		})
	}

	query := "SELECT pass FROM user_info WHERE usr = ?"
	errr := database.DB.QueryRow(query, provided.Username).Scan(&registered.Password)
	if errr == sql.ErrNoRows {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "invalid username",
			"message": "user not exist",
		})
	} else if errr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "internal server error",
		})
	}

	if registered.Password == provided.Password {
		token, err := middleware.JWTgenerate(provided.Password)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   "internal server error",
				"message": "internal server error",
			})
		}
		return c.JSON(fiber.Map{
			"message": "login successful",
			"token":   token,
		})
	}

	return c.JSON(fiber.Map{
		"status":  "ok",
		"message": "NC",
	})
}

func oauthConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SC"),
		RedirectURL:  os.Getenv("REDIRECT_URL"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
}

func LoginWithGoogle(c *fiber.Ctx) error {
	oauthconf := oauthConfig()
	url := oauthconf.AuthCodeURL("state")

	return c.Redirect(url)
}

func GoogleCallback(c *fiber.Ctx) error {
	code := c.Query("code")
	if code == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Code not found")
	}

	oauthconf := oauthConfig()

	token, err := oauthconf.Exchange(context.Background(), code)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to exchange token: " + err.Error())
	}

	client := oauthconf.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to get user info: " + err.Error())
	}
	defer resp.Body.Close()

	var user model.GoogleUser
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to decode user info: " + err.Error())
	}

	return c.Redirect(os.Getenv("DASHBOARD"))
}

func LoginWithFacebook(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "ok",
		"message": "NC",
	})
}
