package variables

import (
    "github.com/gin-gonic/gin"
    "github.com/liCells/controller/api/common/response"
)

type VariablesApi struct{}

var (
    plugins = make(map[string]map[string]any)
)

func (r *VariablesApi) Get(c *gin.Context) {
    pluginData := getPluginData{}
    err := c.ShouldBindJSON(&pluginData)
    if err != nil {
        response.BadRequest(c)
        return
    }

    data := plugins[pluginData.Name]
    if data == nil {
        response.SuccessWithMessage(c, "Not found data")
        return
    }
    if pluginData.Data == nil {
        response.SuccessWithData(c, data)
        return
    }
    res := make(map[string]any)
    for _, v := range pluginData.Data {
        if data[v] == nil {
            continue
        }
        res[v] = data[v]
    }
    response.SuccessWithData(c, res)
}

func (r *VariablesApi) Put(c *gin.Context) {
    pluginData := putPluginData{}
    err := c.ShouldBindJSON(&pluginData)
    if err != nil {
        response.BadRequest(c)
        return
    }
    if pluginData.Data == nil {
        response.Success(c)
        return
    }
    if plugins[pluginData.Name] == nil {
        plugins[pluginData.Name] = make(map[string]any)
    }
    for k, v := range pluginData.Data {
        plugins[pluginData.Name][k] = v
    }
    response.Success(c)
}

type getPluginData struct {
    Name string   `json:"name"`
    Data []string `json:"data"`
}

type putPluginData struct {
    Name string         `json:"name"`
    Data map[string]any `json:"data"`
}
