package support

import (
	"github.com/huichen/wukong/engine"
	"github.com/huichen/wukong/types"
	"github.com/revel/revel"
)

var (
	// searcher是协程安全的
	Searcher = engine.Engine{}
)

func InitSearcher() {
	// 初始化
	Searcher.Init(types.EngineInitOptions{
		SegmenterDictionaries: revel.AppPath + "/service/searchData/dictionary.txt",
		StopTokenFile:         revel.AppPath + "/service/searchData/stop_tokens.txt",
	})
	defer Searcher.Close()
}

func FullTextSearch(keywords string) []int64 {
	// 搜索输出格式见types.SearchResponse结构体
	res := Searcher.Search(types.SearchRequest{Text: keywords})
	revel.TRACE.Println("search result: ", res)
	var blogIDs []int64
	for _, v := range res.Docs {
		blogIDs = append(blogIDs, int64(v.DocId))
	}
	return blogIDs
}
