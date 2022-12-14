package utils

import (
	"github.com/astaxie/beego/orm"
	"strings"
)

// Filters 查询条件结构
type Filters struct {
	Equals   map[string]string
	Likes    map[string]string
	Betweens map[string][]string
	Ins      map[string][]string
	Orders   map[string]string
}

// Paging 分页数据结构
type Page struct {
	Page      int64        `json:"page"`
	PerPage   int64        `json:"per_page"`
	TotalPage int64        `json:"total_page"`
	PrevPage  int64        `json:"prev_page"`
	NextPage  int64        `json:"next_page"`
	Total     int64        `json:"total"`
	List      []orm.Params `json:"list"`
}

// @query orm.QuerySeter		 	// 查询初始化
// @fields []string	 			// 查询字段
// @fieldMap map[string]string	// 允许过滤字段与别名映射
// @filters *Filters	 			// 过滤条件
// @page int64			 			// 当前页码
// @perPage int64		 			// 每页数量
// TODO： per_page 不能超过 200， page 不能超过总页数等等处理
// Paging 分页处理工具
func Paging(
	query orm.QuerySeter,
	fields []string,
	filtersMap map[string]string,
	filters *Filters,
	page int64,
	perPage int64,
) *Page {
	// fieldMap 查询字段映射： 1.没有在这里的字段不允许查询; 2. 表别名时字段别名映射
	// TODO: 别名无效
	// query: =
	for equalKey, equalValue := range filters.Equals {
		fieldAlias, ok := filtersMap[equalKey]

		if ok {
			query = query.Filter(fieldAlias, equalValue)
		}
	}

	// query: like
	for likeKey, likeValue := range filters.Likes {
		fieldAlias, ok := filtersMap[likeKey]

		if ok {
			query = query.Filter(fieldAlias+"__icontains", likeValue)
		}
	}

	// query: between
	for betweenKey, betweenValue := range filters.Betweens {
		fieldAlias, ok := filtersMap[betweenKey]
		betweenValueLen := len(betweenValue)

		// between 参数必须成对
		if ok && betweenValueLen == 2 {
			query = query.Filter(fieldAlias+"__gte", betweenValue[0])
			query = query.Filter(fieldAlias+"__lte", betweenValue[1])
		}
	}

	// query: in
	for inKey, inValue := range filters.Ins {
		fieldAlias, ok := filtersMap[inKey]

		if ok {
			query = query.Filter(fieldAlias+"__in", inValue)
		}
	}

	// 数量统计 total
	total, _ := query.Count()

	// query: order by
	for orderKey, orderValue := range filters.Orders {
		fieldAlias, ok := filtersMap[orderKey]

		if ok {
			if strings.ToLower(orderValue) == "desc" {
				query = query.OrderBy("-" + fieldAlias)
			} else {
				query = query.OrderBy(fieldAlias)
			}
		}
	}

	// 查询
	var list []orm.Params
	query.
		Limit(perPage, getOffset(page, perPage)).
		Values(&list, fields...)

	// 查询结果转 key 转 小写下划线
	var pageList []orm.Params
	for _, listValue := range list {
		m := map[string]interface{}{}

		for k, v := range listValue {
			m[snakeString(k)] = v
		}

		pageList = append(pageList, m)
	}

	// 分页计算
	totalPage := getTotalPage(total, perPage)
	prevPage := getPrevPage(page)
	nextPage := getNextPage(page, totalPage)

	// 分页数据
	Paging := Page{
		Page:      page,
		PerPage:   perPage,
		TotalPage: totalPage,
		PrevPage:  prevPage,
		NextPage:  nextPage,
		Total:     total,
		List:      pageList,
	}

	return &Paging
}

// getOffset 分页偏移量
func getOffset(page, perPage int64) int64 {
	return (page - 1) * perPage
}

// getTotalPage 计算总页数
func getTotalPage(total, perPage int64) int64 {
	totalPage := total / perPage

	if (total % perPage) > 0 {
		totalPage = (total / perPage) + 1
	}

	return totalPage
}

// getPrevPage 上一页页码
func getPrevPage(page int64) int64 {
	prevPage := int64(0)

	if page > 1 {
		prevPage = page - 1
	}

	return prevPage
}

// getNextPage 下一页页码
func getNextPage(page, totalPage int64) int64 {
	nextPage := int64(0)

	if totalPage > page {
		nextPage = page + 1
	}

	return nextPage
}
