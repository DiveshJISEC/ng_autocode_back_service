package service

import (
	serv "ng_autocode_back_service/internal/core/fd/book/service"
	repo "ng_autocode_back_service/internal/infra/repository"

	"github.com/gin-gonic/gin"
)

type bookServices struct {
	serviceLayer serv.ServiceLayer
}

type BookServices interface {
	GetFDOrderBook(c *gin.Context)
}

func NewBookServiceGroup(repo repo.DataObject) BookServices {
	return &bookServices{
		serviceLayer: serv.NewListServicesLayerObject(repo),
	}
}
