package controller

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Auth(c *gin.Context) {
	defer func() {
		if err := recover();err != nil {
			log.Errorln("Recover failed:",err)
			return
		}
	}()
	log.Debugln(c.Params)
	c.JSON(200, gin.H{
		"message": "auth",
	})
}
