package responses

type WhitelistPost struct {
	TxHash string `json:"tx_hash"`
}

type WhitelistGet struct {
	Whitelisted bool `json:"whitelisted"`
}
