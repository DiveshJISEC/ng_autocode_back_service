package dbModel

import "time"

type FDOrderBook struct {
	TransId      int       `gorm:"column:transId"`
	UserId       string    `gorm:"column:userId"`
	MatchAccount int       `gorm:"column:matchAccount"`
	FdSchemeName string    `gorm:"column:fdSchemeName"`
	InterestRate float32   `gorm:"column:interestRate"`
	Amount       int       `gorm:"column:amount"`
	MaturityDays int       `gorm:"column:maturityDays"`
	AutoRenew    bool      `gorm:"column:autoRenew"`
	FdAgentCode  string    `gorm:"column:fdAgentCode"`
	OtherFlags   string    `gorm:"column:otherFlags"`
	DtUpdate     time.Time `gorm:"column:dtUpdate"`
}
