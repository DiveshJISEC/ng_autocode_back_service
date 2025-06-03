package dtoModel

import "time"

type FDOrderBookRequest struct {
	UserId       string  `json:"FML_USER_ID" binding:"required"`
	SessionCode  string  `json:"FML_SSSN_ID" binding:"required"`
	FromDateTime string  `json:"FML_FROM_DT" binding:"required"`
	ToDateTime   string  `json:"FML_TO_DT" binding:"required"`
	FdAgentCode  string  `json:"FML_AGENT_CD" binding:"required"`
	MatchAccount int     `json:"FML_MATCH_ACCT" binding:"required"`
	FdSchemeName string  `json:"FML_SCHEME_NAME" binding:"required"`
	InterestRate float32 `json:"FML_INTRST_RT" binding:"required"`
	MaturityDays int     `json:"FML_MATURITY_DAYS" binding:"required"`
}

type FDOrderBookResponse struct {
	TransId      int       `json:"transId"`
	UserId       string    `json:"userId"`
	MatchAccount int       `json:"matchAccount"`
	FdSchemeName string    `json:"fdSchemeName"`
	InterestRate float32   `json:"interestRate"`
	Amount       int       `json:"amount"`
	MaturityDays int       `json:"maturityDays"`
	AutoRenew    bool      `json:"autoRenew"`
	FdAgentCode  string    `json:"fdAgentCode"`
	OtherFlags   string    `json:"otherFlags"`
	DtUpdate     time.Time `json:"dtUpdate"`
}
