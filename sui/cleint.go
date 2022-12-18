package sui

import (
	"encoding/json"
	"fmt"
	"log"

	"net/url"
	"sync"

	"github.com/coming-chat/go-sui/account"
	"github.com/coming-chat/go-sui/client"
	"github.com/gorilla/websocket"
)

type Sui struct {
	SocketClient *websocket.Conn
	RPC          *client.Client
	Account      *account.Account
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
		fmt.Println("message:", string(message))

		response := EventResponse{}
		err = json.Unmarshal(message, &response)
		if err != nil {
			fmt.Println("Marshal :", err.Error())
			continue
		}

		fmt.Println("RESPONSE======", response)

		fmt.Println("")
		// event := Event{
		// 	Type: response.Params.Result.Event.MoveEvent.TransactionModule,
		// 	Event:
		// }

		ch <- message
	}
}

func subscribe(c *websocket.Conn) {
	params := Params{
		All: []All{
			{
				EventType: "MoveEvent",
			},
			{
				Package: "0x8391436389a857bfd1493294d1c106cbde3a6800",
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
