package repo

import (
	"context"
	"fmt"
	m "ng_autocode_back_service/internal/core/fd/list/service/model/db"
	dto "ng_autocode_back_service/internal/core/fd/list/service/model/dto"
	logger "ng_autocode_back_service/pkg/logger"

	"go.uber.org/zap"
)

func (g *dbSt) GetFDAgentList(c context.Context, dtoRequest *dto.FDAgentListRequest) (dtoResponse []*dto.FDAgentListResponse, e error) {
	logger.Log(c).Info("GetFDAgentList called")

	// Create a slice to hold the database results
	var fdAgentList []m.FDAgent

	// Query the database using raw SQL to avoid quoting iss
	whereQ := ""
	if dtoRequest.AgentBranch != "" {
		whereQ = whereQ + fmt.Sprintf(" where agentBranch = '%s'", dtoRequest.AgentBranch)
	}
	query := g.oracle.WithContext(c).Raw(`
		SELECT agentCode, agentCodeName, agentFirstName, agentLastName,
		       agentBranch, agentCategory, agentActive, agentEmpType,
		       otherFlags, dtUpdate
		FROM fd_agentMaster` + whereQ)

	logger.Log(c).Debug("Executing query", zap.Any("query", query))

	// Execute the query
	if err := query.Scan(&fdAgentList).Error; err != nil {
		logger.Log(c).Error("Failed to fetch FD agent list", zap.Error(err))
		return nil, err
	}

	// Initialize the response slice with proper capacity
	agents := make([]*dto.FDAgentListResponse, 0, len(fdAgentList))

	// Map database entities to DTO responses
	for _, agent := range fdAgentList {
		agents = append(agents, &dto.FDAgentListResponse{
			AgentCode:      agent.AgentCode,
			AgentCodeName:  agent.AgentCodeName,
			AgentFirstName: agent.AgentFirstName,
			AgentLastName:  agent.AgentLastName,
			AgentBranch:    agent.AgentBranch,
			AgentCategory:  agent.AgentCategory,
			AgentActive:    agent.AgentActive,
			AgentEmpType:   agent.AgentEmpType,
			OtherFlags:     agent.OtherFlags,
			DtUpdate:       agent.DtUpdate,
		})
	}

	logger.Log(c).Debug("Successfully retrieved agents", zap.Int("count", len(agents)))
	return agents, nil
}
