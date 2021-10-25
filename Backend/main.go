package main

import (
	"linblog/controllers"

	"github.com/Codexiaoyi/linweb"
)

func main() {
	linweb := linweb.NewLinWeb()
	linweb.AddControllers(&controllers.HomeController{})
	linweb.Run(":6666")
}
