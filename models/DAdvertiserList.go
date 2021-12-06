package models

type DAdvertiserList struct {
	Id            int
	AdverId       string
	AccountId     int
	ChannelId     int
	OperationCode int
}

func (DAdvertiserList) TableName() string {
	return "d_advertiser_list"
}
