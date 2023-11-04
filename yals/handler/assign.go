package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type putURLRequest struct {
	URL string `json:"url"`
}

type putURLResponse struct {
	URL   string `json:"url"`
	Alias string `json:"alias"`
}

func (h yalsHandler) AssignAliasToURL(c *fiber.Ctx) error {
	alias := c.Params("alias")
	force := c.Query("force", "0")

	var body putURLRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).SendString("invalid body")
	}

	isForce := false
	if force == "1" {
		isForce = true
	}

	err := h.Service.SetURLToAlias(alias, body.URL, isForce)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(&putURLResponse{
		URL:   body.URL,
		Alias: alias,
	})
}
