package apis

import "github.com/gin-gonic/gin"

type IAPI interface {
	RegisterRoutes(router *gin.RouterGroup)
}
