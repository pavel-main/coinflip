package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	httpclient "github.com/ddliu/go-httpclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hashicorp/go-uuid"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/pavel-main/coinflip/core"
	"github.com/pavel-main/coinflip/models"
	"github.com/pavel-main/coinflip/payloads"
	"github.com/pavel-main/coinflip/responses"
)

func (h *Coinflip) BlockchainReceive(c echo.Context) error {
	ctx := c.(*core.CoinflipContext)

	// Find available account
	account := models.Account{}
	if h.Database.Where("gap < ?", core.Bip44AddressLimit).First(&account).RecordNotFound() {
		return ctx.JsonError(errors.New(core.ErrNoAvailableAccountsFound))
	}

	// Parse payload
	payload := new(payloads.ReceiveAddress)
	if err := ctx.Bind(payload); err != nil {
		return ctx.JsonError(err)
	}

	// Validate beneficiary address
	if !common.IsHexAddress(payload.Beneficiary) {
		return ctx.JsonError(fmt.Errorf(core.ErrInvalidEthereumAddress, payload.Beneficiary))
	}

	// Check if beneficiary already has pending transfer
	transfer := models.Transfer{}
	if !h.Database.Preload("Address").Where("beneficiary = ?", payload.Beneficiary).First(&transfer).RecordNotFound() {
		if transfer.Address == nil {
			return ctx.JsonError(errors.New(core.ErrUnknown))
		}
		return ctx.JSON(http.StatusOK, transfer.Address)
	}

	// Generate invoice ID
	invoiceID, err := uuid.GenerateUUID()
	if err != nil {
		return ctx.JsonError(err)
	}

	// Send API request
	requestUrl := core.BlockchainInfoBaseUrl + core.BlockchainInfoReceive
	res, err := httpclient.Get(requestUrl, map[string]string{
		"key":      h.Config.BlockchainInfoApiKey,
		"xpub":     account.Xpub,
		"callback": core.GetCallbackUrl(h.Config.Protocol, h.Config.Domain, invoiceID),
	})

	if err != nil {
		return ctx.JsonError(err)
	}

	// Read response body
	bodyBytes, err := res.ReadAll()
	if err != nil {
		return ctx.JsonError(err)
	}

	// Unmarshal response
	response := responses.BlockchainInfoReceive{}
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return ctx.JsonError(err)
	}

	// Check for errors
	if response.Message != nil {
		return ctx.JsonError(errors.New(*response.Message))
	}

	// Begin transaction
	tx := h.Database.Begin()

	// Create new address
	address := models.Address{
		Address:   response.Address,
		AccountID: account.ID,
	}

	// Save new address
	if err := h.Database.Create(&address).Error; err != nil {
		tx.Rollback()
		return ctx.JsonError(err)
	}

	// Create new transfer
	transfer = models.Transfer{
		InvoiceID:   invoiceID,
		Beneficiary: payload.Beneficiary,
		AddressID:   address.ID,
	}

	// Save new transfer
	if err := h.Database.Create(&transfer).Error; err != nil {
		tx.Rollback()
		return ctx.JsonError(err)
	}

	// Increase address gap
	atomicUpdate := h.Database.Model(&account).
		Where("id = ? AND gap < ?", account.ID, core.Bip44AddressLimit).
		Update("gap", gorm.Expr("gap + 1"))

	rowsAffected, err := atomicUpdate.RowsAffected, atomicUpdate.Error
	if err != nil {
		tx.Rollback()
		return ctx.JsonError(err)
	}

	if rowsAffected == 0 {
		tx.Rollback()
		return ctx.JsonError(errors.New(core.ErrNoAvailableAccountsFound))
	}

	// Commit transaction
	tx.Commit()
	address.Transfers = append(address.Transfers, transfer)
	return ctx.JSON(http.StatusOK, address)
}
