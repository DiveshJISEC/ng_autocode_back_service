package service

import (
	"context"
	dto "ng_autocode_back_service/internal/core/fd/list/service/model/dto"
	data "ng_autocode_back_service/internal/core/fd/list/service/repo"
	repo "ng_autocode_back_service/internal/infra/repository"
)

type serviceItem struct {
	dataL data.DataLayer
}

type ServiceLayer interface {
	GetFDAgentList(c context.Context, dtoRequest *dto.FDAgentListRequest) (response []*dto.FDAgentListResponse, e error)
}

func NewListServicesLayerObject(repo repo.DataObject) ServiceLayer {
	return &serviceItem{
		dataL: data.NewDataLayerOb(repo),
	}
}
