package sui

import (
	"encoding/json"
	"fmt"
	"log"

	"net/url"

	"github.com/coming-chat/go-sui/client"
	"github.com/donggyuLim/suino-server/db"
	"github.com/donggyuLim/suino-server/utils"
	"github.com/gorilla/websocket"
)

type Sui struct {
	SocketClient websocket.Conn
	RPC          client.Client
}

func NewClient(socketHost, rpcEndpoint string) Sui {
	// defer wg.Done()
	u := url.URL{Scheme: "wss", Host: socketHost}
	dial := websocket.DefaultDialer

	conn, _, err := dial.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	rpc, err := client.Dial(rpcEndpoint)
	if err != nil {
		log.Fatal("rpc:", err.Error())
	}

	sui := Sui{
		SocketClient: *conn,
		RPC:          *rpc,
	}
	return sui
}

func subscribe(c *websocket.Conn) {
	params := Params{
		All: []All{
			{
				EventType: "MoveEvent",
			},
			{
				Package: utils.LoadENV("CONTRACT"),
			},
		},
	}

	subscribe := SubscribeEvent{
		Jsonrpc: "2.0",
		ID:      1,
		Method:  "sui_subscribeEvent",
		Params:  []Params{params},
	}

	err := c.WriteJSON(subscribe)
	if err != nil {
		log.Println("write:", err)
	}

	_, _, err = c.ReadMessage()
	if err != nil {
		log.Println("write:", err.Error())
	}
	fmt.Println("Subscribe!!!")
}

func (s *Sui) HandleMsg(ch chan interface{}) error {
	defer s.SocketClient.Close()
	subscribe(&s.SocketClient)
	for {
		_, message, err := s.SocketClient.ReadMessage()
		if err != nil {
			fmt.Println("Read:", err)

			return err
		}

		event := EventResponse{}
		err = json.Unmarshal(message, &event)
		if err != nil {
			fmt.Println("Marshal :", err.Error())
			return err
		}
		data := Data{
			TimeStamp: event.Params.Result.Timestamp,
			TxDigest:  event.Params.Result.TxDigest,
			Module:    event.Params.Result.Event.MoveEvent.TransactionModule,
			MoveEvent: MoveEvent{
				Gamer:         event.Params.Result.Event.MoveEvent.Fields.Gamer,
				BetAmount:     event.Params.Result.Event.MoveEvent.Fields.BetAmount,
				BetValue:      event.Params.Result.Event.MoveEvent.Fields.BetValue,
				IsJackpot:     event.Params.Result.Event.MoveEvent.Fields.IsJackpot,
				JackpotAmount: event.Params.Result.Event.MoveEvent.Fields.JackpotAmount,
				JackpotValue:  event.Params.Result.Event.MoveEvent.Fields.JackpotValue,
				PoolBalance:   event.Params.Result.Event.MoveEvent.Fields.PoolBalance,
			},
		}
		fmt.Println(data)

		HandleDB(data)
		// event := Event{
		// 	Type: response.Params.Result.Event.MoveEvent.TransactionModule,
		// 	Event:
		// }

		ch <- data
	}
}

func HandleDB(data Data) {
	switch data.Module {
	case "flip", "race":
		db.Insert("game", data)
	}
}
