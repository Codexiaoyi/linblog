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

	Avatar string
	Slogan string
	Name   string
	Domain string
	Notice string
	Desc   string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("Load config file error!", err)
	}
	LoadSource(file)
	LoadSiteInfo(file)
}

//加载数据资源配置
func LoadSource(file *ini.File) {
	AccessToken = file.Section("gitee").Key("AccessToken").MustString("")
	Owner = file.Section("gitee").Key("Owner").MustString("")
	Repo = file.Section("gitee").Key("Repo").MustString("")
}

//加载站点基本信息
func LoadSiteInfo(file *ini.File) {
	Avatar = file.Section("site").Key("Avatar").MustString("")
	Slogan = file.Section("site").Key("Slogan").MustString("")
	Name = file.Section("site").Key("Name").MustString("")
	Domain = file.Section("site").Key("Domain").MustString("")
	Notice = file.Section("site").Key("Notice").MustString("")
	Desc = file.Section("site").Key("Desc").MustString("")
}
