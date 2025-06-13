package module

import (
	repo "ng_autocode_back_service/common/infra/repository"
	fd "ng_autocode_back_service/internal/fd"
)

// #region L2 - Module ServiceGroup Layer

type fdServiceGroup struct {
	fd.ListServices
	fd.BookServices
	//trans.TransactionServices
	//misc.MiscServices
}

type FDModuleServiceLayer interface {
	fd.ListServices
	fd.BookServices
	//trans.TransactionServices
	//misc.MiscServices
}

func NewFDModuleServiceLayer(repo repo.DataObject) FDModuleServiceLayer {
	return &fdServiceGroup{
		fd.NewListServiceGroup(repo),
		fd.NewBookServiceGroup(repo),
		//trans.NewTransactionServices(repo),
		//misc.NewMiscServices(repo),
	}
}

//#endregion
