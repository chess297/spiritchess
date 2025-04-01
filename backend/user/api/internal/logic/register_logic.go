package logic

import (
	"context"
	"time"

	"user/api/internal/svc"
	"user/api/internal/types"
	"user/api/internal/utils/token"
	"user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	res, err := l.svcCtx.UserRpcClient.CreateUser(l.ctx, &user.CreateUserReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	expireAt := int(time.Now().Unix() + int64(l.svcCtx.Config.Auth.AccessExpire))
	token, err := token.GenJwtToken(l.svcCtx.Config.Auth.AccessSecret, int(time.Now().Unix()), l.svcCtx.Config.Auth.AccessExpire, token.JwtTokenPayload{
		UserID:   res.Uid,
		UserName: req.Username,
	})
	if err != nil {
		return nil, err
	}
	u, err := l.svcCtx.UserRpcClient.GetUserInfo(l.ctx, &user.GetUserInfoReq{
		Id: res.Uid,
	})
	if err != nil {
		return nil, err
	}
	return &types.RegisterResp{
		Id:       res.Uid,
		Username: u.Username,
		Token:    token,
		ExpireAt: expireAt,
	}, nil
}
