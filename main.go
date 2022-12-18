package main

import (
	"sync"

	"github.com/donggyuLim/suino-server/db"
	"github.com/donggyuLim/suino-server/p2p"
	"github.com/donggyuLim/suino-server/rest"
	"github.com/donggyuLim/suino-server/sui"
)

const jsonrpcEndPoint = "https://fullnode.devnet.sui.io:443"
const websocketHost = "fullnode.devnet.sui.io:443"

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(4)
	db.DBConnect()
	defer db.Close()
	ch := make(chan interface{}, 10)
	defer close(ch)
	sui := sui.NewClient(websocketHost, jsonrpcEndPoint, wg)

	go sui.HandleMsg(ch)
	go p2p.HandleMsg(ch)
	go rest.Start(wg)

	// go sui.MoveCall(wg)
	// go sui.SetRandom(wg)
	wg.Wait()

}
