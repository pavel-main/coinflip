package core

import (
	"math/big"

	"github.com/shopspring/decimal"
)

// Common options
const (
	HealthcheckResponse = "OK"
	MinPaymentValue     = 800
)

// Features
const (
	FeatureStats      = "stats"
	FeatureWhitelist  = "whitelist"
	FeatureBlockchain = "blockchain"
	FeatureNewRelic   = "newrelic"
	FeaturePrice      = "price"
)

// Blockchain.info options
const (
	Confirmations            = 3
	BlockchainInfoCallbackOk = "*ok*"
	Bip44AddressLimit        = 20
)

var (
	Zero                = decimal.New(0, 0)
	OneBitcoinInSatoshi = decimal.New(int64(1e+8), 0)
	OneEtherInWei       = decimal.New(int64(1e+18), 0)
	MinPayment          = big.NewInt(MinPaymentValue)
)

// Configuration options
const (
	ConfigOptionAppName              = "app_name"
	ConfigOptionNewRelicLicenseKey   = "newrelic_license_key"
	ConfigOptionDebug                = "debug"
	ConfigOptionLogLevel             = "log_level"
	ConfigOptionDatabase             = "database"
	ConfigOptionPort                 = "port"
	ConfigOptionProtocol             = "protocol"
	ConfigOptionFeatures             = "features"
	ConfigOptionDomain               = "domain"
	ConfigOptionGethIpcPath          = "geth_ipc_path"
	ConfigOptionEthPrivateKey        = "eth_private_key"
	ConfigOptionEthSaleContract      = "eth_sale_contract"
	ConfigOptionEthTokenContract     = "eth_token_contract"
	ConfigOptionBlockchainInfoApiKey = "blockchain_info_api_key"
	ConfigOptionBtcEthFallbackRate   = "btceth_fallback_rate"
	ConfigOptionHttpConnectTimeout   = "http_connect_timeout"
	ConfigOptionHttpTimeout          = "http_timeout"
)

// Configuration defaults
const (
	ConfigDefaultAppName            = "coinflip"
	ConfigDefaultDebug              = false
	ConfigDefaultPort               = 3000
	ConfigDefaultLogLevel           = "INFO"
	ConfigDefaultHttpTimeout        = 20000
	ConfigDefaultHttpConnectTimeout = 10000
	ConfigDefaultProtocol           = "https"
	ConfigDefaultFeatures           = "stats whitelist blockchain"
)

// Blockchain.info API defaults
const (
	BlockchainInfoBaseUrl     = "https://api.blockchain.info"
	BlockchainInfoReceive     = "/v2/receive"
	BlockchainInfoCallbackLog = "/v2/receive/callback_log"
	BlockchainInfoAddressGap  = "/v2/receive/checkgap"
)

// CryptoCompare API defaults
const (
	CryptoCompareBaseUrl    = "https://min-api.cryptocompare.com"
	CryptoCompareBtcEthRate = "/data/price?fsym=BTC&tsyms=ETH"
)

func GetCallbackUrl(protocol string, domain string, invoiceID string) string {
	return protocol + "://" + domain + "/blockchain/callback/" + invoiceID
}
