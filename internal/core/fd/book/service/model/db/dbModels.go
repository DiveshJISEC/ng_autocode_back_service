package dbModel

import "time"

type FDOrderBook struct {
	TransId      int       `gorm:"column:transid"`
	UserId       string    `gorm:"column:userid"`
	MatchAccount int       `gorm:"column:matchaccount"`
	FdSchemeName string    `gorm:"column:fdschemename"`
	InterestRate float32   `gorm:"column:interestrate"`
	Amount       int       `gorm:"column:amount"`
	MaturityDays int       `gorm:"column:maturitydays"`
	AutoRenew    bool      `gorm:"column:autorenew"`
	FdAgentCode  string    `gorm:"column:fdagentcode"`
	OtherFlags   string    `gorm:"column:otherflags"`
	DtUpdate     time.Time `gorm:"column:dtupdate"`
}
