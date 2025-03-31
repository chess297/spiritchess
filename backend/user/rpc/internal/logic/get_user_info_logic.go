package logic

import (
	"context"

	"user/rpc/internal/svc"
	"user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 定义一个 GetUserInfo 一元 rpc 方法，请求体和响应体必填。
func (l *GetUserInfoLogic) GetUserInfo(in *user.GetUserInfoReq) (*user.GetUserInfoResp, error) {
	// todo: add your logic here and delete this line
	logc.Info(l.ctx, "GetUserInfo id", in)
	return &user.GetUserInfoResp{
		Data: &user.UserInfo{
			// username
			Id:       "id",
			Username: "username",
		},
	}, nil
}
