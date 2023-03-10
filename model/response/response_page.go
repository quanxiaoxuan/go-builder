package response

import "github.com/quanxiaoxuan/go-builder/model/request"

// 分页结果
type PageResponse struct {
	CurrPage   int64       `json:"currPage" comment:"当前页码"`
	PageSize   int64       `json:"pageSize" comment:"当前页数量"`
	PageTotal  int64       `json:"pageTotal" comment:"总页数"`
	Total      int64       `json:"total" comment:"结果总数量"`
	ResultList interface{} `json:"resultList" comment:"返回结果集"`
}

// 封装分页结果
func BuildPageData(param *request.Page, result interface{}, total int64) *PageResponse {
	var pageData = PageResponse{
		CurrPage:   param.CurrPage,
		PageSize:   param.PageSize,
		PageTotal:  param.GetPageTotal(total),
		Total:      total,
		ResultList: result,
	}
	return &pageData
}
