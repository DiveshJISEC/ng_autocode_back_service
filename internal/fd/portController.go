package fd

import (
	repo "ng_autocode_back_service/common/infra/repository"
	book "ng_autocode_back_service/internal/fd/book"
	list "ng_autocode_back_service/internal/fd/list"

	"github.com/gin-gonic/gin"
)

// #region L3 - Service controller Layer

// #region L3. Book Service controller Layer

type bookServices struct {
	serviceLayer book.ServiceLayer
}

type BookServices interface {
	GetFDOrderBook(c *gin.Context)
}

func NewBookServiceGroup(repo repo.DataObject) BookServices {
	return &bookServices{
		serviceLayer: book.NewListServicesLayerObject(repo),
	}
}

//#endregion

// #region L3. List Service controller Layer

type listServices struct {
	serviceLayer list.ServiceLayer
}

type ListServices interface {
	GetFDAgentList(c *gin.Context)
}

func NewListServiceGroup(repo repo.DataObject) ListServices {
	return &listServices{
		serviceLayer: list.NewListServicesLayerObject(repo),
	}
}

//#endregion

//#endregion
