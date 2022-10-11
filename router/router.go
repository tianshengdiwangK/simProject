package main

import (
	"github.com/gin-gonic/gin"
	setup "github.com/tianshengdiwangK/simProject/log"
)

func init() {
	setup.InitLogger()
}
func main() {
	router := gin.New()
	router.Use(setup.GinLogger(setup.CwLog()), setup.GinRecovery(setup.CwLog(), true))
	router.GET("/", func(context *gin.Context) {
		context.JSON(200, "Hello Gin")
	})
	err := router.Run(":8008")
	if err != nil {
		return
	}
}
