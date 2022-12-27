package sui

import (
	"encoding/json"
	"fmt"
	"log"

	"net/url"
	"sync"

	"github.com/coming-chat/go-sui/client"
	"github.com/donggyuLim/suino-server/db"
	"github.com/donggyuLim/suino-server/utils"
	"github.com/gorilla/websocket"
)

type Sui struct {
	SocketClient *websocket.Conn
	RPC          *client.Client
}

func NewClient(socketHost, rpcEndpoint string, wg *sync.WaitGroup) Sui {
	defer wg.Done()
	u := url.URL{Scheme: "wss", Host: socketHost}

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	rpc, err := client.Dial(rpcEndpoint)
	if err != nil {
		log.Fatal("rpc:", err.Error())
	}

	sui := Sui{
		SocketClient: conn,
		RPC:          rpc,
	}
	return sui
}

type Event struct {
	Type  string      `json:"Type"`
	Event interface{} `json:"event"`
}

func (s *Sui) HandleMsg(ch chan interface{}) {
	defer s.SocketClient.Close()
	subscribe(s.SocketClient)
	for {
		_, message, err := s.SocketClient.ReadMessage()
		if err != nil {
			fmt.Println("read:", err)
		}

		event := EventResponse{}
		err = json.Unmarshal(message, &event)
		if err != nil {
			fmt.Println("Marshal :", err.Error())
			continue
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

type Data struct {
	TimeStamp int64
	TxDigest  string
	Module    string
	MoveEvent MoveEvent
}

type MoveEvent struct {
	Gamer         string
	BetAmount     string
	BetValue      []string
	IsJackpot     bool
	JackpotAmount string
	JackpotValue  []string
	PoolBalance   string
}

func HandleDB(data Data) {
	switch data.Module {
	case "flip", "race":
		db.Insert("game", data)
	}
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

}
