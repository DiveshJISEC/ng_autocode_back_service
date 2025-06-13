package moduleService

import (
	book "ng_autocode_back_service/internal/core/fd/book"
	list "ng_autocode_back_service/internal/core/fd/list"
	repo "ng_autocode_back_service/internal/infra/repository"
)

type serviceGroup struct {
	list.ListServices
	book.BookServices
	//trans.TransactionServices
	//misc.MiscServices
}

type ModuleServiceLayer interface {
	list.ListServices
	book.BookServices
	//trans.TransactionServices
	//misc.MiscServices
}

func NewModuleServiceLayer(repo repo.DataObject) ModuleServiceLayer {
	return &serviceGroup{
		list.NewListServiceGroup(repo),
		book.NewBookServiceGroup(repo),
		//trans.NewTransactionServices(repo),
		//misc.NewMiscServices(repo),
	}
}
