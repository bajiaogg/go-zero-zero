package user

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"go-zero-zero/service/user/model"
	"go-zero-zero/service/user/rpc/userrpc"

	"go-zero-zero/service/user/api/internal/svc"
	"go-zero-zero/service/user/api/internal/types"

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

func (l *RegisterLogic) Register(req *types.RegisterReq) (*types.RegisterResp, error) {
	registerResp, err := l.svcCtx.UserRpc.Register(l.ctx, &userrpc.RegisterReq{
		Mobile:   req.Mobile,
		Password: req.Password,
		RegMode:  model.RegisterByPwd,
		Name:     "admin",
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	var resp types.RegisterResp
	_ = copier.Copy(&resp, registerResp)
	return &resp, nil
}
