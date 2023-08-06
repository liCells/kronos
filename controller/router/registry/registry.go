package registry

import (
    "github.com/gin-gonic/gin"
    "github.com/liCells/controller/api"
)

type RouterGroup struct {
    ApiRouter
}
type ApiRouter struct{}

func (s *ApiRouter) InitApiRouter(Router *gin.RouterGroup) {
    registryApi := api.ApiGroup.RegistryApi
    registryRouter := Router.Group("registry")
    {
        registryRouter.POST("register", registryApi.Register)
        registryRouter.GET("getApplications", registryApi.GetApplications)
    }
}
