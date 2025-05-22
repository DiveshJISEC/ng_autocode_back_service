package repo

import (
	"context"
	dto "ng_autocode_back_service/internal/core/fd/list/service/model/dto"
	repo "ng_autocode_back_service/internal/infra/repository"

	"gorm.io/gorm"
)

type dbSt struct {
	oracle *gorm.DB
}

type DataLayer interface {
	GetFDAgentList(context.Context) (dtoResponse []*dto.FDAgentListResponse, e error)
}

func NewDataLayerOb(repo repo.DataObject) DataLayer {
	return &dbSt{oracle: repo.DBSet.ReadDB.Db}
}
