package p2p

import (
	"fmt"
	"sync"

	"github.com/donggyuLim/suino-server/utils"
)

type Message struct {
	Method  string      `json:"method"`
	Payload interface{} `json:"payload"`
}

func HandleMsg(wg *sync.WaitGroup, ch chan interface{}) {
	defer wg.Done()
	for {
		message, ok := <-ch
		if !ok {
			fmt.Println("we are done")
			break
		}
		BroadcastNewEvent(message)

	}

}

func BroadcastNewEvent(data interface{}) {
	for _, peer := range Peers.peerMap {
		sendNewBetting("new Bet", data, peer)
	}
}

// Message -> json
func makeMessage(method string, payload interface{}) []byte {
	m := Message{
		Payload: payload,
	}
	return utils.ToJSON(m)
}

func sendNewBetting(method string, data interface{}, p *peer) {
	m := makeMessage(method, data)
	p.inbox <- m
}
