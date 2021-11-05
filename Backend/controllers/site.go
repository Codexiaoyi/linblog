package controllers

import (
	"net/http"

	"github.com/Codexiaoyi/linweb/interfaces"
)

type SiteController struct {
}

//[GET("/site")]
func (site *SiteController) GetSiteInfo(c interfaces.IContext) {
	response := &SiteResponseDto{
		Avatar: "https://s2.ax1x.com/2020/01/17/1SCadg.png",
		Slogan: "The way up is not crowded, and most chose ease.",
		Name:   "FZY′blog",
		Domain: "https://www.fengziy.cn",
		Notice: "本博客的Demo数据由Mockjs生成",
		Desc:   "一个It技术的探索者",
	}
	Response(c, http.StatusOK, response)
}

type SiteResponseDto struct {
	Avatar string `json:"avatar"`
	Slogan string `json:"slogan"`
	Name   string `json:"name"`
	Domain string `json:"domain"`
	Notice string `json:"notice"`
	Desc   string `json:"desc"`
}
