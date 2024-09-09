package main

import (
	"log-service/data"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type JSONPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (app *Config) WriteLog(c *fiber.Ctx) error {
	var requestPayload JSONPayload

	if err := c.BodyParser(&requestPayload); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	// insert data
	event := data.LogEntry{
		Name: requestPayload.Name,
		Data: requestPayload.Data,
	}

	err := app.Models.LogEntry.Insert(event)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	resp := jsonResponse{
		Error:   false,
		Message: "logged",
	}

	return c.Status(http.StatusAccepted).JSON(resp)
}
