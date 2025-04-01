// Code generated by goctl. DO NOT EDIT.
// versions:
//  goctl version: 1.8.1

package user

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	usersFieldNames          = builder.RawFieldNames(&Users{})
	usersRows                = strings.Join(usersFieldNames, ",")
	usersRowsExpectAutoSet   = strings.Join(stringx.Remove(usersFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	usersRowsWithPlaceHolder = strings.Join(stringx.Remove(usersFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheSingleDeployUsersIdPrefix = "cache:singleDeploy:users:id:"
)

type (
	usersModel interface {
		Insert(ctx context.Context, data *Users) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Users, error)
		Update(ctx context.Context, data *Users) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUsersModel struct {
		sqlc.CachedConn
		table string
	}

	Users struct {
		Id         int64     `db:"id"`
		CreateTime time.Time `db:"create_time"` // 创建时间
		UpdateTime time.Time `db:"update_time"`
		Username   string    `db:"username"` // 用户名
		Password   string    `db:"password"` // 密码
		Uid        string    `db:"uid"`      // 用户ID
	}
)

func newUsersModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultUsersModel {
	return &defaultUsersModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`users`",
	}
}

func (m *defaultUsersModel) Delete(ctx context.Context, id int64) error {
	singleDeployUsersIdKey := fmt.Sprintf("%s%v", cacheSingleDeployUsersIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, singleDeployUsersIdKey)
	return err
}

func (m *defaultUsersModel) FindOne(ctx context.Context, id int64) (*Users, error) {
	singleDeployUsersIdKey := fmt.Sprintf("%s%v", cacheSingleDeployUsersIdPrefix, id)
	var resp Users
	err := m.QueryRowCtx(ctx, &resp, singleDeployUsersIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", usersRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
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

func (m *defaultUsersModel) Insert(ctx context.Context, data *Users) (sql.Result, error) {
	singleDeployUsersIdKey := fmt.Sprintf("%s%v", cacheSingleDeployUsersIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, usersRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Username, data.Password, data.Uid)
	}, singleDeployUsersIdKey)
	return ret, err
}

func (m *defaultUsersModel) Update(ctx context.Context, data *Users) error {
	singleDeployUsersIdKey := fmt.Sprintf("%s%v", cacheSingleDeployUsersIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, usersRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Username, data.Password, data.Uid, data.Id)
	}, singleDeployUsersIdKey)
	return err
}

func (m *defaultUsersModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheSingleDeployUsersIdPrefix, primary)
}

func (m *defaultUsersModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", usersRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUsersModel) tableName() string {
	return m.table
}
