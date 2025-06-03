package service

import (
	"context"
	dto "ng_autocode_back_service/internal/core/fd/book/service/model/dto"
	data "ng_autocode_back_service/internal/core/fd/book/service/repo"
	repo "ng_autocode_back_service/internal/infra/repository"
)

type serviceItem struct {
	dataL data.DataLayer
}

type ServiceLayer interface {
	GetFDOrderBook(c context.Context, request *dto.FDOrderBookRequest) (response []*dto.FDOrderBookResponse, e error)
}

func NewListServicesLayerObject(repo repo.DataObject) ServiceLayer {
	return &serviceItem{
		dataL: data.NewDataLayerOb(repo),
	}
}
