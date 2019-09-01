package dao

import (
	"context"
	"fmt"
	"github.com/bilibili/kratos/pkg/database/sql"
	"github.com/pkg/errors"
	"kratos-demo/internal/model"
)

// ItemCount item count.
func (d *Dao) ItemCount(c context.Context, arg *model.ArgItemList) (count int64, err error) {

	sqlStr := _itemCountSQL

	var row = d.db.QueryRow(c, sqlStr)
	if err = row.Scan(&count); err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

// ItemList item list.
func (d *Dao) ItemList(c context.Context, arg *model.ArgItemList) (res []*model.Items, err error) {
	sqlStr := _itemListSQL

	if arg.Page <= 0 {
		arg.Page = _defPage
	}
	if arg.Limit <= 0  {
		arg.Limit = _defPageSize
	}
	sqlStr += fmt.Sprintf(" ORDER BY ID DESC LIMIT %v,%v", (arg.Page-1)*arg.Limit, arg.Limit)
	var rows *sql.Rows
	if rows, err = d.db.Query(c, sqlStr); err != nil {
		err = errors.WithStack(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		r := new(model.Items)
		if err = rows.Scan(&r.ID, &r.Appkey, &r.Channel, &r.Name, &r.Photo, &r.Datail, &r.LastDated, &r.Dated); err != nil {
			err = errors.WithStack(err)
			res = nil
			return
		}
		res = append(res, r)
	}
	err = rows.Err()
	return
}
