package repo

import (
	"context"
	"fmt"
	m "ng_autocode_back_service/internal/core/fd/book/service/model/db"
	dto "ng_autocode_back_service/internal/core/fd/book/service/model/dto"
	logger "ng_autocode_back_service/pkg/logger"

	"go.uber.org/zap"
)

func (g *dbSt) GetFDOrderBook(c context.Context, dtoRequest *dto.FDOrderBookRequest) (dtoResponse []*dto.FDOrderBookResponse, e error) {
	logger.Log(c).Info("GetfdOrderBook called")

	// Create a slice to hold the results
	var orderBook []*dto.FDOrderBookResponse

	// Query the database for FDorder list
	var fdOrderBook []m.FDOrderBook

	// Query the database using raw SQL to avoid quoting issues
	whereQ := fmt.Sprintf(" where interestrate > %f", dtoRequest.InterestRate)
	if dtoRequest.MatchAccount != 0 {
		whereQ = whereQ + fmt.Sprintf(" and matchaccount = %d", dtoRequest.MatchAccount)
	}
	if dtoRequest.FdAgentCode != "" {
		whereQ = whereQ + fmt.Sprintf(" and fdagentcode = '%s'", dtoRequest.FdAgentCode)
	}
	query := g.oracle.WithContext(c).Raw(`
		SELECT transId, userId, matchAccount, fdSchemeName, 
		       interestRate, amount, maturityDays, autoRenew,
		       fdagentCode, otherFlags, dtUpdate
		FROM fd_trans` + whereQ)

	logger.Log(c).Debug("Executing query", zap.Any("query", query))

	if err := query.Scan(&fdOrderBook).Error; err != nil {
		logger.Log(c).Error("Failed to execute query", zap.Error(err))
		return nil, err
	}

	// Map the results to the response dto
	for _, order := range fdOrderBook {
		orderBook = append(orderBook, &dto.FDOrderBookResponse{
			TransId:      order.TransId,
			UserId:       order.UserId,
			MatchAccount: order.MatchAccount,
			FdSchemeName: order.FdSchemeName,
			InterestRate: order.InterestRate,
			Amount:       order.Amount,
			MaturityDays: order.MaturityDays,
			AutoRenew:    order.AutoRenew,
			FdAgentCode:  order.FdAgentCode,
			OtherFlags:   order.OtherFlags,
			DtUpdate:     order.DtUpdate,
		})
	}

	logger.Log(c).Debug("Query result", zap.Any("orders", orderBook))
	return orderBook, nil
}
