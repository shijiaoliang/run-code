/**
 *******************************************51cto********************************************
 * Copyright (c)  www.51cto.com
 * Created by run-code.
 * User: shijl@51cto.com
 * Date: 2018/12/03
 * Time: 10:18
 ********************************************************************************************
 */

package run

import (
	"os/exec"
	"strings"
	"time"

	"github.com/json-iterator/go"
)

// 运行容器 & 获取计算结果 & 销毁容器
func DockerRun(d DockerRunInDto) (runOut string, err error) {
	codeData := d.CodeData
	language := d.Language

	// 超时/单位秒
	timeout := 30

	// 最大内存 20M
	maxMemory := "20m"

	chMaxTime := make(chan bool)
	go func() {
		cmd := exec.Command("docker", "run", "-m", maxMemory, "-i", "--rm", "glot/"+language)
		cmd.Stdin = strings.NewReader(codeData)
		out, e := cmd.Output()

		runOut = strings.Replace(string(out), "\n", "", -1)
		err = e

		// 写入chan
		chMaxTime <- true
	}()

	select {
	case <-chMaxTime:
		//正常执行

	case <-time.After(time.Duration(timeout) * time.Second):
		//超时

		dockerRunOutDto := &DockerRunOutDto{
			Stdout: "",
			Stderr: "",
			Error:  ErrRunTimeOut.Error(),
		}

		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		b, _ := json.Marshal(&dockerRunOutDto)

		runOut = string(b)
	}

	return
}
