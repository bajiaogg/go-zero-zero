package user

import (
	"context"
	"github.com/jinzhu/copier"
	"go-zero-zero/service/user/model"
	"go-zero-zero/service/user/rpc/userrpc"

	"go-zero-zero/service/user/api/internal/svc"
	"go-zero-zero/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (*types.LoginResp, error) {
	loginResp, err := l.svcCtx.UserRpc.Login(l.ctx, &userrpc.LoginReq{
		AuthType:  model.LoginWithPwd,
		AuthKey:   req.Mobile,
		AuthValue: req.Password,
	})
	if err != nil {
		return nil, err
	}

	var resp types.LoginResp
	_ = copier.Copy(&resp, loginResp)

	return &resp, nil
}
