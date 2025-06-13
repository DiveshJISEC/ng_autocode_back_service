package service

import (
	"net/http"
	e "ng_autocode_back_service/common/http"
	dto "ng_autocode_back_service/internal/core/fd/book/service/model/dto"
	logger "ng_autocode_back_service/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (f *bookServices) GetFDOrderBook(c *gin.Context) {
	logger.Log(c).Info("GetFDOrderBook called")

	// Create a new request object
	dtoRequest := dto.FDOrderBookRequest{}

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
