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
	"github.com/mitchellh/mapstructure"
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
		fmt.Println("Handle MSG")
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
		// fmt.Println(event.Params.Result.Event.MoveEvent)
		handleData(event, ch)

	}
}

func handleData(data EventResponse, ch chan interface{}) {
	result := data.Params.Result
	switch result.Event.MoveEvent.TransactionModule {
	case "flip":

		f := &FlipEvent{}

		fields := result.Event.MoveEvent.Fields
		err := mapstructure.Decode(fields, f)
		if err != nil {
			log.Fatal(err)
		}
		data := Data{
			Module:        result.Event.MoveEvent.TransactionModule,
			TimeStamp:     result.Timestamp,
			TxDigest:      result.TxDigest,
			Gamer:         f.Gamer,
			BetAmount:     f.BetAmount,
			BetValue:      f.BetValue,
			IsJackpot:     f.IsJackpot,
			JackpotAmount: f.JackpotAmount,
			JackpotValue:  f.JackpotValue,
			PoolBalance:   f.PoolBalance,
		}
		log.Println(data)
		db.Insert("game", data)
		ch <- data
	case "nft":

	}

}
