package handlers

import (
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo"
	"github.com/pavel-main/coinflip/core"
	"github.com/pavel-main/coinflip/payloads"
	"github.com/pavel-main/coinflip/responses"
)

func (h *Coinflip) WhitelistPost(c echo.Context) error {
	ctx := c.(*core.CoinflipContext)

	// Parse payload
	payload := new(payloads.WhitelistAdd)
	if err := ctx.Bind(payload); err != nil {
		return ctx.JsonError(err)
	}

	// Convert raw input to Ethereum addresses
	addresses := []common.Address{}
	for _, address := range payload.Addresses {
		if common.IsHexAddress(address) {
			addresses = append(addresses, common.HexToAddress(address))
		} else {
			err := fmt.Errorf(core.ErrInvalidEthereumAddress, address)
			return ctx.JsonError(err)
		}
	}

	// Send transaction
	transaction, err := h.SaleContract.WhitelistAdd(h.TxOpts, addresses)
	if err != nil {
		return ctx.JsonError(err)
	}

	// Return response
	return ctx.JSON(http.StatusOK, responses.WhitelistPost{
		TxHash: transaction.Hash().Hex(),
	})
}
