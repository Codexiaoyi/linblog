package main

import (
	"linblog/controllers"
	"linblog/middlewares"

	"github.com/Codexiaoyi/linweb"
)

func main() {
	linweb := linweb.NewLinWeb()
	linweb.AddTransient()
	linweb.AddMiddlewares(middlewares.Cors)
	linweb.AddControllers(&controllers.SiteController{})
	linweb.Run(":5002")
}
