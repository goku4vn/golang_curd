package internal

type PaginationParams struct {
	Page       int   `json:"page" form:"page"`
	PageSize   int   `json:"pageSize" form:"pageSize"`
	TotalCount int64 `json:"totalCount" form:"totalCount"`
}

func (p *PaginationParams) GetLimitOffset() (offset int, limit int) {
	if p.Page < 1 {
		p.Page = 1
	}
	if p.PageSize < 1 {
		p.PageSize = 2 // Default page size
	}

	offset = (p.Page - 1) * p.PageSize
	limit = p.PageSize

	return offset, limit
}
