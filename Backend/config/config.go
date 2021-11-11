package config

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	//Gitee
	AccessToken string
	Owner       string
	Repo        string
)

func init() {
	file, err := ini.Load("config/config_debug.ini")
	if err != nil {
		fmt.Println("Load config file error!", err)
	}
	LoadServer(file)
}

//加载服务器设置
func LoadServer(file *ini.File) {
	AccessToken = file.Section("gitee").Key("AccessToken").MustString("")
	Owner = file.Section("gitee").Key("Owner").MustString("")
	Repo = file.Section("gitee").Key("Repo").MustString("")
}
