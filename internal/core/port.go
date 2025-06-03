package module

import (
	"net/http"
	cmn "ng_autocode_back_service/common/http"
	fd "ng_autocode_back_service/internal/core/fd"
	repo "ng_autocode_back_service/internal/infra/repository"

	"github.com/gin-gonic/gin"
)

type module struct {
	fd.ModuleServiceLayer
	//pms.ModuleServiceLayer
	//common.ModuleServiceLayer
}

type ModuleLayer interface {
	GetSystemMatrics(c *gin.Context)

	fd.ModuleServiceLayer
	//pms.ModuleServiceLayer
	//common.ModuleServiceLayer
}

func NewAppModuleLayer(repo repo.DataObject) ModuleLayer {
	return &module{
		fd.NewModuleServiceLayer(repo),
		//pms.NewModuleServiceLayer(repo),
		//common.NewModuleServiceLayer(repo),
	}
}

func (m *module) GetSystemMatrics(c *gin.Context) {
	c.JSON(http.StatusOK, cmn.SuccesssResponse("I AM HEALTHY 2025-06-03 12:00"))
}
