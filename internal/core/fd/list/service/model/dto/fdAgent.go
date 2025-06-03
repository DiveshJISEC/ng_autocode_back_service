package dtoModel

type FDAgentListRequest struct {
	UserId        string `json:"FML_USER_ID" binding:"required"`
	SessionCode   string `json:"FML_SSSN_ID" binding:"required"`
	FromDateTime  string `json:"FML_FROM_DT" binding:"required"`
	ToDateTime    string `json:"FML_TO_DT" binding:"required"`
	FdAgentCode   string `json:"FML_AGENT_CD" binding:"required"`
	AgentBranch   string `json:"FML_AGENT_BRANCH" binding:"required"`
	AgentCategory int    `json:"FML_AGENT_CATEGORY" binding:"required"`
	AgentActive   bool   `json:"FML_AGENT_ACTIVE" binding:"required"`
	AgentEmpType  string `json:"FML_AGENT_EMP_TYPE" binding:"required"`
}

type FDAgentListResponse struct {
	AgentCode      int    `json:"agentCode"`
	AgentCodeName  string `json:"agentCodeName"`
	AgentFirstName string `json:"agentFirstName"`
	AgentLastName  string `json:"agentLastName"`
	AgentBranch    string `json:"agentBranch"`
	AgentCategory  int    `json:"agentCategory"`
	AgentActive    bool   `json:"agentActive"`
	AgentEmpType   string `json:"agentEmpType"`
	OtherFlags     string `json:"otherFlags"`
	DtUpdate       string `json:"dtUpdate"`
}
