package api

import "github.com/gin-gonic/gin"

type Router interface {
	Register(router *gin.RouterGroup)
}
