package dbModel

type FDAgent struct {
	AgentCode      int    `gorm:"primaryKey;column:agentCode" json:"agentCode"`
	AgentCodeName  string `gorm:"column:agentCodeName" json:"agentCodeName"`
	AgentFirstName string `gorm:"column:agentFirstName" json:"agentFirstName"`
	AgentLastName  string `gorm:"column:agentLastName" json:"agentLastName"`
	AgentBranch    string `gorm:"column:agentBranch" json:"agentBranch"`
	AgentCategory  int    `gorm:"column:agentCategory" json:"agentCategory"`
	AgentActive    bool   `gorm:"column:agentActive" json:"agentActive"`
	AgentEmpType   string `gorm:"column:agentEmpType" json:"agentEmpType"`
	OtherFlags     string `gorm:"column:otherFlags" json:"otherFlags"`
	DtUpdate       string `gorm:"column:dtUpdate" json:"dtUpdate"`
}
