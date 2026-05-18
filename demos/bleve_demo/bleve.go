package bleve_demo

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/mapping"
	"github.com/yanyiwu/gojieba"

	// 只为注册 gojieba tokenizer / analyzer
	_ "github.com/ttys3/gojieba-bleve/v2"
)

type Article struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Tag     string `json:"tag"`
}

func Demo() {
	indexPath := "data/bleve/articles.bleve"
	_ = os.RemoveAll(indexPath)
	defer os.RemoveAll(indexPath)

	indexMapping := buildIndexMapping()

	index, err := bleve.New(indexPath, indexMapping)
	if err != nil {
		log.Fatalf("create index failed: %v", err)
	}
	defer func() {
		// gojieba-bleve 的示例里明确演示了需要在结束时 Free，
		// 用来清理 cgo 分配的内存。
		if jiebaAnalyzer, ok := index.Mapping().AnalyzerNamed("gojieba").(interface{ Free() }); ok {
			jiebaAnalyzer.Free()
		}
		_ = index.Close()
	}()

	docs := []Article{
		{
			ID:      "1",
			Title:   "Go 语言后端开发",
			Content: "程序员喜欢使用 Go 语言开发高性能后端服务",
			Tag:     "golang",
		},
		{
			ID:      "2",
			Title:   "全文检索入门",
			Content: "Bleve 是 Go 生态里常见的全文检索库，适合做站内搜索",
			Tag:     "search",
		},
		{
			ID:      "3",
			Title:   "程序员成长路线",
			Content: "程序员需要理解数据结构、倒排索引、缓存和分布式系统",
			Tag:     "career",
		},
		{
			ID:      "4",
			Title:   "旅游随笔",
			Content: "今天去了长江大桥，风景很好",
			Tag:     "life",
		},
	}

	batch := index.NewBatch()
	for _, doc := range docs {
		if err := batch.Index(doc.ID, doc); err != nil {
			log.Fatalf("batch index failed: %v", err)
		}
	}
	if err := index.Batch(batch); err != nil {
		log.Fatalf("execute batch failed: %v", err)
	}

	fmt.Println("索引构建完成")

	runQuery(index, "程序员")
	runQuery(index, "Go 语言")
	runQuery(index, "长江大桥")
	runQuery(index, "全文检索")
}

func buildIndexMapping() *mapping.IndexMappingImpl {
	indexMapping := bleve.NewIndexMapping()

	// 注册自定义 tokenizer
	err := indexMapping.AddCustomTokenizer("gojieba",
		map[string]interface{}{
			"dictpath":     gojieba.DICT_PATH,
			"hmmpath":      gojieba.HMM_PATH,
			"userdictpath": gojieba.USER_DICT_PATH,
			"idf":          gojieba.IDF_PATH,
			"stop_words":   gojieba.STOP_WORDS_PATH,
			"type":         "gojieba",
		},
	)
	if err != nil {
		log.Fatalf("add custom tokenizer failed: %v", err)
	}

	// 注册自定义 analyzer
	err = indexMapping.AddCustomAnalyzer("gojieba",
		map[string]interface{}{
			"type":      "gojieba",
			"tokenizer": "gojieba",
		},
	)
	if err != nil {
		log.Fatalf("add custom analyzer failed: %v", err)
	}

	// 默认分析器改为 gojieba
	indexMapping.DefaultAnalyzer = "gojieba"

	// 文档映射
	docMapping := bleve.NewDocumentMapping()

	textField := bleve.NewTextFieldMapping()
	textField.Store = true
	textField.IncludeTermVectors = true
	textField.Analyzer = "gojieba"

	tagField := bleve.NewTextFieldMapping()
	tagField.Store = true
	tagField.Analyzer = "keyword"

	docMapping.AddFieldMappingsAt("title", textField)
	docMapping.AddFieldMappingsAt("content", textField)
	docMapping.AddFieldMappingsAt("tag", tagField)

	indexMapping.DefaultMapping = docMapping
	indexMapping.TypeField = "type"

	return indexMapping
}

func runQuery(index bleve.Index, q string) {
	fmt.Printf("\n===== 查询: %s =====\n", q)

	// 用 QueryStringQuery 更接近实际搜索框体验
	query := bleve.NewQueryStringQuery(q)
	req := bleve.NewSearchRequest(query)
	req.Fields = []string{"title", "content", "tag"}
	req.Highlight = bleve.NewHighlight()
	req.Size = 10

	res, err := index.Search(req)
	if err != nil {
		log.Fatalf("search failed: %v", err)
	}

	fmt.Printf("命中总数: %d\n", res.Total)
	for i, hit := range res.Hits {
		fmt.Printf("\n--- 第 %d 条 ---\n", i+1)
		fmt.Printf("ID: %s\n", hit.ID)
		fmt.Printf("Score: %.4f\n", hit.Score)

		if len(hit.Fragments) > 0 {
			b, _ := json.MarshalIndent(hit.Fragments, "", "  ")
			fmt.Printf("Fragments: %s\n", string(b))
		}

		b, _ := json.MarshalIndent(hit.Fields, "", "  ")
		fmt.Printf("Fields: %s\n", string(b))
	}
}
