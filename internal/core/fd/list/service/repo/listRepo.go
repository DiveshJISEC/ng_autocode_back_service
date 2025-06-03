package repo

import (
	"context"
	m "ng_autocode_back_service/internal/core/fd/list/service/model/db"
	dto "ng_autocode_back_service/internal/core/fd/list/service/model/dto"
	logger "ng_autocode_back_service/pkg/logger"

	"go.uber.org/zap"
)

func (g *dbSt) GetFDAgentList(c context.Context) (dtoResponse []*dto.FDAgentListResponse, e error) {
	logger.Log(c).Info("GetFDAgentList called")

	// Create a slice to hold the results
	var agents []*dto.FDAgentListResponse

	// Query the database for FDAgent list
	var fdAgentList []m.FDAgent
	query := g.oracle.WithContext(c).Table("fd_agent_master").
		Select(
			`agentCode`,
			`agentCodeName`,
			`agentFirstName`,
			`agentLastName`,
			`agentBranch`,
			`agentCategory`,
			`agentActive`,
			`agentEmpType`,
			`otherFlags`,
			`dtUpdate`,
		)

	logger.Log(c).Debug("Executing query", zap.Any("query", query))
	err := query.Scan(&fdAgentList).Error
	if err != nil {
		logger.Log(c).Error("Failed to execute query", zap.Error(err))
		return nil, err
	}

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

	logger.Log(c).Debug("Query result", zap.Any("agents", agents))
	return agents, nil
}
