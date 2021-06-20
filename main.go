package main

import (
	echorelic "github.com/jessie-codes/echo-relic"
	"github.com/labstack/echo"
	newrelic "github.com/newrelic/go-agent"
	"github.com/pavel-main/coinflip/core"
	"github.com/pavel-main/coinflip/handlers"
	"github.com/pavel-main/echo-logrusmiddleware"
	log "github.com/sirupsen/logrus"
)

// Provided by `govvv` tool
var (
	GitCommit  string
	GitBranch  string
	GitState   string
	GitSummary string
	BuildDate  string
)

func main() {
	// Init config & Coinflip handler
	build := core.NewBuild(GitCommit, GitBranch, GitState, GitSummary, BuildDate)
	cfg := core.NewConfig("cf", build)
	coinflip := handlers.NewCoinflip(cfg)

	// Init echo
	e := echo.New()
	e.HideBanner = true

	// Set up logrus middleware
	e.Logger = logrusmiddleware.Logger{log.StandardLogger()}
	e.Use(logrusmiddleware.Hook())

	// Configure NewRelic if necessary
	config := newrelic.NewConfig(cfg.AppName, cfg.NewRelicLicenseKey)
	app, err := newrelic.NewApplication(config)
	if err != nil {
		log.Fatalf(core.ErrNewRelicAgent, err.Error())
	}
	e.Use(echorelic.Middleware(app))

	// Wire up custom context
	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := &core.CoinflipContext{c}
			return h(ctx)
		}
	})

	// Built-in routes
	e.GET("/", coinflip.Healthcheck)
	e.GET("/version", coinflip.Version)

	// Stats feature
	if coinflip.HasFeature(core.FeatureStats) {
		g1 := e.Group("stats")
		g1.GET("", coinflip.StatsGet)
		g1.GET("/balance/:address", coinflip.BalanceGet)
	}

	// Whitelist feature
	if coinflip.HasFeature(core.FeatureWhitelist) {
		g2 := e.Group("whitelist")
		g2.GET("/:address", coinflip.WhitelistGet)
		g2.POST("", coinflip.WhitelistPost)
		g2.DELETE("", coinflip.WhitelistDelete)
	}

	// Blockchain.info feature
	if coinflip.HasFeature(core.FeatureBlockchain) {
		g3 := e.Group("blockchain")
		g3.POST("/receive", coinflip.BlockchainReceive)
		g3.GET("/callback/:invoice_id", coinflip.BlockchainCallback)
		g3.GET("/invoice/:id", coinflip.BlockchainInvoiceGet)
		g3.GET("/callback/logs/:invoice_id", coinflip.BlockchainCallbackLogs)
		g3.GET("/gap/:xpub", coinflip.BlockchainGapCheck)
	}

	// Price update feature
	if coinflip.HasFeature(core.FeaturePrice) {
		g4 := e.Group("price")
		g4.POST("/update", coinflip.UpdatePrice)
	}

	// Print configuration variables
	log.WithFields(log.Fields{
		"from":                              coinflip.TxOpts.From.String(),
		core.ConfigOptionAppName:            cfg.AppName,
		core.ConfigOptionDomain:             cfg.Domain,
		core.ConfigOptionPort:               cfg.Port,
		core.ConfigOptionProtocol:           cfg.Protocol,
		core.ConfigOptionDebug:              cfg.Debug,
		core.ConfigOptionFeatures:           cfg.Features,
		core.ConfigOptionGethIpcPath:        cfg.GethIpcPath,
		core.ConfigOptionEthSaleContract:    cfg.EthSaleContract,
		core.ConfigOptionEthTokenContract:   cfg.EthTokenContract,
		core.ConfigOptionBtcEthFallbackRate: cfg.BtcEthFallbackRate,
		core.ConfigOptionHttpConnectTimeout: cfg.HttpConnectTimeout,
		core.ConfigOptionHttpTimeout:        cfg.HttpTimeout,
	}).Info("Coinflip configuration")

	// Print build metadata
	log.WithFields(log.Fields{
		"git_commit":  cfg.Build.GitCommit,
		"git_branch":  cfg.Build.GitBranch,
		"git_state":   cfg.Build.GitState,
		"git_summary": cfg.Build.GitSummary,
		"build_date":  cfg.Build.BuildDate,
	}).Info("Build metadata")

	// Start server
	e.Logger.Fatal(e.Start(":" + cfg.Port))
}
