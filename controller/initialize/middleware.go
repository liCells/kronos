package initialize

import (
    "github.com/liCells/controller/global"
    "github.com/olivere/elastic/v7"
    "strconv"
)

func InitEs() {
    es := global.Config.Es
    address := es.Host + ":" + strconv.Itoa(int(es.Port))
    esClient, err := elastic.NewClient(
        elastic.SetURL(address),
        elastic.SetSniff(false),
    )
    if err != nil {
        panic("ES client init error")
    }
    global.ESClient = esClient
}
