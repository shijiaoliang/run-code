/**
 *******************************************51cto********************************************
 * Copyright (c)  www.51cto.com
 * Created by run-code.
 * User: shijl@51cto.com
 * Date: 2018/12/04
 * Time: 10:18
 ********************************************************************************************
 */

package run

import (
	"errors"

	"run-code/util"
)

var (
	ErrRunTimeOut         = util.AppErrorNew(errors.New("run code timeout"), util.ERROR)
	ErrLanguageNotSupport = util.AppErrorNew(errors.New("language not support"), util.ERROR)
)
