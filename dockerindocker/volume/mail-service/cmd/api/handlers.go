package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (app *Config) SendMail(c *fiber.Ctx) error {
	type mailMessage struct {
		From    string `json:"from"`
		To      string `json:"to"`
		Subject string `json:"subject"`
		Message string `json:"message"`
	}

	var requestPayload mailMessage

	if err := c.BodyParser(&requestPayload); err != nil {
		errorResp := errorResponse{
			Error:   true,
			Message: err.Error(),
		}
		return c.Status(http.StatusBadRequest).JSON(errorResp)
	}

	msg := Message{
		From:    requestPayload.From,
		To:      requestPayload.To,
		Subject: requestPayload.Subject,
		Data:    requestPayload.Message,
	}

	err := app.Mailer.SendSMTPMessage(msg)
	if err != nil {
		errorResp := errorResponse{
			Error:   true,
			Message: err.Error(),
		}
		return c.Status(http.StatusBadRequest).JSON(errorResp)
	}

	payload := jsonResponse{
		Error:   false,
		Message: "sent to" + requestPayload.To,
	}

	return c.Status(http.StatusAccepted).JSON(payload)

}
