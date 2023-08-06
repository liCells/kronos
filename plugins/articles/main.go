package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"time"
)

var (
	pluginName        = "Articles"
	indexName         = "articles"
	saveDataUrlSuffix = "/data/doc/save/data"
	contentType       = "application/json"
	serverUrl         = "http://localhost:18000"
	port              = "18001"
)

func main() {
	args := os.Args

	if len(args) == 3 {
		serverUrl = args[1]
		port = args[2]
	} else if len(args) == 2 {
		serverUrl = args[1]
	}

	r := gin.Default()
	r.Use(cors())

	r.POST("/save", func(c *gin.Context) {
		var article Article
		if c.ShouldBindJSON(&article) != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "Bad Request"})
			return
		}
		if article.Title == "" {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "Title is empty"})
			return
		}
		if article.Content == "" {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "Content is empty"})
			return
		}
		if article.Url == "" {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "Url is empty"})
			return
		}

		article.RecordTime = time.Now()

		pluginData := pluginData{
			Name:  pluginName,
			Index: indexName,
			Id:    article.Title,
			Data:  article,
		}

		jsonData, err := json.Marshal(pluginData)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "Data conversion failed"})
			return
		}

		resp, err := http.Post(serverUrl+saveDataUrlSuffix, contentType, bytes.NewBuffer(jsonData))
		res := response{}
		body, _ := io.ReadAll(resp.Body)
		_ = json.Unmarshal(body, &res)

		if err != nil && res.Code != 200 {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "Save data failed"})
			return
		}
	})

	err := r.Run(":" + port)
	if err != nil {
		panic(err.Error())
	}
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

type Article struct {
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Type       []string  `json:"type"`
	Url        string    `json:"url"`
	RecordTime time.Time `json:"recordTime"`
}

type response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type pluginData struct {
	Name  string      `json:"name"`
	Index string      `json:"index"`
	Id    string      `json:"id"`
	Data  interface{} `json:"data"`
}
