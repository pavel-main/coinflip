package handlers

import (
	"encoding/json"
	"net/http"

	httpclient "github.com/ddliu/go-httpclient"
	"github.com/labstack/echo"
	"github.com/pavel-main/coinflip/core"
	"github.com/pavel-main/coinflip/responses"
)

func (h *Coinflip) BlockchainGapCheck(c echo.Context) error {
	ctx := c.(*core.CoinflipContext)

	// Perform API request
	requestUrl := core.BlockchainInfoBaseUrl + core.BlockchainInfoAddressGap
	res, err := httpclient.Get(requestUrl, map[string]string{
		"key":  h.Config.BlockchainInfoApiKey,
		"xpub": c.Param("xpub"),
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
	response := responses.BlockchainInfoGap{}
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return ctx.JsonError(err)
	}

	return ctx.JSON(http.StatusOK, response)
}
