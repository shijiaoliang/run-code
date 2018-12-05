/**
 *******************************************51cto********************************************
 * Copyright (c)  www.51cto.com
 * Created by run-code.
 * User: shijl@51cto.com
 * Date: 2018/12/03
 * Time: 10:18
 ********************************************************************************************
 */

package config

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/json-iterator/go"

	"run-code/util"
)

type AppConfig struct {
	Model        string `json:"model"`
	Address      string `json:"address"`
	RunAuthToken string `json:"run_auth_token"`
}

var (
	AppConf *AppConfig
)

func init() {
	AppConf = new(AppConfig)
	AppConf.Model = gin.ReleaseMode
	AppConf.Address = "8989"
	AppConf.RunAuthToken = "123456"

	// 如果存在配置json文件 以配置文件为准
	if util.FileExist("conf/app.json") {
		appFile, _ := os.Open("conf/app.json")
		defer appFile.Close()

		decoder := jsoniter.NewDecoder(appFile)
		err := decoder.Decode(&AppConf)

		if err != nil {
			panic("解析app配置文件失败")
		}
	}
}
