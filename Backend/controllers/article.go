package controllers

import (
	"fmt"
	"linblog/repository"
	"net/http"
	"strconv"

	"github.com/Codexiaoyi/linweb"
	"github.com/Codexiaoyi/linweb/interfaces"
)

type ArticleController struct {
	ArticleRepo *repository.ArticleRepository
}

//[GET("/post/list")]
func (a *ArticleController) GetHomeArticles(c interfaces.IContext) {
	page, _ := strconv.Atoi(c.Request().Query("page"))
	size, _ := strconv.Atoi(c.Request().Query("size"))
	if page == 0 {
		page = 1
	}
	if size == 0 {
		size = 5
	}
	articles, total, err := a.ArticleRepo.GetAllArticles(page, size)
	if err != nil {
		Response(c, http.StatusInternalServerError, nil)
		return
	}
	response := &articleListResponseDto{
		Total:       total,
		Items:       make([]*articleResponseDto, 0, len(articles)),
		HasNextPage: total >= page*size,
		Page:        page,
		Size:        size,
	}
	for _, article := range articles {
		dto := &articleResponseDto{}
		err := linweb.NewModel(article).MapToByFieldName(dto).ModelError()
		if err != nil {
			Response(c, http.StatusInternalServerError, nil)
			return
		}
		response.Items = append(response.Items, dto)
	}
	Response(c, http.StatusOK, response)
}

//[GET("/article/:cate/:title")]
func (a *ArticleController) GetArticle(c interfaces.IContext) {
	cate := c.Request().Param("cate")
	title := c.Request().Param("title")
	article, err := a.ArticleRepo.GetArticleContent(cate, title)
	if err != nil {
		Response(c, http.StatusInternalServerError, nil)
		return
	}
	fmt.Println(article)
	Response(c, http.StatusOK, article)
}

type articleListResponseDto struct {
	Total       int                   `json:"total"`
	Items       []*articleResponseDto `json:"items"`
	HasNextPage bool                  `json:"hasNextPage"`
	Page        int                   `json:"page"`
	Size        int                   `json:"size"`
}

type articleResponseDto struct {
	Id            int    `json:"id"`
	IsTop         bool   `json:"isTop"`
	Banner        string `json:"banner"`
	IsHot         bool   `json:"isHot"`
	PubTime       int    `json:"pubTime"`
	Title         string `json:"title"`
	Summary       string `json:"summary"`
	Category      string `json:"category"`
	ViewsCount    int    `json:"viewsCount"`
	CommentsCount int    `json:"commentsCount"`
}
