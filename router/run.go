/**
 ********************************************************************************************
 * Created by run-code.
 * User: shijl
 * Date: 2018/12/03
 * Time: 10:18
 ********************************************************************************************
 */

package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/json-iterator/go"

	"run-code/util"
	"run-code/middleware"
	"run-code/service/run"
	"run-code/service"
)

func InitRunRouter(r *gin.Engine) {
	runRouter := r.Group("/run")

	//run-code & RunCodeMiddleware
	runRouter.POST("run-code", middleware.RunCodeMiddleware(), func(c *gin.Context) {
		dto := new(run.CodeRunDto)
		if err := c.ShouldBindJSON(dto); err != nil {
			err = errors.WithStack(service.ErrParamError)
			util.CheckErr(err, c)
			return
		}

		//docker run
		d, err := run.GenerateDockerRunInDto(dto)
		if err != nil {
			util.CheckErr(err, c)
			return
		}

		runOut, err := run.DockerRun(d)
		if err != nil {
			util.CheckErr(err, c)
			return
		}

		var result run.DockerRunOutDto
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		json.Unmarshal([]byte(runOut), &result)

		util.ResSuccess(c, result, "")
	})

	//languages
	runRouter.GET("languages", func(c *gin.Context) {
		util.ResSuccess(c, run.Languages(), "")
	})
}
