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
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"run-code/log"
)

//appError 系统通用error
type AppError struct {
	Err  error
	Code int
}

func (a *AppError) Error() string {
	return a.Err.Error()
}

func (a *AppError) GetCode() int {
	return a.Code
}

func AppErrorNew(err error, code int) *AppError {
	return &AppError{
		err,
		code,
	}
}

func CheckErr(err error, c *gin.Context) {
	if err != nil {
		//记录日志
		log.Log.Error(fmt.Sprintf("%+v", err))

		//res abort
		msg := ""

		switch err := errors.Cause(err).(type) {
		case *AppError:
			msg = err.Error()
		default:
			msg = "system error"
		}

		c.AbortWithStatusJSON(http.StatusOK, &Res{"", ERROR, msg})
		return
	}
}
