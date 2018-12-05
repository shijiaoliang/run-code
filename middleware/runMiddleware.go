/**
 *******************************************51cto********************************************
 * Copyright (c)  www.51cto.com
 * Created by run-code.
 * User: shijl@51cto.com
 * Date: 2018/12/03
 * Time: 10:18
 ********************************************************************************************
 */

package middleware

import (
	"github.com/gin-gonic/gin"

	"run-code/config"
	"run-code/util"
)

func RunCodeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		runAuthToken := c.DefaultQuery("run_auth_token", "")

		if runAuthToken != config.AppConf.RunAuthToken {
			util.ResError(c, "invalid token", util.ERROR, "")
			c.Abort()
			return
		}

		c.Next()
	}
}
