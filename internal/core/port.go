package module

import (
	"net/http"
	e "ng_autocode_back_service/common/http"
	cmnUtil "ng_autocode_back_service/common/util"
	fd "ng_autocode_back_service/internal/core/fd"
	repo "ng_autocode_back_service/internal/infra/repository"
	"time"

	"github.com/gin-gonic/gin"
)

type module struct {
	fd.ModuleServiceLayer
	//pms.ModuleServiceLayer
	//common.ModuleServiceLayer
}

type ModuleLayer interface {
	GetSystemMatrics(c *gin.Context)
	GetAppVersion(c *gin.Context)
	GetReadTextData(c *gin.Context)
	GetReadJsonData(c *gin.Context)

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
	c.JSON(http.StatusOK, e.SuccesssResponse("I AM HEALTHY 2025-06-03 16:00"))
}

func (m *module) GetAppVersion(c *gin.Context) {
	sDate := time.Now().Format("2006-01-02 15:04:05")
	c.JSON(http.StatusOK, e.SuccesssResponse("Version 7.3.1.4 - "+sDate))
}

func (m *module) GetReadTextData(c *gin.Context) {
	sDate := cmnUtil.ReadTextFile("C:\\data\\info.txt")
	c.JSON(http.StatusOK, e.SuccesssResponse(sDate))
}

func (m *module) GetReadJsonData(c *gin.Context) {
	sDate := cmnUtil.ReadTextFile("C:\\data\\sampleResponse.json")
	c.JSON(http.StatusOK, e.SuccesssResponse(sDate))
}
