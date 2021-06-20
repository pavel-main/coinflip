package core

import (
	"strings"

	"github.com/joho/godotenv"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	AppName              string
	Build                *Build
	Debug                bool
	Database             string
	Domain               string
	LogLevel             string
	Port                 string
	Protocol             string
	Features             []string
	GethIpcPath          string
	EthPrivateKey        string
	EthSaleContract      string
	EthTokenContract     string
	BlockchainInfoApiKey string
	BtcEthFallbackRate   decimal.Decimal
	NewRelicLicenseKey   string
	HttpConnectTimeout   int
	HttpTimeout          int
}

func NewConfig(prefix string, build *Build) *Config {
	// Load .env if exists
	godotenv.Load()

	// Viper instance
	viper.AutomaticEnv()
	viper.SetEnvPrefix(prefix)
	viper.SetDefault(ConfigDefaultAppName, ConfigDefaultAppName)
	viper.SetDefault(ConfigOptionPort, ConfigDefaultPort)
	viper.SetDefault(ConfigOptionDebug, ConfigDefaultDebug)
	viper.SetDefault(ConfigOptionLogLevel, ConfigDefaultLogLevel)
	viper.SetDefault(ConfigOptionFeatures, ConfigDefaultFeatures)
	viper.SetDefault(ConfigOptionProtocol, ConfigDefaultProtocol)
	viper.SetDefault(ConfigOptionHttpConnectTimeout, ConfigDefaultHttpConnectTimeout)
	viper.SetDefault(ConfigOptionHttpTimeout, ConfigDefaultHttpTimeout)

	// Load configuration variables
	cfg := new(Config)
	cfg.Build = build
	cfg.AppName = viper.GetString(ConfigOptionAppName)
	cfg.Port = viper.GetString(ConfigOptionPort)
	cfg.Features = viper.GetStringSlice(ConfigOptionFeatures)
	cfg.Debug = viper.GetBool(ConfigOptionDebug)
	cfg.LogLevel = viper.GetString(ConfigOptionLogLevel)
	cfg.NewRelicLicenseKey = viper.GetString(ConfigOptionNewRelicLicenseKey)
	cfg.Protocol = viper.GetString(ConfigOptionProtocol)
	cfg.HttpConnectTimeout = viper.GetInt(ConfigOptionHttpConnectTimeout)
	cfg.HttpTimeout = viper.GetInt(ConfigOptionHttpTimeout)

	// Fail fast
	cfg.Database = viper.GetString(ConfigOptionDatabase)
	if cfg.Database == "" {
		option := strings.ToUpper(prefix + "_" + ConfigOptionDatabase)
		log.Fatalf(ErrConfigOptionRequired, option)
	}

	cfg.Domain = viper.GetString(ConfigOptionDomain)
	if cfg.Domain == "" {
		option := strings.ToUpper(prefix + "_" + ConfigOptionDomain)
		log.Fatalf(ErrConfigOptionRequired, option)
	}

	cfg.GethIpcPath = viper.GetString(ConfigOptionGethIpcPath)
	if cfg.GethIpcPath == "" {
		option := strings.ToUpper(prefix + "_" + ConfigOptionGethIpcPath)
		log.Fatalf(ErrConfigOptionRequired, option)
	}

	cfg.EthPrivateKey = viper.GetString(ConfigOptionEthPrivateKey)
	if cfg.EthPrivateKey == "" {
		option := strings.ToUpper(prefix + "_" + ConfigOptionEthPrivateKey)
		log.Fatalf(ErrConfigOptionRequired, option)
	}

	cfg.EthSaleContract = viper.GetString(ConfigOptionEthSaleContract)
	if cfg.EthSaleContract == "" {
		option := strings.ToUpper(prefix + "_" + ConfigOptionEthSaleContract)
		log.Fatalf(ErrConfigOptionRequired, option)
	}

	cfg.EthTokenContract = viper.GetString(ConfigOptionEthTokenContract)
	if cfg.EthTokenContract == "" {
		option := strings.ToUpper(prefix + "_" + ConfigOptionEthTokenContract)
		log.Fatalf(ErrConfigOptionRequired, option)
	}

	cfg.BlockchainInfoApiKey = viper.GetString(ConfigOptionBlockchainInfoApiKey)
	if cfg.BlockchainInfoApiKey == "" {
		option := strings.ToUpper(prefix + "_" + ConfigOptionBlockchainInfoApiKey)
		log.Fatalf(ErrConfigOptionRequired, option)
	}

	fallbackRate, err := decimal.NewFromString(viper.GetString(ConfigOptionBtcEthFallbackRate))
	if err != nil {
		log.Fatal(err)
	}

	// Set log level
	level, err := log.ParseLevel(cfg.LogLevel)
	if err != nil {
		log.Fatal(err)
	}

	log.SetLevel(level)

	// Set fallback rate
	cfg.BtcEthFallbackRate = fallbackRate
	return cfg
}
