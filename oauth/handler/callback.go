package handler

import (
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
)

func (h oauthHandler) Callback(c *fiber.Ctx) error {
	code := c.Query("code")
	clientID := os.Getenv("GITHUB_CLIENT_ID")
	clientSecret := os.Getenv("GITHUB_CLIENT_SECRET")

	if clientID == "" || clientSecret == "" {
		return c.Status(http.StatusInternalServerError).SendString("Missing GITHUB_CLIENT_ID or GITHUB_CLIENT_SECRET")
	}

	accessToken, err := h.OAuthService.GetAccessToken(clientID, clientSecret, code)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(http.StatusOK).SendString(accessToken)
}
