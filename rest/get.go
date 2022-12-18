package rest

import (
	"fmt"

	"github.com/donggyuLim/suino-server/db"
	"github.com/gin-gonic/gin"
)

func RecentBettingList(c *gin.Context) {
	list, err := db.Find("betting", "", "", 10)
	if err != nil {
		fmt.Println(err)
		c.String(404, err.Error())
		return
	}

	c.JSON(200, list)
}

func UserBettingList(c *gin.Context) {
	user := c.Param("user")
	list, err := db.Find("betting", "sender", user, 0)
	if err != nil {
		fmt.Println(err)
		c.String(404, err.Error())
		return
	}

	c.JSON(200, list)
}
