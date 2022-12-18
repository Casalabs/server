package rest

// import (
// 	"fmt"

// 	"github.com/donggyuLim/suino-server/db"
// 	"github.com/donggyuLim/suino-server/rpc"
// 	"github.com/donggyuLim/suino-server/types"
// 	"github.com/gin-gonic/gin"
// )

// func FlipGame(c *gin.Context) {
// 	r := types.GameRequest{}
// 	err := c.ShouldBindJSON(&r)
// 	if err != nil {
// 		c.String(400, err.Error())
// 	}
// 	tx := r.Tx

// 	flipResonse := rpc.GetFlipEvent(tx) //websocket -> send
// 	var flip types.FlipEventResponse
// 	err = db.FindOne("suino", "betting", "transaction", tx, flip)
// 	fmt.Println(flip)
// 	if err == nil {
// 		c.String(404, "Exsists")
// 		return
// 	}

// 	db.Insert("suino", "betting", flipResonse)
// 	c.JSON(200, flipResonse)

// }
