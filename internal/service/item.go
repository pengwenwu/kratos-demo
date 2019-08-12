package service

import (
	"context"
	"github.com/pkg/errors"
	"kratos-demo/internal/model"
)

func (s *Service) OrderList(c context.Context, arg *model.ArgItemList) (res []*model.Items, count int64, err error) {
	if count, err = s.dao.OrderCount(c, arg); err != nil {
		err = errors.WithStack(err)
		return
	}
	if res, err = s.dao.OrderList(c, arg); err != nil {
		err = errors.WithStack(err)
	}
	return
}
