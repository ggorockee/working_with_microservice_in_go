package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type ResponseHTTP struct {
	Success bool   `json:"success"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

// HealthCheck is a function to get health condition
// @Summary Get Health Check
// @Description Get Health Check
// @Tags HealthCheck
// @Accept json
// @Produce json
// @Success 200 {object} ResponseHTTP{data=any}
// @Failure 503 {object} ResponseHTTP{}
// @Router /healthz/ready [get]
func HealthCheck(c *fiber.Ctx) error {
	responseHTTP := ResponseHTTP{
		Success: true,
		Data:    nil,
		Message: "Welcome",
	}
	return c.Status(http.StatusOK).JSON(responseHTTP)
}
