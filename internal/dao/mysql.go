package dao

const (
	// item
	_itemListSQL  = "SELECT id,appkey,channel,name,photo,detail,last_dated,dated FROM items WHERE 1=1"
	_itemCountSQL = "SELECT COUNT(1) FROM item WHERE 1=1"

	_defPage     = 1
	_defPageSize = 20
)
