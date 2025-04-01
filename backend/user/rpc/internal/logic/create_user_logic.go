package logic

import (
	"context"

	model_user "user/model/user"
	"user/rpc/internal/svc"
	"user/rpc/internal/utils"
	"user/rpc/pb/user"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *user.CreateUserReq) (*user.CreateUserResp, error) {
	pwd, err := utils.EncryptPassword(in.Password)
	if err != nil {
		return &user.CreateUserResp{
			Success: false,
			Message: err.Error(),
		}, nil
	}
	uid := uuid.New().String()
	_, err = l.svcCtx.UserModel.Insert(l.ctx, &model_user.Users{
		Uid:      uid,
		Username: in.Username,
		Password: pwd,
	})
	if err != nil {
		return &user.CreateUserResp{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &user.CreateUserResp{
		Success: true,
		Message: "创建成功",
		Uid:     uid,
	}, nil
}
