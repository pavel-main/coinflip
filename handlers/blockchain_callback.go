package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"strconv"

	httpclient "github.com/ddliu/go-httpclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo"
	"github.com/pavel-main/coinflip/core"
	"github.com/pavel-main/coinflip/models"
	"github.com/pavel-main/coinflip/responses"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
)

func (h *Coinflip) BlockchainCallback(c echo.Context) error {
	ctx := c.(*core.CoinflipContext)

	// Parse number of confirmations
	confirmations, err := strconv.Atoi(c.QueryParam("confirmations"))
	if err != nil {
		return ctx.JsonError(err)
	}

	// Validate required confirmations
	if confirmations < core.Confirmations {
		message := fmt.Errorf(core.ErrNotEnoughConfirmations, core.Confirmations, confirmations)
		return ctx.JsonError(message)
	}

	// Parse transfer value
	value, err := decimal.NewFromString(c.QueryParam("value"))
	if err != nil {
		return ctx.JsonError(err)
	}

	// Find transfer by invoice
	invoiceID := c.Param("invoice_id")
	transfer := models.Transfer{}
	if h.Database.Preload("Address.Account").Where("invoice_id = ?", invoiceID).First(&transfer).RecordNotFound() {
		return ctx.JsonError(errors.New(core.ErrTransferNotFound))
	}

	// Retrieve currency conversion rate
	cryptoCompareFailure := false
	requestUrl := core.CryptoCompareBaseUrl + core.CryptoCompareBtcEthRate
	res, err := httpclient.Get(requestUrl)
	if err != nil {
		cryptoCompareFailure = true
		log.Errorf("Error during CryptoCompare API request: %s", err.Error())
	}

	// Read response body
	bodyBytes, err := res.ReadAll()
	if err != nil {
		cryptoCompareFailure = true
		log.Errorf("Error during CryptoCompare response read: %s", err.Error())
	}

	// Unmarshal response
	response := responses.CryptoCompareBtcEthRate{}
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		cryptoCompareFailure = true
		log.Errorf("Error during CryptoCompare response decode: %s", err.Error())
	}

	// Figure out which rate to use
	cryptoCompareRate := decimal.NewFromFloatWithExponent(response.ETH, -2)
	exchangeRate := h.Config.BtcEthFallbackRate
	if !cryptoCompareFailure && cryptoCompareRate.GreaterThan(core.Zero) {
		log.WithFields(log.Fields{"rate": cryptoCompareRate}).Info("Successfully fetched exchange rate")
		exchangeRate = cryptoCompareRate
	}

	// Calculate value in wei
	bitcoinValue := value.Div(core.OneBitcoinInSatoshi)
	etherValue := bitcoinValue.Mul(exchangeRate)
	weiValue := etherValue.Mul(core.OneEtherInWei)
	transferValue, success := new(big.Int).SetString(weiValue.String(), 10)

	// Debug calculations
	log.WithFields(log.Fields{
		"satoshi":  value.String(),
		"bitcoin":  bitcoinValue.String(),
		"ether":    etherValue.String(),
		"wei":      weiValue.String(),
		"transfer": transferValue,
	}).Info("Calculating transfer value")

	if !success {
		return ctx.JsonError(fmt.Errorf(core.ErrBtcEthConversionFailure, value.String()))
	}

	// Call smart contract
	beneficiary := common.HexToAddress(transfer.Beneficiary)
	transaction, err := h.SaleContract.BuyTokensBTC(h.TxOpts, beneficiary, transferValue)
	if err != nil {
		return ctx.JsonError(err)
	}

	// Begin transaction
	tx := h.Database.Begin()

	// Update transfer
	transfer.Rate = exchangeRate.String()
	transfer.TxIn = c.QueryParam("transaction_hash")
	transfer.TxOut = transaction.Hash().String()
	transfer.ValueIn = value.String()
	transfer.ValueOut = transferValue.String()
	if err := h.Database.Save(&transfer).Error; err != nil {
		tx.Rollback()
		return ctx.JsonError(err)
	}

	// Delete transfer
	if err := h.Database.Delete(transfer).Error; err != nil {
		tx.Rollback()
		return ctx.JsonError(err)
	}

	// Delete address
	if err := h.Database.Delete(transfer.Address).Error; err != nil {
		tx.Rollback()
		return ctx.JsonError(err)
	}

	// Decrease gap
	transfer.Address.Account.Gap = transfer.Address.Account.Gap - 1
	if err := h.Database.Save(transfer.Address.Account).Error; err != nil {
		tx.Rollback()
		return ctx.JsonError(err)
	}

	// Commit
	tx.Commit()

	return ctx.String(http.StatusOK, core.BlockchainInfoCallbackOk)
}
