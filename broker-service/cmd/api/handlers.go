package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"errors"
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

	if err := c.BodyParser(&a); err != nil {
		return c.Status(http.StatusUnauthorized).JSON(err.Error())
	}

	var encodePayload bytes.Buffer

	enc := gob.NewEncoder(&encodePayload)
	_ = enc.Encode(a)

	request, err := http.NewRequest("POST", "http://authentication-service/authticate", bytes.NewBuffer(encodePayload.Bytes()))
	if err != nil {
		return err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		return errors.New("invalid credentials")
	} else if response.StatusCode != http.StatusAccepted {
		return errors.New("error calling auth service")
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
	return nil
}
