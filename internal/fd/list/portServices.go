package fdList

import (
	"context"
	repo "ng_autocode_back_service/common/infra/repository"

	"gorm.io/gorm"
)

// #region L4 - Application logic Layer

type serviceItem struct {
	dataL DataLayer
}

type ServiceLayer interface {
	GetFDAgentList(c context.Context, dtoRequest *FDAgentListRequest) (response []*FDAgentListResponse, e error)
}

func NewListServicesLayerObject(repo repo.DataObject) ServiceLayer {
	return &serviceItem{
		dataL: NewDataLayerOb(repo),
	}
}

//#endregion

// #region L5 - Repository Layer

type dbSt struct {
	oracle *gorm.DB
}

type DataLayer interface {
	GetFDAgentList(c context.Context, dtoRequest *FDAgentListRequest) (dtoResponse []*FDAgentListResponse, e error)
}

func NewDataLayerOb(repo repo.DataObject) DataLayer {
	return &dbSt{oracle: repo.DBSet.ReadDB.Db}
}

//#endregion
