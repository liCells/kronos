package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
    "net/url"
    "os"
    "strconv"
)

var (
    pluginName            = "Github_Stars"
    indexName             = "github_stars"
    page                  = 1
    perPage               = 50
    pageName              = "page"
    getVariablesUrlSuffix = "/variables/get"
    putVariablesUrlSuffix = "/variables/put"
    saveDataUrlSuffix     = "/data/doc/save/data"
    contentType           = "application/json"
    serverUrl             = "http://localhost:18000"
    proxyUrl              = "http://127.0.0.1:7890"
    client                = &http.Client{}
)

func main() {
    args := os.Args
    // 1: failedFilePath
    // 2: username
    // 3: get var serverUrl
    if len(args) != 5 && len(args) != 4 {
        // usage
        fmt.Println("usage: <failedFilePath> <username> <serverUrl default:http://localhost:18000> <proxy:http://127.0.0.1:7890>")
        fmt.Println("       ./failedFilePath liCells http://192.168.123.200:18000 http://127.0.0.1:7890")
        return
    }
    serverUrl = args[3]
    if len(args) == 5 {
        proxyUrl = args[4]
    }

    proxyURL, err := url.Parse(proxyUrl)
    if err != nil {
        panic(err)
    }
    client.Transport = &http.Transport{
        Proxy: http.ProxyURL(proxyURL),
    }

    getVariable()

    pull(args[1], args[2])

    updateVariable()
}

func updateVariable() {
    putData := putPluginVariable{
        Name: pluginName,
        Data: map[string]any{
            pageName: page,
        },
    }

    jsonData, err := json.Marshal(putData)
    if err != nil {
        log.Fatal(err)
    }

    resp, err := http.Post(serverUrl+putVariablesUrlSuffix, contentType, bytes.NewBuffer(jsonData))

    res := response{}
    body, _ := io.ReadAll(resp.Body)
    _ = json.Unmarshal(body, &res)

    if err != nil && res.Code != 200 {
        log.Fatal(err)
    }
    defer resp.Body.Close()
}

func getVariable() {
    getData := getPluginVariable{
        Name: pluginName,
        Key:  []string{pageName},
    }
    jsonData, err := json.Marshal(getData)
    if err != nil {
        log.Fatal(err)
    }

    resp, err := http.Post(serverUrl+getVariablesUrlSuffix, contentType, bytes.NewBuffer(jsonData))
    res := response{}
    body, _ := io.ReadAll(resp.Body)
    _ = json.Unmarshal(body, &res)

    if err != nil && res.Code != 200 {
        log.Fatal(err)
    }
    defer resp.Body.Close()

    if res.Data != nil {
        page = int(res.Data.(map[string]any)[pageName].(float64))
    }
}

