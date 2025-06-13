package dtoModel

type FDAgentListRequest struct {
	UserId        string `json:"FML_USER_ID" binding:"required"`
	SessionCode   string `json:"FML_SSSN_ID" binding:"required"`
	FromDateTime  string `json:"FML_FROM_DT"`
	ToDateTime    string `json:"FML_TO_DT"`
	FdAgentCode   string `json:"FML_AGENT_CD"`
	AgentBranch   string `json:"FML_AGENT_BRANCH"`
	AgentCategory int    `json:"FML_AGENT_CATEGORY"`
	AgentActive   bool   `json:"FML_AGENT_ACTIVE"`
	AgentEmpType  string `json:"FML_AGENT_EMP_TYPE"`
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
