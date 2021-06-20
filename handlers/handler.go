package handlers

import (
	httpclient "github.com/ddliu/go-httpclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pavel-main/coinflip/contracts"
	"github.com/pavel-main/coinflip/core"
	"github.com/pavel-main/coinflip/models"
	"github.com/pavel-main/gormrus"
	log "github.com/sirupsen/logrus"
)

type Coinflip struct {
	Config        *core.Config
	Database      *gorm.DB
	SaleContract  *contracts.TokenSale
	TokenContract *contracts.ShopToken
	TxOpts        *bind.TransactOpts
}

func NewCoinflip(cfg *core.Config) *Coinflip {
	// Configure http client
	httpclient.Defaults(httpclient.Map{
		httpclient.OPT_DEBUG:             cfg.Debug,
		httpclient.OPT_CONNECTTIMEOUT_MS: cfg.HttpConnectTimeout,
		httpclient.OPT_TIMEOUT_MS:        cfg.HttpTimeout,
		"Content-Type":                   "application/json",
		"Accept":                         "application/json",
	})

	// Configure database connection
	db, err := gorm.Open("postgres", cfg.Database)
	if err != nil {
		log.Fatalf(core.ErrDbConnFailure, err)
	}

	// Configure database handler
	db.LogMode(cfg.Debug)
	db.SetLogger(gormrus.New())

	// Auto-migrate
	db.AutoMigrate(&models.Account{}, &models.Address{}, &models.Transfer{})

	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial(cfg.GethIpcPath)
	if err != nil {
		log.Fatalf(core.ErrGethConnFailrue, err)
	}

	// Instantiate sale contract
	saleContract, err := contracts.NewTokenSale(common.HexToAddress(cfg.EthSaleContract), conn)
	if err != nil {
		log.Fatalf(core.ErrContractInit, err)
	}

	// Instantiate token contract
	tokenContract, err := contracts.NewShopToken(common.HexToAddress(cfg.EthTokenContract), conn)
	if err != nil {
		log.Fatalf(core.ErrContractInit, err)
	}

	// Load private key
	ecdsaKey, err := crypto.HexToECDSA(cfg.EthPrivateKey)
	if err != nil {
		log.Fatalf(core.ErrPrivateKeyInit, err)
	}

	// Instantiate handler
	h := new(Coinflip)
	h.Config = cfg
	h.Database = db
	h.SaleContract = saleContract
	h.TokenContract = tokenContract
	h.TxOpts = bind.NewKeyedTransactor(ecdsaKey)

	return h
}

func (h *Coinflip) HasFeature(feature string) bool {
	for _, item := range h.Config.Features {
		if feature == item {
			return true
		}
	}

	return false
}
