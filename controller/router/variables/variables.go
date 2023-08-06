package variables

import (
    "github.com/gin-gonic/gin"
    "github.com/liCells/controller/api"
)

type RouterGroup struct {
    ApiRouter
}
type ApiRouter struct{}

func (s *ApiRouter) InitApiRouter(Router *gin.RouterGroup) {
    variablesApi := api.ApiGroup.VariablesApi
    variablesRouter := Router.Group("variables")
    {
        variablesRouter.POST("get", variablesApi.Get)
        variablesRouter.POST("put", variablesApi.Put)
    }
}
