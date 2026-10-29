package pagination

import (
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Params struct {
	Page    int    `json:"page"`
	PerPage int    `json:"per_page"`
	Search  string `json:"search"`
	SortBy  string `json:"sort_by"`
	Order   string `json:"order"`
}

type Meta struct {
	CurrentPage int   `json:"current_page"`
	PerPage     int   `json:"per_page"`
	TotalItems  int64 `json:"total_items"`
	TotalPages  int   `json:"total_pages"`
}

func GetPaginationParams(c *gin.Context) Params {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	if page < 1 {
		page = 1
	}

	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "10"))
	if perPage < 1 || perPage > 100 {
		perPage = 10
	}

	search := c.Query("search")
	sortBy := c.DefaultQuery("sort_by", "created_at")
	order := c.DefaultQuery("order", "desc")
	if order != "asc" && order != "desc" {
		order = "desc"
	}

	return Params{
		Page:    page,
		PerPage: perPage,
		Search:  search,
		SortBy:  sortBy,
		Order:   order,
	}
}

func (p *Params) Offset() int {
	return (p.Page - 1) * p.PerPage
}

func BuildMeta(params Params, totalItems int64) *Meta {
	totalPages := int(math.Ceil(float64(totalItems) / float64(params.PerPage)))
	if totalPages == 0 {
		totalPages = 1
	}

	return &Meta{
		CurrentPage: params.Page,
		PerPage:     params.PerPage,
		TotalItems:  totalItems,
		TotalPages:  totalPages,
	}
}
