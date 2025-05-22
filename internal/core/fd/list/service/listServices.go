package service

import (
	"context"
	dto "ng_autocode_back_service/internal/core/fd/list/service/model/dto"
)

func (g *serviceItem) GetFDAgentList(c context.Context, request *dto.FDAgentListRequest) (dtoResponse []*dto.FDAgentListResponse, e error) {
	// Call the data layer to get the FDAgent list
	dtoResponse, err := g.dataL.GetFDAgentList(c)
	if err != nil {
		return nil, err
	}

	return dtoResponse, nil
}
