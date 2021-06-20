package handlers

import (
	"fmt"
	"math/big"
	"net/http"

	"github.com/labstack/echo"
	"github.com/pavel-main/coinflip/core"
	"github.com/pavel-main/coinflip/payloads"
)

func (h *Coinflip) UpdatePrice(c echo.Context) error {
	ctx := c.(*core.CoinflipContext)

	// Parse payload
	payload := new(payloads.PriceUpdate)
	if err := ctx.Bind(payload); err != nil {
		return ctx.JsonError(err)
	}

	// Prepare Ethereum value
	empty := new(big.Int)
	price, ok := empty.SetString(payload.Value, 10)
	if !ok {
		message := fmt.Errorf(core.ErrPriceUpdateFailure, payload.Value)
		return ctx.JsonError(message)
	}

	// Minimal validation
	if price.Cmp(core.MinPayment) != 1 {
		message := fmt.Errorf(core.ErrPriceUpdateMinValue, payload.Value, core.MinPaymentValue)
		return ctx.JsonError(message)
	}

	// Send transaction
	transaction, err := h.SaleContract.UpdatePrice(h.TxOpts, price)
	if err != nil {
		return ctx.JsonError(err)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"txHash": transaction.Hash().String(),
	})
}
