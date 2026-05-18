结合源码，可以把 Bleve 的倒排查询过程总结成这 5 步：

### 第一层：API 层

你写 NewMatchQuery + NewSearchRequest + index.Search

### 第二层：请求编译层

SearchInContext() 打开 reader，准备 scoring/context

### 第三层：查询执行层

req.Query.Searcher(...) 把 Query 编译成 Searcher 树
这里真正把“match/phrase/bool”翻译成底层执行计划

### 第四层：倒排遍历层

底层 TermSearcher / PhraseSearcher / BooleanSearcher 等，从 indexReader 读取词典和 posting list，产生 DocumentMatch

### 第五层：结果整理层

collector.Collect(...) 做 TopN、排序、分页，再补字段、高亮、facet，最终返回结果