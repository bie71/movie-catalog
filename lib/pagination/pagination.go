package pagination

import (
	"math"
	"movie-catalog/internal/core/domain/entity"
)

type PaginationInter interface {
	AddPagination(totalData, page, perPage int) (*entity.Page, error)
}

type Options struct {
}

func (o *Options) AddPagination(totalData, page, perPage int) (*entity.Page, error) {
	newPage := page
	if newPage <= 0 {
		return nil, ErrorPage
	}
	limitData := 10
	if perPage > 0 {
		limitData = perPage
	}
	totalPage := int(math.Ceil(float64(totalData) / float64(limitData)))
	last := (newPage * limitData)
	firts := last - limitData

	if totalData < last {
		last = totalData
	}
	zeroPage := &entity.Page{
		Page:      newPage,
		PageCount: 1,
	}

	if totalData == 0 && newPage == 1 {
		return zeroPage, nil
	}
	if newPage > totalPage {
		return nil, ErrorMaxPage
	}

	pages := &entity.Page{
		Page:       newPage,
		PerPage:    limitData,
		PageCount:  totalPage,
		TotalCount: totalData,
		Firts:      firts,
		Last:       last,
	}

	return pages, nil
}

func NewPagination() PaginationInter {
	pagination := new(Options)
	return pagination
}
