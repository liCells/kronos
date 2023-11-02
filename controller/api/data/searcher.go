package data

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/liCells/controller/api/common/response"
	"github.com/liCells/controller/global"
	"github.com/olivere/elastic/v7"
	"strings"
)

func (r *DataApi) BasicSearch(c *gin.Context) {
	searcher := Searcher{
		From: 0,
		Size: 30,
	}
	if buildSearcher(c, &searcher) {
		return
	}

	esSearcher := global.ESClient.Search()
	if searcher.Indexes != nil {
		// set indexes
		esSearcher = esSearcher.
			Index(searcher.Indexes...)
	}
	if global.Config.Es.Analyzer != "" {
		esSearcher = esSearcher.
			Query(elastic.NewQueryStringQuery(searcher.Content).Analyzer(global.Config.Es.Analyzer))
	} else {
		esSearcher = esSearcher.
			Query(elastic.NewQueryStringQuery(searcher.Content))
	}

	res, err := esSearcher.
		From(searcher.From).
		Size(searcher.Size).
		Do(context.Background())

	esSearcher.Explain(true)
	if err != nil {
		response.FailWithMessage(c, "Search failure")
		return
	}

	response.SuccessWithData(c, Records{
		Data:  res.Hits.Hits,
		Total: res.TotalHits(),
	})
}

func buildSearcher(c *gin.Context, searcher *Searcher) bool {
	err := c.ShouldBindJSON(&searcher)
	if err != nil {
		response.BadRequest(c)
		return true
	}
	if strings.TrimSpace(searcher.Content) == "" {
		response.FailWithMessage(c, "Content is empty")
		return true
	}
	return false
}

type Records struct {
	Data  any   `json:"data"`
	Total int64 `json:"total"`
}

type Searcher struct {
	Content string   `json:"content"`
	Indexes []string `json:"indexes"`
	From    int      `json:"from"`
	Size    int      `json:"size"`
}
