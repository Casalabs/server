package sui

// type SubscibeResponse struct {
// 	Jsonrpc string `json:"jsonrpc"`
// 	Result  int    `json:"result"`
// 	ID      int    `json:"id"`
// }
// type SubscribeEvent struct {
// 	Jsonrpc string   `json:"jsonrpc"`
// 	ID      int      `json:"id"`
// 	Method  string   `json:"method"`
// 	Params  []Params `json:"params"`
// }
// type Params struct {
// 	Param []Param `json:"All"`
// }

// type Param struct {
// 	EventType     string `json:"EventType,omitempty"`
// 	Package       string `json:"Package,omitempty"`
// 	Module        string `json:"Module,omitempty"`
// 	SenderAddress string `json:"SenderAddress,omitempty"`
// }

// func (s Sui) SubScribeEventModule(EventType, Package, Module string) {
// 	params := Params{
// 		Param: []Param{
// 			{
// 				EventType: EventType,
// 			},
// 			{
// 				Package: Package,
// 			},
// 			{
// 				Module: Module,
// 			},
// 		},
// 	}
// 	subscribe := SubscribeEvent{
// 		Jsonrpc: "2.0",
// 		ID:      1,
// 		Method:  "sui_subscribeEvent",
// 		Params:  []Params{params},
// 	}
// 	s.SocketClient.WriteJSON(subscribe)
// }

// func (c Sui) SubScribeEventOnlyPublish() {
// 	params := Params{
// 		Param: []Param{
// 			{
// 				EventType: "Publish",
// 			},
// 		},
// 	}
// 	subscribe := SubscribeEvent{
// 		Jsonrpc: "2.0",
// 		ID:      1,
// 		Method:  "sui_subscribeEvent",
// 		Params:  []Params{params},
// 	}
// 	c.SocketClient.WriteJSON(subscribe)
// }

// func (c Sui) SubscribeAllEvent() {
// 	params := Params{
// 		Param: []Param{
// 			{},
// 		},
// 	}
// 	subscribe := SubscribeEvent{
// 		Jsonrpc: "2.0",
// 		ID:      1,
// 		Method:  "sui_subscribeEvent",
// 		Params:  []Params{params},
// 	}
// 	c.SocketClient.WriteJSON(subscribe)

// }

// func (c Sui) SubscribeActivityFromAddress(address string) {
// 	params := Params{
// 		Param: []Param{
// 			{
// 				SenderAddress: address,
// 			},
// 		},
// 	}
// 	subscribe := SubscribeEvent{
// 		Jsonrpc: "2.0",
// 		ID:      1,
// 		Method:  "sui_subscribeEvent",
// 		Params:  []Params{params},
// 	}
// 	c.SocketClient.WriteJSON(subscribe)

// }

// func (c Sui) SubscribeCoinBalanceChange(address string) {
// 	params := Params{
// 		Param: []Param{
// 			{
// 				SenderAddress: address,
// 			},
// 			{
// 				EventType: "CoinBalanceChange",
// 			},
// 		},
// 	}
// 	subscribe := SubscribeEvent{
// 		Jsonrpc: "2.0",
// 		ID:      1,
// 		Method:  "sui_subscribeEvent",
// 		Params:  []Params{params},
// 	}
// 	c.SocketClient.WriteJSON(subscribe)

// }

// func (c Sui) SubscribeAllTransferObject() {
// 	params := Params{
// 		Param: []Param{

// 			{
// 				EventType: "TransferObject",
// 			},
// 		},
// 	}
// 	subscribe := SubscribeEvent{
// 		Jsonrpc: "2.0",
// 		ID:      1,
// 		Method:  "sui_subscribeEvent",
// 		Params:  []Params{params},
// 	}
// 	c.SocketClient.WriteJSON(subscribe)
// }

// func (c Sui) SubscribeOnlyAddressTransferObject(address string) {
// 	params := Params{
// 		Param: []Param{

// 			{
// 				EventType: "TransferObject",
// 			},
// 			{
// 				SenderAddress: address,
// 			},
// 		},
// 	}
// 	subscribe := SubscribeEvent{
// 		Jsonrpc: "2.0",
// 		ID:      1,
// 		Method:  "sui_subscribeEvent",
// 		Params:  []Params{params},
// 	}
// 	c.SocketClient.WriteJSON(subscribe)
// }

// func (c Sui) SubscribePackageModifiy(Package string) {
// 	params := Params{
// 		Param: []Param{
// 			{
// 				EventType: "MoveEvent",
// 			},
// 			{
// 				Package: Package,
// 			},
// 		},
// 	}
// 	subscribe := SubscribeEvent{
// 		Jsonrpc: "2.0",
// 		ID:      1,
// 		Method:  "sui_subscribeEvent",
// 		Params:  []Params{params},
// 	}
// 	c.SocketClient.WriteJSON(subscribe)
// }
