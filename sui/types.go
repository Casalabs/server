package sui

type GetCoreResponse struct {
	Description      string `json:"description"`
	GamingFeePercent int    `json:"gaming_fee_percent"`
	ID               struct {
		ID string `json:"id"`
	} `json:"id"`
	Lock           bool   `json:"lock"`
	LotteryAmount  int    `json:"lottery_amount"`
	LotteryPercent int    `json:"lottery_percent"`
	Name           string `json:"name"`
	Owners         int    `json:"owners"`
	Pool           int    `json:"pool"`
	Random         struct {
		Type   string `json:"type"`
		Fields struct {
			Hash string `json:"hash"`
			Salt string `json:"salt"`
		} `json:"fields"`
	} `json:"random"`
	RandomFee  int `json:"random_fee"`
	RewardPool int `json:"reward_pool"`
	Sign       struct {
		Type   string `json:"type"`
		Fields struct {
			Contents string `json:"contents"`
		} `json:"fields"`
	} `json:"sign"`
}

type All struct {
	EventType string `json:"EventType,omitempty"`
	Package   string `json:"Package,omitempty"`
	Module    string `json:"Module,omitempty"`
}

type Params struct {
	All []All `json:"All"`
}

type SubscribeEvent struct {
	Jsonrpc string   `json:"jsonrpc"`
	ID      int      `json:"id"`
	Method  string   `json:"method"`
	Params  []Params `json:"params"`
}

type EventResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  struct {
		Subscription int64 `json:"subscription"`
		Result       struct {
			Timestamp int64  `json:"timestamp"`
			TxDigest  string `json:"txDigest"`
			ID        struct {
				TxSeq    int `json:"txSeq"`
				EventSeq int `json:"eventSeq"`
			} `json:"id"`
			Event struct {
				MoveEvent struct {
					PackageID         string                 `json:"packageId"`
					TransactionModule string                 `json:"transactionModule"`
					Sender            string                 `json:"sender"`
					Type              string                 `json:"type"`
					Fields            map[string]interface{} `json:"fields"`
					Bcs               string                 `json:"bcs"`
				} `json:"moveEvent"`
			} `json:"event"`
		} `json:"result"`
	} `json:"params"`
}

type Data struct {
	Module        string
	TimeStamp     int64
	TxDigest      string
	Gamer         string
	BetAmount     string
	BetValue      []string
	IsJackpot     bool
	JackpotAmount string
	JackpotValue  []string
	PoolBalance   string
}

type FlipEvent struct {
	BetAmount     string   `mapstructure:"bet_amount"`
	BetValue      []string `mapstructure:"bet_value"`
	Gamer         string   `mapstructure:"gamer"`
	IsJackpot     bool     `mapstructure:"is_jackpot"`
	JackpotAmount string   `mapstructure:"jackpot_amount"`
	JackpotValue  []string `mapstructure:"jackpot_value"`
	PoolBalance   string   `mapstructure:"pool_balance"`
}

type MoveEvent struct {
}
