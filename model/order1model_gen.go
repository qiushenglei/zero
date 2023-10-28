// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	order1FieldNames          = builder.RawFieldNames(&Order1{})
	order1Rows                = strings.Join(order1FieldNames, ",")
	order1RowsExpectAutoSet   = strings.Join(stringx.Remove(order1FieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	order1RowsWithPlaceHolder = strings.Join(stringx.Remove(order1FieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	order1Model interface {
		Insert(ctx context.Context, data *Order1) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Order1, error)
		FindOneByOrderId(ctx context.Context, orderId string) (*Order1, error)
		Update(ctx context.Context, data *Order1) error
		Delete(ctx context.Context, id int64) error
	}

	defaultOrder1Model struct {
		conn  sqlx.SqlConn
		table string
	}

	Order1 struct {
		Id         int64     `db:"id"`
		OrderId    string    `db:"order_id"`
		AppId      string    `db:"app_id"`
		Fee        int64     `db:"fee"`
		AddTime    time.Time `db:"add_time"`
		UpdateTime time.Time `db:"update_time"`
		IsDeleted  int64     `db:"is_deleted"`
	}
)

func newOrder1Model(conn sqlx.SqlConn) *defaultOrder1Model {
	return &defaultOrder1Model{
		conn:  conn,
		table: "`order1`",
	}
}

func (m *defaultOrder1Model) withSession(session sqlx.Session) *defaultOrder1Model {
	return &defaultOrder1Model{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`order1`",
	}
}

func (m *defaultOrder1Model) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultOrder1Model) FindOne(ctx context.Context, id int64) (*Order1, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", order1Rows, m.table)
	var resp Order1
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultOrder1Model) FindOneByOrderId(ctx context.Context, orderId string) (*Order1, error) {
	var resp Order1
	query := fmt.Sprintf("select %s from %s where `order_id` = ? limit 1", order1Rows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, orderId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultOrder1Model) Insert(ctx context.Context, data *Order1) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, order1RowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.OrderId, data.AppId, data.Fee, data.AddTime, data.IsDeleted)
	return ret, err
}

func (m *defaultOrder1Model) Update(ctx context.Context, newData *Order1) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, order1RowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.OrderId, newData.AppId, newData.Fee, newData.AddTime, newData.IsDeleted, newData.Id)
	return err
}

func (m *defaultOrder1Model) tableName() string {
	return m.table
}
