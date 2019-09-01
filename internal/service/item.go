package service

import (
	"context"
	"fmt"
	"kratos-demo/internal/model"
)

func (s *Service) ItemList(c context.Context, arg *model.ArgItemList) (res []*model.Items, count int64, err error) {
	fmt.Println("aaa", s)
	s.dao.Ping(c)
	//if count, err = s.dao.OrderCount(c, arg); err != nil {
	//	err = errors.WithStack(err)
	//	return
	//}
	//if res, err = s.dao.OrderList(c, arg); err != nil {
	//	err = errors.WithStack(err)
	//}
	return
}
