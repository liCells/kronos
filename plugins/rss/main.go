package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mmcdole/gofeed"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	pluginName        = "Rss_Pull"
	indexName         = "rss_pull"
	saveDataUrlSuffix = "/data/doc/save/data"
	contentType       = "application/json"
	serverUrl         = "http://localhost:18000"
)

func main() {
	args := os.Args
	// 1: symbol
	if len(args) != 3 {
		// usage
		fmt.Println("usage: <config_file_path> <symbol>")
		return
	}

	config := parseConfiguration(args[1])

	fp := gofeed.NewParser()
	data := config.Resources

	if len(data[args[2]]) == 0 {
		return
	}

	for _, url := range data[args[2]] {
		feed, _ := fp.ParseURL(url)
		for _, item := range feed.Items {
			pluginData := pluginData{
				Name:  pluginName,
				Index: indexName,
				Id:    item.Link,
				Data: &rssPullData{
					SourceUrl: item.Link,
					Title:     item.Title,
					Content:   item.Content,
					Tag:       feed.Title,
				},
			}
			saveData(pluginData)
		}
	}
}

func saveData(pluginData pluginData) {
	jsonData, err := json.Marshal(pluginData)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(serverUrl+saveDataUrlSuffix, contentType, bytes.NewBuffer(jsonData))

	res := response{}
	body, _ := io.ReadAll(resp.Body)
	_ = json.Unmarshal(body, &res)
}

func parseConfiguration(path string) configuration {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Config file read error")
		panic(err)
	}
	configJson, _ := io.ReadAll(file)
	config := configuration{}
	err = json.Unmarshal(configJson, &config)
	if err != nil {
		panic("Json parse error")
	}
	if config.ServerUrl != "" {
		serverUrl = config.ServerUrl
	}
	if config.SaveDataUrlSuffix != "" {
		saveDataUrlSuffix = config.SaveDataUrlSuffix
	}
	return config
}

type configuration struct {
	ServerUrl         string              `json:"serverUrl"`
	SaveDataUrlSuffix string              `json:"saveDataUrlSuffix"`
	Resources         map[string][]string `json:"resources"`
}

type rssPullData struct {
	SourceUrl string `json:"source_url"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Tag       string `json:"tag"`
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
