package extension

import (
    "github.com/gin-gonic/gin"
    "github.com/liCells/controller/api"
)

type RouterGroup struct {
    ApiRouter
}
type ApiRouter struct{}

func (s *ApiRouter) InitApiRouter(Router *gin.RouterGroup) {
    extensionApi := api.ApiGroup.ExtensionApi
    extensionRouter := Router.Group("script")
    {
        extensionRouter.GET("execute", extensionApi.Execute)
    }
}
