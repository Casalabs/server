package rest

import (
	"sync"
	"time"

	"github.com/donggyuLim/suino-server/p2p"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const port = ":3306"

func Start(wg *sync.WaitGroup) {
	r := gin.Default()
	gin.ForceConsoleColor()

	//CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://suino.io"},
		AllowMethods: []string{"GET", "POST"},
		MaxAge:       12 * time.Hour,
	}))

	r.GET("/", RecentBettingList)
	r.GET("/user/:user", UserBettingList)
	r.GET("/ws", p2p.Upgrade)
	r.Run(port)
	defer wg.Done()
}
