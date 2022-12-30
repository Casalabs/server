package rest

import (
	"sync"

	"github.com/donggyuLim/suino-server/p2p"
	"github.com/gin-gonic/gin"
)

const port = ":8080"

func Start(wg *sync.WaitGroup) {
	r := gin.Default()
	gin.ForceConsoleColor()
	r.GET("/", RecentBettingList)
	r.GET("/user/:user", UserBettingList)
	r.GET("/ws", p2p.Upgrade)
	r.Run(port)
	defer wg.Done()
}
