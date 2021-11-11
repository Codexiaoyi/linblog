package repository

import (
	"linblog/model"
	"linblog/repository/articleSource"
)

type CategoryRepository struct {
}

func (cate *CategoryRepository) GetCategories() []*model.Category {
	cates, err := articleSource.Article.GetCategories()
	if err != nil {
		return nil
	}
	res := make([]*model.Category, 0, len(cates))
	for index, cate := range cates {
		res = append(res, &model.Category{Id: index, Title: cate, Href: "category/" + cate})
	}
	return res
}
