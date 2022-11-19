/**
 ********************************************************************************************
 * Created by run-code.
 * User: shijl
 * Date: 2018/12/04
 * Time: 10:18
 ********************************************************************************************
 */

package run

import (
	"github.com/json-iterator/go"
	"github.com/pkg/errors"
)

type language struct {
	Language     string `json:"language"`
	LanguageName string `json:"language_name"`
	Ext          string `json:"ext"`
	Cmd          string `json:"cmd"`
}

var (
	LanguageMap = map[string]language{
		"java": {
			Language:     "java",
			LanguageName: "java",
			Ext:          "java",
			Cmd:          "",
		},
		"c": {
			Language:     "c",
			LanguageName: "c",
			Ext:          "c",
			Cmd:          "",
		},
		"cpp": {
			Language:     "cpp",
			LanguageName: "c++",
			Ext:          "cpp",
			Cmd:          "",
		},
		"php": {
			Language:     "php",
			LanguageName: "php",
			Ext:          "php",
			Cmd:          "php main.php",
		},
		"python": {
			Language:     "python",
			LanguageName: "python",
			Ext:          "py",
			Cmd:          "python main.python",
		},
		"lua": {
			Language:     "lua",
			LanguageName: "lua",
			Ext:          "lua",
			Cmd:          "lua main.lua",
		},
		"go": {
			Language:     "go",
			LanguageName: "go",
			Ext:          "go",
			Cmd:          "go run main.go",
		},
		"csharp": {
			Language:     "csharp",
			LanguageName: "c#",
			Ext:          "cs",
			Cmd:          "",
		},
		"scala": {
			Language:     "scala",
			LanguageName: "scala",
			Ext:          "scala",
			Cmd:          "",
		},
		"javascript": {
			Language:     "javascript",
			LanguageName: "javascript",
			Ext:          "js",
			Cmd:          "node main.js",
		},
		"kotlin": {
			Language:     "kotlin",
			LanguageName: "kotlin",
			Ext:          "kt",
			Cmd:          "",
		},
		"swift": {
			Language:     "swift",
			LanguageName: "swift",
			Ext:          "swift",
			Cmd:          "",
		},
	}
)

// Languages
func Languages() (languages map[string]string) {
	for k, v := range LanguageMap {
		languages[k] = v.LanguageName
	}

	return
}

// infoLanguage
func infoLanguage(lang string) (l language, err error) {
	l, ok := LanguageMap[lang]
	if !ok {
		err = errors.WithStack(ErrLanguageNotSupport)
	}

	return
}

// 生成 DockerRunInDto
func GenerateDockerRunInDto(c *CodeRunDto) (d DockerRunInDto, err error) {
	//e.g: "{\"files\":[{\"name\":\"main.php\",\"content\":\"<?php\n\necho \\"Hello World\\";\"}],\"stdin\":\"1\n2\",\"command\":\"\"}"

	//语言信息
	language, err := infoLanguage(c.Language)
	if err != nil {
		return
	}

	codeData := make(map[string]interface{})
	codeData["language"] = c.Language

	//stdin
	if c.Stdin != "" {
		codeData["stdin"] = c.Stdin
	}

	//command
	if c.Argv != "" {
		codeData["command"] = language.Cmd + " " + c.Argv
	}

	//files
	files := []map[string]string{
		{
			"name":    "main." + language.Ext,
			"content": c.Script,
		},
	}
	codeData["files"] = files

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	b, _ := json.Marshal(&codeData)

	d.CodeData = string(b)
	d.Language = language.Language

	return
}
