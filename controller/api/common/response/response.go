package response

import (
    "github.com/gin-gonic/gin"
)

const (
    ERROR      = 500
    SUCCESS    = 200
    SuccessMsg = "success"
)

func BadRequest(c *gin.Context) {
    c.JSON(400, Response{
        Code: 400,
        Msg:  "Bad request",
    })
}

func FailWithMessage(c *gin.Context, msg string) {
    c.JSON(ERROR, Response{
        Code: ERROR,
        Msg:  msg,
    })
}

func FailWithCodeAndMessage(c *gin.Context, code int, msg string) {
    c.JSON(code, Response{
        Code: code,
        Msg:  msg,
    })
}

func SuccessWithMessage(c *gin.Context, msg string) {
    c.JSON(SUCCESS, Response{
        Code: SUCCESS,
        Msg:  msg,
    })
}

func SuccessWithDataAndMessage(c *gin.Context, msg string, data interface{}) {
    c.JSON(SUCCESS, Response{
        Code: SUCCESS,
        Msg:  msg,
        Data: data,
    })
}

func SuccessWithData(c *gin.Context, data interface{}) {
    c.JSON(SUCCESS, Response{
        Code: SUCCESS,
        Msg:  SuccessMsg,
        Data: data,
    })
}

func Success(c *gin.Context) {
    c.JSON(SUCCESS, Response{
        Code: SUCCESS,
        Msg:  SuccessMsg,
    })
}

type Response struct {
    Code int         `json:"code"`
    Msg  string      `json:"msg"`
    Data interface{} `json:"data"`
}
