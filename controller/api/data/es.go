package data

import (
    "context"
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/liCells/controller/api/common/response"
    "github.com/liCells/controller/global"
)

func (r *DataApi) CreateIndex(c *gin.Context) {
    pluginIndex := pluginIndex{}
    err := c.ShouldBindJSON(&pluginIndex)
    if err != nil {
        response.BadRequest(c)
        return
    }
    _, err = global.ESClient.CreateIndex(pluginIndex.Index).BodyJson(pluginIndex.Data).Do(context.Background())
    if err != nil {
        response.FailWithMessage(c, "Create failure")
        return
    }
    response.Success(c)
}

func (r *DataApi) GetData(c *gin.Context) {
    pluginData := pluginData{}
    err := c.ShouldBindJSON(&pluginData)
    if err != nil {
        response.BadRequest(c)
        return
    }
    res, err := global.ESClient.Get().
        Index(pluginData.Index).
        Id(pluginData.Id).
        Do(context.Background())

    if err != nil {
        response.FailWithMessage(c, "Get failure")
        return
    }

    fmt.Print(res)
    response.SuccessWithData(c, res.Source)
}

func (r *DataApi) SaveData(c *gin.Context) {
    pluginData := pluginData{}
    err := c.ShouldBindJSON(&pluginData)
    if err != nil {
        response.BadRequest(c)
        return
    }
    _, err = global.ESClient.Index().
        Index(pluginData.Index).
        Id(pluginData.Id).
        BodyJson(pluginData.Data).
        Do(context.Background())

    if err != nil {
        response.FailWithMessage(c, "Save failure")
        return
    }
    response.Success(c)
}

func (r *DataApi) UpdateData(c *gin.Context) {
    pluginData := pluginData{}
    err := c.ShouldBindJSON(&pluginData)
    if err != nil {
        response.BadRequest(c)
        return
    }
    _, err = global.ESClient.Update().
        Index(pluginData.Index).
        Id(pluginData.Id).
        Doc(pluginData.Data).
        Do(context.Background())

    if err != nil {
        response.FailWithMessage(c, "Update failure")
        return
    }
    response.Success(c)
}

func (r *DataApi) DeleteData(c *gin.Context) {
    pluginData := pluginData{}
    err := c.ShouldBindJSON(&pluginData)
    if err != nil {
        response.BadRequest(c)
        return
    }
    _, err = global.ESClient.Delete().
        Index(pluginData.Index).
        Id(pluginData.Id).
        Do(context.Background())

    if err != nil {
        response.FailWithMessage(c, "Delete failure")
        return
    }
    response.Success(c)
}

type pluginIndex struct {
    Name  string      `json:"name"`
    Index string      `json:"index"`
    Data  interface{} `json:"data"`
}

type pluginData struct {
    Name  string      `json:"name"`
    Index string      `json:"index"`
    Id    string      `json:"id"`
    Data  interface{} `json:"data"`
}
