package controllers

import "github.com/gofiber/fiber"

// GetNewAccessToken method for creating new access token.
// @Description Create a new access token.
// @Summary creates a new access token
// @Tags Token
// @Accept json
// @Produce json
// @Success 200 {string} status "ok"
// @Router /v1/token/new [get]
func GetNewAccessToken(c *fiber.Ctx) error {
	Generate
	a
	new
	Access
	token.
		token, err := utils.Ge
}
