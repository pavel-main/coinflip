package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/pavel-main/coinflip/core"
	"github.com/pavel-main/coinflip/responses"
)

func (h *Coinflip) StatsGet(c echo.Context) error {
	ctx := c.(*core.CoinflipContext)

	// Retrieve data from smart contract
	active, err := h.SaleContract.IsActiveSale(nil)
	availableUnits, err := h.SaleContract.AvailableUnits(nil)
	availableBonus, err := h.SaleContract.AvailableBonus(nil)
	bonusUsed, err := h.SaleContract.BonusUsed(nil)
	duration, err := h.SaleContract.Duration(nil)
	minPayment, err := h.SaleContract.MinPayment(nil)
	price, err := h.SaleContract.Price(nil)
	startTime, err := h.SaleContract.StartTime(nil)
	unitsSold, err := h.SaleContract.UnitsSold(nil)
	weiReceived, err := h.SaleContract.WeiReceived(nil)

	if err != nil {
		return ctx.JsonError(err)
	}

	return ctx.JSON(http.StatusOK, responses.Stats{
		Active:         active,
		AvailableUnits: availableUnits.Uint64(),
		AvailableBonus: availableBonus.Uint64(),
		BonusUsed:      bonusUsed.Uint64(),
		Duration:       duration.Uint64(),
		MinPayment:     minPayment.Uint64(),
		Price:          price.Uint64(),
		StartTime:      startTime.Uint64(),
		UnitsSold:      unitsSold.Uint64(),
		WeiReceived:    weiReceived.Uint64(),
	})
}
