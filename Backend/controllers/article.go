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
	CategoryRepo *repository.CategoryRepository
	ArticleRepo  *repository.ArticleRepository
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
	articles, total, err := a.ArticleRepo.GetArticles(page, size)
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

//[GET("/category/:cate")]
func (a *ArticleController) GetCategoryArticles(c interfaces.IContext) {
	fmt.Println(c.Request().Param("cate"))
}

//[GET("/category")]
func (article *ArticleController) GetCategories(c interfaces.IContext) {
	cates := article.CategoryRepo.GetCategories()
	response := make([]*categoryResponseDto, 0, len(cates))
	for _, category := range cates {
		dto := &categoryResponseDto{}
		err := linweb.NewModel(category).MapToByFieldName(dto).ModelError()
		if err != nil {
			Response(c, http.StatusInternalServerError, nil)
			return
		}
		response = append(response, dto)
	}
	Response(c, http.StatusOK, response)
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
	Content       string `json:"content"`
	ViewsCount    int    `json:"viewsCount"`
	CommentsCount int    `json:"commentsCount"`
}

type categoryResponseDto struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Href  string `json:"href"`
}
