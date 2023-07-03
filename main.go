package main

import (
	db "./dbs"
	"./libs"
	"./routers"
	"github.com/gin-gonic/gin"
)

func main()  {
	defer db.Conns.Close()
	gin.SetMode(gin.DebugMode)
	router := routers.InitRouter()
	router.Run(":" + libs.Conf.Read("site", "httpport"))
}
