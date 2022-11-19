/**
 ********************************************************************************************
 * Created by run-code.
 * User: shijl
 * Date: 2018/12/04
 * Time: 10:18
 ********************************************************************************************
 */

package service

import (
	"errors"

	"run-code/util"
)

var (
	ErrParamError = util.AppErrorNew(errors.New("params error"), util.ERROR)
)
