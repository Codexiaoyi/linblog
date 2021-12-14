package controllers

import (
	"io/ioutil"

	"github.com/Codexiaoyi/linweb/interfaces"
)

type VideoController struct {
}

//[GET("/video/:name")]
func (video *VideoController) GetVideo(c interfaces.IContext) {
	file, _ := ioutil.ReadFile("/home/code/" + c.Request().Param("name"))
	c.Response().Data(200, file)
}
