package dtoModel

type FDAgentListRequest struct {
	UserId      string `json:"FML_USER_ID" binding:"required"`
	SessionCode string `json:"FML_SSSN_ID" binding:"required"`
}

type FDAgentListResponse struct {
	Name string `json:"FML_AGENT_NAME" binding:"required"`
	Code string `json:"FML_AGENT_CD" binding:"required"`
}
