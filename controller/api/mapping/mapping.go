package mapping

import (
    "github.com/gin-gonic/gin"
    "github.com/liCells/controller/api/common/response"
    "github.com/liCells/controller/global"
)

var extensions = map[string]global.Mapping{}

type MappingApi struct{}

func (r *MappingApi) GetMapping(c *gin.Context) {
    for _, extension := range global.Config.Scripts {
        extensions[extension.EsIndexSetting.Name] = extension.EsIndexSetting.Mapping
    }
    for _, extension := range global.Config.Services {
        extensions[extension.EsIndexSetting.Name] = extension.EsIndexSetting.Mapping
    }
    response.SuccessWithData(c, extensions)
}
