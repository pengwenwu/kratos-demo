package http

import (
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	"kratos-demo/internal/model"
)

func itemList(c *bm.Context) {
	arg := new(model.ArgItemList)
	if err := c.Bind(arg); err != nil {
		return
	}
	res, count, err :=  svc.ItemList(c, arg)
	info := new(model.PageInfo)
	info.Count = int(count)
	info.Item = res
	info.CurrentPage = arg.Page
	c.JSON(info, err)
}
