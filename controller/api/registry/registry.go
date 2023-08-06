package registry

import "github.com/gin-gonic/gin"

type RegistryApi struct{}

func (r *RegistryApi) Register(c *gin.Context) {
}

func (r *RegistryApi) GetApplications(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "pong",
    })
}
