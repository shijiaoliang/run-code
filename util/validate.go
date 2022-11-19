/**
 ********************************************************************************************
 * Created by run-code.
 * User: shijl
 * Date: 2018/12/03
 * Time: 10:18
 ********************************************************************************************
 */

package util

import (
	"gopkg.in/go-playground/validator.v9"
)

// 全局验证实例
var Validate *validator.Validate

func init()  {
	Validate = validator.New()
}
