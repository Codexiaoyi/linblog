package controllers

import (
	"fmt"
	"linblog/repository"
	"net/http"

	"github.com/Codexiaoyi/linweb"
	"github.com/Codexiaoyi/linweb/interfaces"
)

type CategoryController struct {
	CategoryRepo *repository.CategoryRepository
}

//[GET("/category/:cate")]
func (cate *CategoryController) GetCategoryArticles(c interfaces.IContext) {
	fmt.Println(c.Request().Param("cate"))
}

//[GET("/category")]
func (cate *CategoryController) GetCategories(c interfaces.IContext) {
	cates := cate.CategoryRepo.GetCategories()
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

type categoryResponseDto struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Href  string `json:"href"`
}
