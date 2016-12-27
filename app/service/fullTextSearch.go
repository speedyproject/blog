package service

import (
	"blog/app/models"
	"log"

	"github.com/huichen/wukong/engine"
	"github.com/huichen/wukong/types"
	"github.com/revel/revel"
)

var (
	// searcher是协程安全的
	searcher = engine.Engine{}
)

func InitSearcher() {
	// 初始化
	searcher.Init(types.EngineInitOptions{
		SegmenterDictionaries: revel.AppPath + "/service/searchData/dictionary.txt",
		StopTokenFile:         revel.AppPath + "/service/searchData/stop_tokens.txt",
	})
	defer searcher.Close()

	// 将文档加入索引，docId 从1开始
	blogModel := new(models.Blogger)
	blogs, _ := blogModel.FindList()
	for _, v := range blogs {
		searcher.IndexDocument(uint64(v.Id), types.DocumentIndexData{Content: v.ContentHTML}, false)
	}

	// 等待索引刷新完毕
	searcher.FlushIndex()
}

func FullTextSearch(keywords string) []int64 {
	// 搜索输出格式见types.SearchResponse结构体
	res := searcher.Search(types.SearchRequest{Text: keywords})
	log.Println("search result: ", res)
	var blogIDs []int64
	for _, v := range res.Docs {
		blogIDs = append(blogIDs, int64(v.DocId))
	}
	return blogIDs
}
