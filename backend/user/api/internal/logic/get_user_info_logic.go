package logic

import (
	"context"

	"user/api/internal/svc"
	"user/rpc/user_client"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo() (resp *user_client.GetUserInfoResp, err error) {
	// todo: add your logic here and delete this line
	// uc:= userclient.NewUser()
	uInfo, err := l.svcCtx.UserRpcClient.GetUserInfo(l.ctx, &user_client.GetUserInfoReq{
		Id: "id",
	})
	return uInfo, err
}
