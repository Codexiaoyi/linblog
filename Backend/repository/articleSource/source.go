package articleSource

import (
	"io/ioutil"
	"linblog/config"
	"linblog/model"
	"net/http"
)

var Article ArticleSource = NewGitee()

type ArticleSource interface {
	GetCategories() ([]string, error)
	GetArticleNames(category string) ([]string, error)
	GetArticleHtml(category string, articleName string) (string, error)
	GetArticleInfo(category string, articleName string) (*model.Article, error)
	GetImageUrl(category string, articleName string, imageName string) (string, error)
}

func DoRequest(url string) ([]byte, error) {
	req, _ := http.NewRequest("GET", url+"?access_token="+config.AccessToken, nil)
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()                // 这步是必要的，防止以后的内存泄漏，切记
	body, err := ioutil.ReadAll(response.Body) // 读取响应 body, 返回为 []byte
	if err != nil {
		return nil, err
	}
	return body, nil
}
