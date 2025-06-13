package service

import (
	serv "ng_autocode_back_service/internal/core/fd/list/service"
	repo "ng_autocode_back_service/internal/infra/repository"

	"github.com/gin-gonic/gin"
)

type listServices struct {
	serviceLayer serv.ServiceLayer
}

type ListServices interface {
	GetFDAgentList(c *gin.Context)
}

func NewListServiceGroup(repo repo.DataObject) ListServices {
	return &listServices{
		serviceLayer: serv.NewListServicesLayerObject(repo),
	}
}
