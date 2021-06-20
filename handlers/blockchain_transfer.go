package handlers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo"
	"github.com/pavel-main/coinflip/core"
	"github.com/pavel-main/coinflip/models"
)

func (h *Coinflip) BlockchainInvoiceGet(c echo.Context) error {
	ctx := c.(*core.CoinflipContext)

	invoiceID := c.Param("id")
	transfer := models.Transfer{}
	if h.Database.Unscoped().Preload("Address").Where("invoice_id = ?", invoiceID).First(&transfer).RecordNotFound() {
		return ctx.JsonError(errors.New(core.ErrTransferNotFound))
	}

	return ctx.JSON(http.StatusOK, transfer)
}
