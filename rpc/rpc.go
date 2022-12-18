package rpc

import (
	"context"
	"fmt"
	"time"

	// rpc "github.com/DonggyuLim/grc20/protos/RPC"
	"github.com/donggyuLim/suino-server/types"
	rpc "github.com/ybbus/jsonrpc/v3"
)

const URL = "https://fullnode.devnet.sui.io:443"

const CORE = "0x5346c7e1958236278abfc247e4aa9bffeac880d6"

// type Client struct {
// 	SocketClient *websocket.Conn
// }

var client rpc.RPCClient

func Init() {
	rpcClient := rpc.NewClient(URL)
	client = rpcClient
}

func GetFlipEvent(Transaction string) types.Flip {
	exp := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), exp)
	defer cancel()
	tx := Tx{
		Transaction,
	}
	cursor := Cursor{
		TxSeq:    500,
		EventSeq: 1,
	}

	var response flipEventResponse

	client.CallFor(ctx, &response, "sui_getEvents", tx, cursor, 8)
	fmt.Println(response)
	poolBalance := GetPoolBalance()

	data := response.Data[0]
	flip := types.Flip{
		PoolBalance:   poolBalance,
		Timestamp:     data.Timestamp,
		Transaction:   Transaction,
		Address:       data.Event.MoveEvent.Sender,
		BetAmount:     data.Event.MoveEvent.Fields.BetAmount,
		BetValue:      data.Event.MoveEvent.Fields.BetValue,
		IsJackpot:     data.Event.MoveEvent.Fields.IsJackpot,
		JackpotAmount: data.Event.MoveEvent.Fields.JackpotAmount,
		JackpotValue:  data.Event.MoveEvent.Fields.JackpotValue,
	}

	return flip
}

func GetPoolBalance() int {
	exp := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), exp)
	defer cancel()
	var response getPoolBalanceResponse
	client.CallFor(ctx, &response, "sui_getObject", CORE)
	balance := response.Details.Data.Fields.Pool
	return balance
}
