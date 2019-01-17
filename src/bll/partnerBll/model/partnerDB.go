package model

// 合作商DB接口
type PartnerDB struct {
	PartnerID int

	PartnerName string

	PartnerAlias string

	AppID string

	LoginKey string

	ChargeConfig string

	OtherConfigInfo string

	GameVersionUrl string

	ChargeServerUrl string

	PartnerType int

	Weight int
}
