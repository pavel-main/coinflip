package handlers

import (
	"encoding/json"
	"net/http"
	"net/url"

	httpclient "github.com/ddliu/go-httpclient"
	"github.com/labstack/echo"
	"github.com/pavel-main/coinflip/core"
	"github.com/pavel-main/coinflip/responses"
)

func (h *Coinflip) BlockchainCallbackLogs(c echo.Context) error {
	ctx := c.(*core.CoinflipContext)

	// Prepare callback URL
	callbackUrl := url.URL{
		Scheme: h.Config.Protocol,
		Host:   h.Config.Domain,
		Path:   "/blockchain/callback/" + ctx.QueryParam("invoice_id"),
	}

	// Perform API request
	requestUrl := core.BlockchainInfoBaseUrl + core.BlockchainInfoCallbackLog
	res, err := httpclient.Get(requestUrl, map[string]string{
		"key":      h.Config.BlockchainInfoApiKey,
		"callback": callbackUrl.String(),
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
	response := []responses.BlockchainInfoCallback{}
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return ctx.JsonError(err)
	}

	return ctx.JSON(http.StatusOK, response)
}
