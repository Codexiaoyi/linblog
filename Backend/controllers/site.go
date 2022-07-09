package controllers

import (
	"linblog/repository"
	"net/http"

	"github.com/Codexiaoyi/linweb"
	"github.com/Codexiaoyi/linweb/interfaces"
)

type SiteController struct {
	SiteRepo *repository.SiteRepository
}

//[GET("/social")]
func (site *SiteController) GetSocials(c interfaces.IContext) {
	socials := site.SiteRepo.GetSocials()
	response := make([]*socialResponseDto, 0, len(socials))
	for _, social := range socials {
		dto := &socialResponseDto{}
		err := linweb.NewModel(social).MapToByFieldName(dto).ModelError()
		if err != nil {
			Response(c, http.StatusInternalServerError, nil)
			return
		}
		response = append(response, dto)
	}
	Response(c, http.StatusOK, response)
}

type siteResponseDto struct {
	Avatar string `json:"avatar"`
	Slogan string `json:"slogan"`
	Name   string `json:"name"`
	Domain string `json:"domain"`
	Notice string `json:"notice"`
	Desc   string `json:"desc"`
}

type socialResponseDto struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Icon  string `json:"icon"`
	Color string `json:"color"`
	Href  string `json:"href"`
}
