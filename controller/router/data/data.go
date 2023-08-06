package data

import (
    "github.com/gin-gonic/gin"
    "github.com/liCells/controller/api"
)

type RouterGroup struct {
    ApiRouter
}
type ApiRouter struct{}

func (s *ApiRouter) InitApiRouter(Router *gin.RouterGroup) {
    dataApi := api.ApiGroup.DataApi
    dataRouter := Router.Group("data")
    {
        dataRouter.GET("doc/get/data", dataApi.GetData)
        dataRouter.POST("doc/create/index", dataApi.CreateIndex)
        dataRouter.POST("doc/save/data", dataApi.SaveData)
        dataRouter.POST("doc/update/data", dataApi.UpdateData)
        dataRouter.POST("doc/delete/data", dataApi.DeleteData)
        dataRouter.POST("basic/search", dataApi.BasicSearch)
    }
}
