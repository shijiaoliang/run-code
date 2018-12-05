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

// CodeRunDto
type CodeRunDto struct {
	Language string `json:"language"`
	Stdin    string `json:"stdin"`
	Argv     string `json:"argv"`
	Script   string `json:"script"`
}

// DockerRunInDto
type DockerRunInDto struct {
	CodeData string `json:"code_data"`
	Language string `json:"language"`
}

// docker run 返回结构体
type DockerRunOutDto struct {
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
	Error  string `json:"error"`
}
