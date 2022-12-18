package rpc

// type SubscribeEvent struct {
// 	Jsonrpc string   `json:"jsonrpc"`
// 	ID      int      `json:"id"`
// 	Method  string   `json:"method"`
// 	Params  []Params `json:"params"`
// }

// type All struct {
// 	EventType string `json:"EventType,omitempty"`
// 	Package   string `json:"Package,omitempty"`
// 	Module    string `json:"Module,omitempty"`
// }
// type Params struct {
// 	All []All `json:"All"`
// }

// func SuiClient(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	u := url.URL{Scheme: "wss", Host: "fullnode.devnet.sui.io:443"}
// 	fmt.Println(u.String())
// 	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
// 	if err != nil {
// 		log.Fatal("dial:", err.Error())
// 	}
// 	go HandleMsg(c)
// }

// func HandleMsg(c *websocket.Conn) {
// 	defer c.Close()
// 	subscribe(c)
// 	for {
// 		_, message, err := c.ReadMessage()
// 		if err != nil {
// 			fmt.Println("read:", err)
// 		}
// 		fmt.Println("message:", string(message))
// 	}

// }

// func subscribe(c *websocket.Conn) {
// 	params := Params{
// 		All: []All{
// 			{
// 				EventType: "MoveEvent",
// 			},
// 			{
// 				Package: "0x1e354db4a3cf0bfa288072e44757538d96916c85",
// 			},
// 			{
// 				Module: "flip",
// 			},
// 		},
// 	}

// 	subscribe := SubscribeEvent{
// 		Jsonrpc: "2.0",
// 		ID:      1,
// 		Method:  "sui_subscribeEvent",
// 		Params:  []Params{params},
// 	}

// 	err := c.WriteJSON(subscribe)
// 	if err != nil {
// 		log.Println("write:", err)
// 	}

// }
