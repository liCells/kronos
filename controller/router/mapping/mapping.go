package mapping

import (
    "github.com/gin-gonic/gin"
    "github.com/liCells/controller/api"
)

type RouterGroup struct {
    ApiRouter
}
type ApiRouter struct{}

func (s *ApiRouter) InitApiRouter(Router *gin.RouterGroup) {
    mappingApi := api.ApiGroup.MappingApi
    mappingRouter := Router.Group("mapping")
    {
        mappingRouter.GET("get", mappingApi.GetMapping)
    }
}
