package main

import (
	"linblog/controllers"
	"linblog/middlewares"
	"linblog/repository"

	"github.com/Codexiaoyi/linweb"
)

func main() {
	linweb := linweb.NewLinWeb()
	linweb.AddSingleton(&repository.SiteRepository{})
	linweb.AddSingleton(&repository.ArticleRepository{}, &repository.CategoryRepository{})
	linweb.AddMiddlewares(middlewares.Cors, middlewares.Logs)
	linweb.AddControllers(&controllers.SiteController{}, &controllers.ArticleController{}, &controllers.CategoryController{}, &controllers.VideoController{})
	linweb.Run(":5002")
}
