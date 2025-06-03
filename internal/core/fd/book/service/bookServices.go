package service

import (
	"context"
	dto "ng_autocode_back_service/internal/core/fd/book/service/model/dto"
)

func (g *serviceItem) GetFDOrderBook(c context.Context, request *dto.FDOrderBookRequest) (dtoResponse []*dto.FDOrderBookResponse, e error) {
	// Call the data layer to get the FDAgent list
	dtoResponse, err := g.dataL.GetFDOrderBook(c)
	if err != nil {
		return nil, err
	}

	return dtoResponse, nil
}
