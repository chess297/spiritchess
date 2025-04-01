package user

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UsersModel = (*customUsersModel)(nil)

type (
	// UsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUsersModel.
	UsersModel interface {
		usersModel
		FindOneByUid(ctx context.Context, uid string) (*Users, error)
	}

	customUsersModel struct {
		*defaultUsersModel
	}
)

// NewUsersModel returns a model for the database table.
func NewUsersModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UsersModel {
	return &customUsersModel{
		defaultUsersModel: newUsersModel(conn, c, opts...),
	}
}

func (m *defaultUsersModel) FindOneByUid(ctx context.Context, uid string) (*Users, error) {
	singleDeployUsersIdKey := fmt.Sprintf("%s%v", cacheSingleDeployUsersIdPrefix, uid)
	var resp Users
	err := m.QueryRowCtx(ctx, &resp, singleDeployUsersIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `uid` = ? limit 1", usersRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, uid)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
