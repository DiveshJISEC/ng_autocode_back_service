package fdBook

import (
	"context"
	repo "ng_autocode_back_service/common/infra/repository"

	"gorm.io/gorm"
)

// #region L4 - Application Logic Layer

type serviceItem struct {
	dataL DataLayer
}

type ServiceLayer interface {
	GetFDOrderBook(c context.Context, request *FDOrderBookRequest) (response []*FDOrderBookResponse, e error)
}

func NewListServicesLayerObject(repo repo.DataObject) ServiceLayer {
	return &serviceItem{
		dataL: NewDataLayerOb(repo),
	}
}

//#endregion

// #region L5 - Repository controller Layer

type dbSt struct {
	oracle *gorm.DB
}

type DataLayer interface {
	GetFDOrderBook(c context.Context, dtoRequest *FDOrderBookRequest) (dtoResponse []*FDOrderBookResponse, e error)
}

func NewDataLayerOb(repo repo.DataObject) DataLayer {
	return &dbSt{oracle: repo.DBSet.ReadDB.Db}
}

//#endregion