func pull(failedFilePath string, username string) {
    var notOverYet = true
    var failedStars []GithubStar

    var retry = true
    _, err := os.Stat(failedFilePath)
    if os.IsNotExist(err) {
        retry = false
    }

    var stillWrong []GithubStar
    if retry {
        file, _ := os.OpenFile(failedFilePath, os.O_CREATE|os.O_RDWR, 0666)
        defer file.Close()
        failedStarsJson, _ := io.ReadAll(file)
        _ = json.Unmarshal(failedStarsJson, &failedStars)

        for _, star := range failedStars {
            readmeContent := getReadmeContent(star, stillWrong)

            star.ReadmeContent = readmeContent

            pluginData := pluginData{
                Name:  pluginName,
                Index: indexName,
                Id:    star.FullName,
                Data:  star,
            }

            jsonData, err := json.Marshal(pluginData)
            if err != nil {
                log.Fatal(err)
            }

            resp, err := http.Post(serverUrl+saveDataUrlSuffix, contentType, bytes.NewBuffer(jsonData))

            res := response{}
            body, _ := io.ReadAll(resp.Body)
            _ = json.Unmarshal(body, &res)

            if err != nil && res.Code != 200 {
                stillWrong = append(stillWrong, star)
                continue
            }
        }
        // clear succeeded
        failedStars = stillWrong
    }

    if retry {
        // merge stillWrong and failedStars
        failedStars = append(stillWrong, failedStars...)
    }

    for notOverYet {
        stars := getStars(username, perPage, page)
        for _, star := range stars {
            readmeContent := getReadmeContent(star, failedStars)

            star.ReadmeContent = readmeContent

            pluginData := pluginData{
                Name:  pluginName,
                Index: indexName,
                Id:    star.FullName,
                Data:  star,
            }

            jsonData, err := json.Marshal(pluginData)
            if err != nil {
                log.Fatal(err)
            }

            resp, err := http.Post(serverUrl+saveDataUrlSuffix, contentType, bytes.NewBuffer(jsonData))

            res := response{}
            body, _ := io.ReadAll(resp.Body)
            _ = json.Unmarshal(body, &res)

            if err != nil && res.Code != 200 {
                stillWrong = append(stillWrong, star)
                continue
            }
        }

        if len(failedStars) > 0 {
            failedStarsJson, _ := json.Marshal(failedStars)
            file, _ := os.OpenFile(failedFilePath, os.O_CREATE|os.O_RDWR, 0666)

            // clear file
            _ = file.Truncate(0)
            _, _ = file.Seek(0, 0)

            _, _ = file.Write(failedStarsJson)
            _ = file.Close()
        }

        if perPage > len(stars) {
            notOverYet = false
        } else {
            page++
        }
    }
}

func getReadmeContent(star GithubStar, failedStars []GithubStar) string {
    readmeUrl := buildReadmeUrl(star)
    resp, err := client.Get(readmeUrl)
    if err != nil {
        fmt.Println(err)
        star.ReadmeUrl = readmeUrl
        failedStars = append(failedStars, star)
        return ""
    }
    readmeContent, _ := io.ReadAll(resp.Body)
    // Get plain text
    //md := markdown.Parse(readmeContent, nil)
    //return extractPlainText(md)
    return string(readmeContent)
}

func getStars(username string, perPage int, page int) []GithubStar {
    var url = "https://api.github.com/users/" +
        username + "/starred?per_page=" + strconv.Itoa(perPage) + "&page=" + strconv.Itoa(page)

    req, _ := http.NewRequest("GET", url, nil)
    req.Header.Add("Accept", "application/vnd.github+json")
    req.Header.Add("X-GitHub-Api-Version", "2022-11-28")

    resp, _ := client.Do(req)
    defer func(Body io.ReadCloser) {
        _ = Body.Close()
    }(resp.Body)
    body, _ := io.ReadAll(resp.Body)
    var stars []GithubStar
    _ = json.Unmarshal(body, &stars)
    return stars
}

//func extractPlainText(node ast.Node) string {
//    var plaintext strings.Builder
//
//    ast.WalkFunc(node, func(node ast.Node, entering bool) ast.WalkStatus {
//        if txt, ok := node.(*ast.Text); ok {
//            plaintext.Write(txt.Literal)
//        } else if _, ok := node.(*ast.Hardbreak); ok {
//            plaintext.WriteByte('\n')
//        }
//        return ast.GoToNext
//    })
//
//    return plaintext.String()
//}

func buildReadmeUrl(star GithubStar) string {
    return "https://raw.githubusercontent.com/" + star.FullName + "/" + star.DefaultBranch + "/README.md"
}

type GithubStar struct {
    Name          string `json:"name"`
    FullName      string `json:"full_name"`
    HtmlUrl       string `json:"html_url"`
    CloneUrl      string `json:"clone_url"`
    DefaultBranch string `json:"default_branch"`
    Language      string `json:"language"`
    ReadmeUrl     string `json:"readme_url"`
    ReadmeContent string `json:"readme_content"`
}

type putPluginVariable struct {
    Name string         `json:"name"`
    Data map[string]any `json:"data"`
}

type getPluginVariable struct {
    Name string   `json:"name"`
    Key  []string `json:"key"`
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
