package common

type APP_TYPE int

const (
	APP_NONE APP_TYPE = iota
	APP_FIXED_DEPOSIT
	APP_PMS
	APP_NSP
	APP_IPO
	APP_BOND
	APP_MAX_LEN = APP_BOND
)

func (t APP_TYPE) String() string {
	if t <= APP_NONE && t >= APP_MAX_LEN {
		return APP_NONE.String()
	}
	return [...]string{"app_none", "fixed_deposit", "pms", "nps", "ipo", "bond", "app_max"}[t]
}
