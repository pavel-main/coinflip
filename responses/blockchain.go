package responses

type BlockchainInfoGap struct {
	Gap int `json:"gap"`
}

type BlockchainInfoCallback struct {
	Callback     string `json:"callback"`
	CalledAt     string `json:"called_at"`
	RawResponse  string `json:"raw_response"`
	ResponseCode int    `json:"response_code"`
}

type BlockchainInfoReceive struct {
	Address  string  `json:"address"`
	Index    int     `json:"index"`
	Callback string  `json:"callback"`
	Message  *string `json:"message"`
}
