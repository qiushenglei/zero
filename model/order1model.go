package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ Order1Model = (*customOrder1Model)(nil)

type (
	// Order1Model is an interface to be customized, add more methods here,
	// and implement the added methods in customOrder1Model.
	Order1Model interface {
		order1Model
	}

	customOrder1Model struct {
		*defaultOrder1Model
	}
)

// NewOrder1Model returns a model for the database table.
func NewOrder1Model(conn sqlx.SqlConn) Order1Model {
	return &customOrder1Model{
		defaultOrder1Model: newOrder1Model(conn),
	}
}
