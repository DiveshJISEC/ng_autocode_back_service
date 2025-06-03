package repo

import (
	"context"
	m "ng_autocode_back_service/internal/core/fd/book/service/model/db"
	dto "ng_autocode_back_service/internal/core/fd/book/service/model/dto"
	logger "ng_autocode_back_service/pkg/logger"

	"go.uber.org/zap"
)

func (g *dbSt) GetFDOrderBook(c context.Context) (dtoResponse []*dto.FDOrderBookResponse, e error) {
	logger.Log(c).Info("GetfdOrderBook called")

	// Create a slice to hold the results
	var orderBook []*dto.FDOrderBookResponse

	// Query the database for FDorder list
	var fdOrderBook []m.FDOrderBook
	query := g.oracle.WithContext(c).Table("fd_trans").
		Select(
			`transId`,
			`userId`,
			`matchAccount`,
			`fdSchemeName`,
			`interestRate`,
			`amount`,
			`maturityDays`,
			`autoRenew`,
			`fdorderCode`,
			`otherFlags`,
			`dtUpdate`,
		)

	logger.Log(c).Debug("Executing query", zap.Any("query", query))
	err := query.Scan(&fdOrderBook).Error
	if err != nil {
		logger.Log(c).Error("Failed to execute query", zap.Error(err))
		return nil, err
	}

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
