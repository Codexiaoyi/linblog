package middlewares

import (
	"fmt"

	"github.com/Codexiaoyi/linweb/interfaces"
)

func Logs(context interfaces.IContext) {
	fmt.Println("Request method: ", context.Request().Method())
	fmt.Println("Request path: ", context.Request().Path())
	fmt.Println("Request body: ", context.Request().Body())
	//处理请求
	context.Next()
}
