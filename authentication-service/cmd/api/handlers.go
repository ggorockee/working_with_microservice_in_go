package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ggorockee/working_with_microservice_in_go/authentication-service/data"
	"github.com/gofiber/fiber/v2"
)

func (app *Config) Authenticate(c *fiber.Ctx) error {

	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&requestPayload); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	var user data.User
	_, err := user.GetByEmail(requestPayload.Email)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	// log authentication
	err = app.logRequest("authentication", fmt.Sprintf("%s logged in", user.Email))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in user %s", user.Email),
		Data:    user,
	}

	return c.Status(http.StatusOK).JSON(payload)
}

func (app *Config) logRequest(name, data string) error {
	var entry struct {
		Name string `json:"name"`
		Data string `json:"data"`
	}

	entry.Name = name
	entry.Data = data

	jsonData, _ := json.MarshalIndent(entry, "", "\t")
	logServiceURL := "http://logger-service/log"

	request, err := http.NewRequest("POST", logServiceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	client := &http.Client{}
	_, err = client.Do(request)

	if err != nil {
		return err
	}

	return nil
}

func (app *Config) HealthCheck(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"health": "ok",
		"data":   nil,
	})
}
