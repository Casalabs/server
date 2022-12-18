package p2p

import (
	"fmt"
	"net/http"

	"github.com/donggyuLim/suino-server/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func Upgrade(c *gin.Context) {
	openPort := c.Request.URL.Query().Get("openPort")
	ip := utils.Splitter(c.Request.RemoteAddr, ":", 0)
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return openPort != "" || ip != ""
	}
	fmt.Printf("%s,wants an upgrade\n", openPort)

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	utils.HandleErr(err)
	fmt.Println("Successful websocket")
	peer := initPeer(conn, ip, openPort)
	peer.inbox <- []byte("hello from 3000!")

}

// func BroadcastNewBet(method string, flip types.Flip) {
// 	for _, p := range Peers.peerMap {
// 		sendNewBetting(method, flip, p)
// 	}
// }
