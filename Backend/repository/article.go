package repository

import (
	"linblog/model"
	"linblog/repository/articleSource"
)

type ArticleRepository struct {
}

func (article *ArticleRepository) GetAllArticles(page, size int) ([]*model.Article, int, error) {
	//先拿到所有的分类
	cates, err := articleSource.Article.GetCategories()
	if err != nil {
		return nil, 0, err
	}
	//按照所有分类拿到所有的文章名
	res := make([]*model.Article, 0)
	min_index, max_index := (page-1)*size, page*size
	totalLength := 0
	for _, cate := range cates {
		articles, err := articleSource.Article.GetArticleNames(cate)
		if err != nil {
			return nil, 0, err
		}
		articleLength := len(articles)
		//小于所需起始条目的直接continue
		if articleLength+totalLength < min_index {
			totalLength += articleLength
			continue
		}

		for _, article := range articles {
			if totalLength >= max_index {
				return res, totalLength, nil
			}
			//每个分类对应的文章
			if totalLength >= min_index && totalLength < max_index {
				newArticle, _ := articleSource.Article.GetArticleInfo(cate, article)
				res = append(res, newArticle)
			}
			totalLength++
		}
	}
	return res, totalLength, nil
}

func (article *ArticleRepository) GetArticlesByCategory(cate string, page, size int) ([]*model.Article, int, error) {
	//按照所有分类拿到所有的文章名
	res := make([]*model.Article, 0)
	min_index, max_index := (page-1)*size, page*size
	totalLength := 0
	articles, err := articleSource.Article.GetArticleNames(cate)
	if err != nil {
		return nil, 0, err
	}
	for _, article := range articles {
		if totalLength >= max_index {
			return res, totalLength, nil
		}
		//每个分类对应的文章
		if totalLength >= min_index && totalLength < max_index {
			newArticle, _ := articleSource.Article.GetArticleInfo(cate, article)
			res = append(res, newArticle)
		}
		totalLength++
	}
	return res, totalLength, nil
}

func (article *ArticleRepository) GetArticleContent(category, articleName string) (string, error) {
	return articleSource.Article.GetArticleHtml(category, articleName)
}

func (article *ArticleRepository) GetArticleInfo(category string, articleName string) (*model.Article, error) {
	return articleSource.Article.GetArticleInfo(category, articleName)
}
