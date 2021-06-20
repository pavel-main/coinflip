package responses

type Stats struct {
	Active         bool   `json:"active"`
	AvailableUnits uint64 `json:"available_units"`
	AvailableBonus uint64 `json:"available_bonus"`
	BonusUsed      uint64 `json:"bonus_used"`
	Duration       uint64 `json:"duration"`
	MinPayment     uint64 `json:"min_payment"`
	Price          uint64 `json:"price"`
	StartTime      uint64 `json:"start_time"`
	UnitsSold      uint64 `json:"units_sold"`
	WeiReceived    uint64 `json:"wei_received"`
}

type Balance struct {
	Balance uint64 `json:"balance"`
}
