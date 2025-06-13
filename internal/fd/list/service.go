package fdList

import (
	"context"
)

func (g *serviceItem) GetFDAgentList(c context.Context, dtoRequest *FDAgentListRequest) (dtoResponse []*FDAgentListResponse, e error) {
	// Call the data layer to get the FDAgent list
	dtoResponse, err := g.dataL.GetFDAgentList(c, dtoRequest)
	if err != nil {
		return nil, err
	}

	return dtoResponse, nil
}
