package core

// Core errors
const (
	ErrUnknown              = "Unknown error"
	ErrConfigOptionRequired = "Please set mandatory %s configuration variable"
	ErrDbConnFailure        = "Failed to connect to the Postgres: %s"
	ErrGethConnFailrue      = "Failed to connect to the Ethereum node: %s"
	ErrContractInit         = "Failed to instantiate contract: %s"
	ErrPrivateKeyInit       = "Failed to load private key: %v"
	ErrNewRelicAgent        = "Error creating NewRelic agent: %s"
)

// Handlers errors
const (
	ErrInvalidEthereumAddress    = "Invalid Ethereum address: %s"
	ErrNotEnoughConfirmations    = "Required number of confirmations: %d, got %d"
	ErrNoAvailableAccountsFound  = "No available accounts found"
	ErrTransferNotFound          = "Transfer not found"
	ErrTransferNotFoundCompleted = "Transfer not found or completed"
	ErrBtcEthConversionFailure   = "Failed converting %s satoshi to wei"
	ErrPriceUpdateFailure        = "Failed to set price to value %s"
	ErrPriceUpdateMinValue       = "Failed to set price to value %s, should be bigger than %d"
)
