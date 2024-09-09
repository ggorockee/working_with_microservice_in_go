package main

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omitempty"`
}

func Broker(c *fiber.Ctx) error {
	payload := &jsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	//if err := c.BodyParser(payload); err != nil {
	//	return c.Status(http.StatusBadRequest).JSON(fiber.Map{
	//		"status": "failed",
	//		"error": err.Error(),
	//	})
	//}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": "ok",
		"data":   payload,
	})
}

func HealthCheck(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": "ok",
		"data":   nil,
	})
}

func HandleSubmission(c *fiber.Ctx) error {
	var requestPayload RequestPayload

	if err := c.BodyParser(&requestPayload); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	switch requestPayload.Action {
	case "auth":
		return authenticate(c, requestPayload.Auth)
	default:
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "unknown action",
		})
	}
}

func authenticate(c *fiber.Ctx, a AuthPayload) error {

	//a AuthPayload = {
	//	Email:    "sample@sample.com",
	//	Password: "password",
	//}

	jsonData, _ := json.MarshalIndent(a, "", "  ")

	request, err := http.NewRequest("POST", "http://authentication-service/authenticate", bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json")

	if err != nil {
		return err
	}

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	} else if response.StatusCode == http.StatusUnauthorized {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "authentication failed",
		})
	}

	var jsonFromService jsonResponse
	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		return err
	}

	if jsonFromService.Error {
		return err
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "authenticated!"
	payload.Data = jsonFromService.Data

	return c.Status(http.StatusOK).JSON(payload)
}
