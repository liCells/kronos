package data

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/liCells/controller/api/common/response"
	"github.com/liCells/controller/global"
	"github.com/olivere/elastic/v7"
	"strconv"
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

	if searcher.Indexes == nil {
		searcher.Indexes = global.Indexes
	}

	boolQuery := elastic.NewBoolQuery()

	if global.SearchSchemes == nil {
		for _, index := range searcher.Indexes {
			schemes, isExist := global.SearchSchemes[index]
			if isExist {
				for _, scheme := range schemes {
					name := index + "." + scheme.Field + "^" + strconv.Itoa(scheme.Boost)
					matchQuery := elastic.NewMatchQuery(name, searcher.Content)

					if global.Config.Es.Analyzer != "" {
						matchQuery.Analyzer(global.Config.Es.Analyzer)
					}

					boolQuery.Should(matchQuery)
				}
			}
		}
	} else {
		boolQuery.Should(elastic.NewQueryStringQuery(searcher.Content))
	}

	res, err := global.ESClient.Search().
		Index(searcher.Indexes...).
		From(searcher.From).
		Size(searcher.Size).
		Query(boolQuery).
		Highlight(elastic.NewHighlight().PreTags("<em>").PostTags("</em>")).
		Do(context.Background())

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
