package fdList

import (
	"context"
	"fmt"
	"ng_autocode_back_service/common/pkg/logger"

	"go.uber.org/zap"
)

func (g *dbSt) GetFDAgentList(c context.Context, dtoRequest *FDAgentListRequest) (dtoResponse []*FDAgentListResponse, e error) {
	logger.Log(c).Info("GetFDAgentList called")

	// Create a slice to hold the database results
	var fdAgentList []FDAgent

	// Query the database using raw SQL to avoid quoting iss
	whereQ := fmt.Sprintf(" where agentActive = %t", dtoRequest.AgentActive)
	if dtoRequest.AgentBranch != "" {
		whereQ = whereQ + fmt.Sprintf(" and agentBranch = '%s'", dtoRequest.AgentBranch)
	}
	if dtoRequest.FdAgentCode != "" {
		whereQ = whereQ + fmt.Sprintf(" and agentCodeName = '%s'", dtoRequest.FdAgentCode)
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
	agents := make([]*FDAgentListResponse, 0, len(fdAgentList))

	// Map database entities to DTO responses
	for _, agent := range fdAgentList {
		agents = append(agents, &FDAgentListResponse{
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
