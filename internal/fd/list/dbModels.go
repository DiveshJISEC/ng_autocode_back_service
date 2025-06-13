package fdList

type FDAgent struct {
	AgentCode      int    `gorm:"primaryKey;column:agentcode" json:"agentCode"`
	AgentCodeName  string `gorm:"column:agentcodename" json:"agentCodeName"`
	AgentFirstName string `gorm:"column:agentfirstname" json:"agentFirstName"`
	AgentLastName  string `gorm:"column:agentlastname" json:"agentLastName"`
	AgentBranch    string `gorm:"column:agentbranch" json:"agentBranch"`
	AgentCategory  int    `gorm:"column:agentcategory" json:"agentCategory"`
	AgentActive    bool   `gorm:"column:agentactive" json:"agentActive"`
	AgentEmpType   string `gorm:"column:agentemptype" json:"agentEmpType"`
	OtherFlags     string `gorm:"column:otherflags" json:"otherFlags"`
	DtUpdate       string `gorm:"column:dtupdate" json:"dtUpdate"`
}
