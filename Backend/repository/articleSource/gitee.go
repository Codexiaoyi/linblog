package articleSource

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"linblog/config"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
)

var _ ArticleSource = &Gitee{}

type Gitee struct {
	access_token string
	owner        string
	repo         string
}

type GiteeResponse struct {
	Content     string `json:"content"`
	DownloadUrl string `json:"download_url"`
	Encoding    string `json:"encoding"`
	HtmlUrl     string `json:"html_url"`
	Name        string `json:"name"`
	Path        string `json:"path"`
	Sha         string `json:"sha"`
	Size        int    `json:"size"`
	Type        string `json:"type"`
	Url         string `json:"url"`
}

func NewGitee() ArticleSource {
	return &Gitee{
		access_token: config.AccessToken,
		owner:        config.Owner,
		repo:         config.Repo,
	}
}

func (g *Gitee) GetCategories() ([]string, error) {
	response, err := DoRequest(fmt.Sprintf("https://gitee.com/api/v5/repos/%s/%s/contents/", g.owner, g.repo))
	if err != nil {
		return nil, err
	}
	categories := []*GiteeResponse{}
	err = json.Unmarshal(response, &categories)
	if err != nil {
		return nil, err
	}
	res := make([]string, 0, len(categories))
	for _, category := range categories {
		res = append(res, category.Name)
	}
	return res, nil
}

func (g *Gitee) GetArticleNames(category string) ([]string, error) {
	response, err := DoRequest(fmt.Sprintf("https://gitee.com/api/v5/repos/%s/%s/contents/%s", g.owner, g.repo, category))
	if err != nil {
		return nil, err
	}
	articles := []*GiteeResponse{}
	err = json.Unmarshal(response, &articles)
	if err != nil {
		return nil, err
	}
	res := make([]string, 0, len(articles))
	for _, article := range articles {
		res = append(res, article.Name)
	}
	return res, nil
}

func (g *Gitee) GetArticleHtml(category string, articleName string) (string, error) {
	//直接获取相应的文章markdown字符串
	response, err := DoRequest(fmt.Sprintf("https://gitee.com/api/v5/repos/%s/%s/contents/%s/%s/%s.md", g.owner, g.repo, category, articleName, articleName))
	if err != nil {
		return "", err
	}
	article := &GiteeResponse{}
	err = json.Unmarshal(response, &article)
	if err != nil {
		return "", err
	}
	mdStr, err := base64.StdEncoding.DecodeString(article.Content)
	if err != nil {
		return "", err
	}

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	renderImageFunc := func(w io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {
		//这里将md中的图片在转换为html的时候替换
		image, ok := node.(*ast.Image)
		if !ok {
			return ast.GoToNext, false
		}
		name := string(image.Destination)
		url, err := g.GetImageUrl(category, articleName, name)
		if err != nil {
			return ast.GoToNext, false
		}
		image.Destination = []byte(url)
		return ast.GoToNext, false
	}

	opts := html.RendererOptions{Flags: htmlFlags, RenderNodeHook: renderImageFunc}
	renderer := html.NewRenderer(opts)
	html := markdown.ToHTML(mdStr, nil, renderer)
	return string(html), nil
}

func (g *Gitee) GetImageUrl(category string, articleName string, imageName string) (string, error) {
	return fmt.Sprintf("https://gitee.com/%s/%s/raw/master/%s/%s/%s", g.owner, g.repo, category, articleName, imageName), nil
}
