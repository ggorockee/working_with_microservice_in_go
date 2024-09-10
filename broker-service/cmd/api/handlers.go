// mongodb uri: // mongodb://admin:password@localhost:27017/logs?authSource=admin&readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false

package main

import (
	"bytes"
	"encoding/json"
	"github.com/ggorockee/working_with_microservice_in_go/broker-service/event"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type MailPayload struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omitempty"`
	Log    LogPayload  `json:"log,omitempty"`
	Mail   MailPayload `json:"mail,omitempty"`
}

type LogPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) Broker(c *fiber.Ctx) error {
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

func (app *Config) HealthCheck(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": "ok",
		"data":   nil,
	})
}

func (app *Config) HandleSubmission(c *fiber.Ctx) error {
	var requestPayload RequestPayload

	if err := c.BodyParser(&requestPayload); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	switch requestPayload.Action {
	case "auth":
		return app.authenticate(c, requestPayload.Auth)
	case "log":
		return app.logEventViaRabbit(c, requestPayload.Log)
	case "mail":
		return app.sendMail(c, requestPayload.Mail)
	default:
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "unknown action",
		})
	}
}

func (app *Config) authenticate(c *fiber.Ctx, a AuthPayload) error {

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

func (app *Config) logItem(c *fiber.Ctx, entry LogPayload) error {
	jsonData, _ := json.MarshalIndent(entry, "", "\t")
	logServiceURL := "http://logger-service/log"

	request, err := http.NewRequest("POST", logServiceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusAccepted {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "bad request!!",
		})
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "logged"
	return c.Status(http.StatusAccepted).JSON(payload)

}

func (app *Config) sendMail(c *fiber.Ctx, msg MailPayload) error {

	jsonData, _ := json.MarshalIndent(msg, "", "\t")

	mailServiceURL := "http://mail-service/send"

	request, err := http.NewRequest("POST", mailServiceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		jsonResp := jsonResponse{
			Error:   true,
			Message: err.Error(),
		}
		return c.Status(http.StatusBadRequest).JSON(jsonResp)
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		jsonResp := jsonResponse{
			Error:   true,
			Message: err.Error(),
		}
		return c.Status(http.StatusBadRequest).JSON(jsonResp)
	}
	defer response.Body.Close()

	errorbodyWithByte, _ := io.ReadAll(response.Body)

	if response.StatusCode != http.StatusAccepted {
		jsonResp := jsonResponse{
			Error:   true,
			Message: string(errorbodyWithByte),
		}
		return c.Status(http.StatusBadRequest).JSON(jsonResp)
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Message sent to" + msg.To,
	}

	return c.Status(http.StatusAccepted).JSON(payload)

}

func (app *Config) logEventViaRabbit(c *fiber.Ctx, l LogPayload) error {
	err := app.pushToQueue(l.Name, l.Data)
	if err != nil {
		errResp := jsonResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(http.StatusBadRequest).JSON(errResp)
	}

	payload := jsonResponse{
		Error:   false,
		Message: "logged via RabbitMQ",
		Data:    nil,
	}
	return c.Status(http.StatusAccepted).JSON(payload)
}

func (app *Config) pushToQueue(name, msg string) error {
	emitter, err := event.NewEventEmitter(app.Rabbit)
	if err != nil {
		return nil
	}

	payload := LogPayload{
		Name: name,
		Data: msg,
	}

	j, _ := json.MarshalIndent(&payload, "", "\t")
	err = emitter.Push(string(j), "log.INFO")
	if err != nil {
		return err
	}
	return nil
}
