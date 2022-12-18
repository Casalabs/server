package rpc

type flipEventResponse struct {
	Data []struct {
		Timestamp int64  `json:"timestamp"`
		TxDigest  string `json:"txDigest"`
		ID        struct {
			TxSeq    int `json:"txSeq"`
			EventSeq int `json:"eventSeq"`
		} `json:"id"`
		Event struct {
			MoveEvent struct {
				PackageID         string `json:"packageId"`
				TransactionModule string `json:"transactionModule"`
				Sender            string `json:"sender"`
				Type              string `json:"type"`
				Fields            struct {
					BetAmount      int    `json:"bet_amount"`
					BetValue       []int  `json:"bet_value"`
					IsJackpot      bool   `json:"is_jackpot"`
					JackpotAddress string `json:"jackpot_address"`
					JackpotAmount  int    `json:"jackpot_amount"`
					JackpotValue   []int  `json:"jackpot_value"`
				} `json:"fields"`
				Bcs string `json:"bcs"`
			} `json:"moveEvent"`
		} `json:"event"`
	} `json:"data"`
	NextCursor interface{} `json:"nextCursor"`
	ID         int         `json:"id"`
}


type Test struct {
	Timestamp int64  `json:"timestamp"`
	TxDigest  string `json:"txDigest"`
	ID        struct {
		TxSeq    int `json:"txSeq"`
		EventSeq int `json:"eventSeq"`
	} `json:"id"`
	Event struct {
		MoveEvent struct {
			PackageID         string `json:"packageId"`
			TransactionModule string `json:"transactionModule"`
			Sender            string `json:"sender"`
			Type              string `json:"type"`
			Fields            struct {
				BetAmount      int    `json:"bet_amount"`
				BetValue       []int  `json:"bet_value"`
				IsJackpot      bool   `json:"is_jackpot"`
				JackpotAddress string `json:"jackpot_address"`
				JackpotAmount  int    `json:"jackpot_amount"`
				JackpotValue   []int  `json:"jackpot_value"`
			} `json:"fields"`
			Bcs string `json:"bcs"`
		} `json:"moveEvent"`
	} `json:"event"`
}

type getPoolBalanceResponse struct {
	Status  string `json:"status"`
	Details struct {
		Data struct {
			DataType          string `json:"dataType"`
			Type              string `json:"type"`
			HasPublicTransfer bool   `json:"has_public_transfer"`
			Fields            struct {
				Description      string `json:"description"`
				GamingFeePercent int    `json:"gaming_fee_percent"`
				ID               struct {
					ID string `json:"id"`
				} `json:"id"`
				Lock           bool   `json:"lock"`
				LotteryPercent int    `json:"lottery_percent"`
				Name           string `json:"name"`
				Owners         int    `json:"owners"`
				Pool           int    `json:"pool"`
				RewardPool     int    `json:"reward_pool"`
				Sign           struct {
					Type   string `json:"type"`
					Fields struct {
						Contents string `json:"contents"`
					} `json:"fields"`
				} `json:"sign"`
			} `json:"fields"`
		} `json:"data"`
		Owner struct {
			Shared struct {
				InitialSharedVersion int `json:"initial_shared_version"`
			} `json:"Shared"`
		} `json:"owner"`
		PreviousTransaction string `json:"previousTransaction"`
		StorageRebate       int    `json:"storageRebate"`
		Reference           struct {
			ObjectID string `json:"objectId"`
			Version  int    `json:"version"`
			Digest   string `json:"digest"`
		} `json:"reference"`
	} `json:"details"`
}
