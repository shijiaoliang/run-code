/**
 ********************************************************************************************
 * Created by run-code.
 * User: shijl
 * Date: 2018/12/03
 * Time: 10:18
 ********************************************************************************************
 */

package main

import (
	"time"

	"github.com/gin-gonic/gin"

	"run-code/log"
	"run-code/router"
	"run-code/config"
	"run-code/middleware"
)

func main() {
	r := gin.New()

	//=====config=====
	// config init

	//=====mode=====
	//gin.SetMode(gin.ReleaseMode)
	gin.SetMode(config.AppConf.Model)

	//=====logger=====
	log.ConfigLocalFilesystemLogger(log.Log, "runtime", "app.log", time.Hour*360, time.Hour)

	//=====middleware=====
	r.Use(middleware.LoggerMiddleware(log.Log))
	r.Use(gin.Recovery())

	//=====route=====
	router.InitRunRouter(r)

	//=====run=====
	r.Run(config.AppConf.Address)
}
