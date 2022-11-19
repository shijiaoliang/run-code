/**
 ********************************************************************************************
 * Created by run-code.
 * User: shijl
 * Date: 2018/12/04
 * Time: 10:18
 ********************************************************************************************
 */

package util

import "os"

// 目录 or 文件 是否存在
func FileExist(file string) bool {
	_, err := os.Stat(file)
	if err != nil && os.IsNotExist(err) {
		return false
	}

	return true
}
