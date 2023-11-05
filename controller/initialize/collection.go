package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/liCells/controller/global"
	"strconv"
	"sync"
)

func Initialize() {
	GetConfig()

	InitEsClient()

	router := Routers()

	waitGroup := &sync.WaitGroup{}

	waitGroup.Add(1)
	go initServer(router)

	LoadScheduledTask()

	waitGroup.Wait()
}

func initServer(router *gin.Engine) {
	port := global.Config.Port
	if port == 0 {
		port = 8080
	}
	_ = router.Run(":" + strconv.Itoa(int(port)))
}
