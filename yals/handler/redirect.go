package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h yalsHandler) RedirectFromAlias(c *fiber.Ctx) error {
	alias := c.Params("alias")
	dry := c.Query("dry", "0")

	url, err := h.Service.GetURLFromAlias(alias)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	if dry == "1" {
		return c.Status(http.StatusOK).SendString(url)
	}

	return c.Status(http.StatusPermanentRedirect).Redirect(url)
}
