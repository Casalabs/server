package types

type Flip struct {
	PoolBalance   int    `json:"pool_balance"`
	Timestamp     int64  `json:"timestamp"`
	Transaction   string `json:"transaction"`
	Address       string `json:"address"`
	BetAmount     int    `json:"bet_amount"`
	BetValue      []int  `json:"bet_value"`
	IsJackpot     bool   `json:"is_jackpot"`
	JackpotAmount int    `json:"jackpot_amount"`
	JackpotValue  []int  `json:"jackpot_value"`
}
