package fdBook

import (
	"context"
)

func (g *serviceItem) GetFDOrderBook(c context.Context, dtoRequest *FDOrderBookRequest) (dtoResponse []*FDOrderBookResponse, e error) {
	// Call the data layer to get the FDAgent list
	dtoResponse, err := g.dataL.GetFDOrderBook(c, dtoRequest)
	if err != nil {
		return nil, err
	}

	return dtoResponse, nil
}
