package model

//ArgItemList query item.
type ArgItemList struct {
	Appkey  string `form:"appkey"`
	Channel int    `form:"channel"`
	Page    int    `form:"page" default:"1"`
	Limit   int    `form:"limit" default:"20"`
}
