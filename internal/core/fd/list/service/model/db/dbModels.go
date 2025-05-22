package dbModel

type FDAgent struct {
	Code string `gorm:"column:FML_AGENT_CD"`
	Name string `gorm:"column:FML_DESC"`
}
