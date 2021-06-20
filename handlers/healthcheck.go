package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/pavel-main/coinflip/core"
)

func (h *Coinflip) Healthcheck(c echo.Context) error {
	ctx := c.(*core.CoinflipContext)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"status":   core.HealthcheckResponse,
		"features": h.Config.Features,
	})
}
