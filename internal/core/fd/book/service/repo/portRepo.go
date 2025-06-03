package repo

import (
	"context"
	dto "ng_autocode_back_service/internal/core/fd/book/service/model/dto"
	repo "ng_autocode_back_service/internal/infra/repository"

	"gorm.io/gorm"
)

type dbSt struct {
	oracle *gorm.DB
}

type DataLayer interface {
	GetFDOrderBook(c context.Context, dtoRequest *dto.FDOrderBookRequest) (dtoResponse []*dto.FDOrderBookResponse, e error)
}

func NewDataLayerOb(repo repo.DataObject) DataLayer {
	return &dbSt{oracle: repo.DBSet.ReadDB.Db}
}
