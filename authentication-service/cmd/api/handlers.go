package main

import (
	"fmt"
	"github.com/ggorockee/working_with_microservice_in_go/authentication-service/data"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func Authenticate(c *fiber.Ctx) error {

	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&requestPayload); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	var models data.Models

	user, err := models.User.GetByEmail(requestPayload.Email)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in user %s", user.Email),
		Data:    user,
	}

	return c.Status(http.StatusOK).JSON(payload)
}
