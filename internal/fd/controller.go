package fd

import (
	"net/http"
	e "ng_autocode_back_service/common/model"
	logger "ng_autocode_back_service/common/pkg/logger"

	book "ng_autocode_back_service/internal/fd/book"
	list "ng_autocode_back_service/internal/fd/list"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (f *bookServices) GetFDOrderBook(c *gin.Context) {
	logger.Log(c).Info("GetFDOrderBook called")

	// Create a new request object
	dtoRequest := book.FDOrderBookRequest{}

	// Bind the request data to the struct
	if err := c.BindJSON(&dtoRequest); err != nil {
		logger.Log(c).Error("Invalid request payload", zap.Error(err))
		c.JSON(http.StatusBadRequest, e.FailureResponse(e.ApiErrors.BadRequestError.WithErrorDescription(err.Error())))
		c.Abort()
		return
	}

	// Call the service layer to get the OrderBook list
	dtoResponse, err := f.serviceLayer.GetFDOrderBook(c, &dtoRequest)

	if err != nil {
		logger.Log(c).Error("Failed to get OrderBook list", zap.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, e.FailureResponse(e.ApiErrors.BadRequestError.WithErrorDescription(err.Error())))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, e.SuccesssResponse(dtoResponse))
}

func (f *listServices) GetFDAgentList(c *gin.Context) {
	logger.Log(c).Info("GetFDAgentList called")

	// Create a new request object
	dtoRequest := list.FDAgentListRequest{}

	// Bind the request data to the struct
	if err := c.BindJSON(&dtoRequest); err != nil {
		logger.Log(c).Error("Invalid request payload", zap.Error(err))
		c.JSON(http.StatusBadRequest, e.FailureResponse(e.ApiErrors.BadRequestError.WithErrorDescription(err.Error())))
		c.Abort()
		return
	}

	// Call the service layer to get the FDAgent list
	dtoResponse, err := f.serviceLayer.GetFDAgentList(c, &dtoRequest)

	if err != nil {
		logger.Log(c).Error("Failed to get FDAgent list", zap.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, e.FailureResponse(e.ApiErrors.BadRequestError.WithErrorDescription(err.Error())))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, e.SuccesssResponse(dtoResponse))
}
