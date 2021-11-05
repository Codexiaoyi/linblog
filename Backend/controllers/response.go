package controllers

import "github.com/Codexiaoyi/linweb/interfaces"

// 前端项目需要，将项目状态码与数据分离
func Response(c interfaces.IContext, httpCode int, data interface{}) {
	c.Response().JSON(httpCode, &struct {
		Code int         `json:"code"`
		Data interface{} `json:"data"`
	}{
		Code: httpCode,
		Data: data,
	})
}
