package fdBook

import "time"

type FDOrderBookRequest struct {
	UserId       string  `json:"FML_USER_ID" binding:"required"`
	SessionCode  string  `json:"FML_SSSN_ID" binding:"required"`
	FromDateTime string  `json:"FML_FROM_DT"`
	ToDateTime   string  `json:"FML_TO_DT"`
	FdAgentCode  string  `json:"FML_AGENT_CD"`
	MatchAccount int     `json:"FML_MATCH_ACCT"`
	FdSchemeName string  `json:"FML_SCHEME_NAME"`
	InterestRate float32 `json:"FML_INTRST_RT"`
	MaturityDays int     `json:"FML_MATURITY_DAYS"`
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
