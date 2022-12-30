package main

import (
	"fmt"
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
	wg.Add(3)
	db.DBConnect()
	defer db.Close()
	ch := make(chan interface{}, 10)
	defer close(ch)

	go p2p.HandleMsg(wg, ch)

	go rest.Start(wg)
	go func() {
		defer wg.Done()
		for {
			fmt.Println("Websocket connect")
			sui := sui.NewClient(websocketHost, jsonrpcEndPoint)
			err := sui.HandleMsg(ch)
			if err != nil {
				fmt.Println("Websocket Disconnect")
				fmt.Println(err)
			}
		}
	}()

	wg.Wait()
	// go sui.MoveCall(wg)
	// go sui.SetRandom(wg)

}
