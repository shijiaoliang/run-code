/**
 *******************************************51cto********************************************
 * Copyright (c)  www.51cto.com
 * Created by run-code.
 * User: shijl@51cto.com
 * Date: 2018/12/03
 * Time: 10:18
 ********************************************************************************************
 */

package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//appRes 系统通用返回
const (
	SUCCESS = 1  //成功
	ZERO    = 0  //0值
	ERROR   = -1 //失败
)

const (
	constSuccessMsg = "success"
)

type Res struct {
	Data interface{} `json:"data"`
	Res  int         `json:"res"`
	Msg  string      `json:"msg"`
}

func ResSuccessList(c *gin.Context, dataList interface{}, pagination Pagination, msg string) {
	if msg == "" {
		msg = constSuccessMsg
	}

	c.JSON(http.StatusOK, &Res{
		gin.H{
			"list":       dataList,
			"pagination": pagination,
		},
		SUCCESS,
		msg,
	})
}

func ResSuccess(c *gin.Context, data interface{}, msg string) {
	if msg == "" {
		msg = constSuccessMsg
	}

	c.JSON(http.StatusOK, &Res{
		data,
		SUCCESS,
		msg,
	})
}

func ResError(c *gin.Context, msg string, res int, data interface{}) {
	c.JSON(http.StatusOK, &Res{
		data,
		res,
		msg,
	})
}
