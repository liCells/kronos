package initialize

import (
    "github.com/gin-gonic/gin"
    "github.com/liCells/controller/router"
    "net/http"
)

func Routers() *gin.Engine {
    Router := gin.Default()

    Router.Use(cors())

    registry := router.RouterGroupApp.Registry
    variables := router.RouterGroupApp.Variables
    data := router.RouterGroupApp.Data
    mapping := router.RouterGroupApp.Mapping
    extension := router.RouterGroupApp.Extension
    DefaultGroup := Router.Group("")
    {
        registry.InitApiRouter(DefaultGroup)
        variables.InitApiRouter(DefaultGroup)
        data.InitApiRouter(DefaultGroup)
        mapping.InitApiRouter(DefaultGroup)
        extension.InitApiRouter(DefaultGroup)
    }
    return Router
}

func cors() gin.HandlerFunc {
    return func(c *gin.Context) {
        method := c.Request.Method
        origin := c.Request.Header.Get("Origin")
        c.Header("Access-Control-Allow-Origin", origin)
        c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token,Authorization,Token,X-Token,X-User-Id")
        c.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS,DELETE,PUT")
        c.Header("Access-Control-Expose-Headers", "Content-Length,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Content-Type")
        c.Header("Access-Control-Allow-Credentials", "true")

        // 放行所有OPTIONS方法
        if method == "OPTIONS" {
            c.AbortWithStatus(http.StatusNoContent)
        }
        // 处理请求
        c.Next()
    }
}
