package user

import (
	"context"
	"github.com/jinzhu/copier"
	ctxdata "go-zero-zero/common/ctxData"
	"go-zero-zero/service/user/rpc/userrpc"

	"go-zero-zero/service/user/api/internal/svc"
	"go-zero-zero/service/user/api/internal/types"

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

func (l *GetUserInfoLogic) GetUserInfo(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	infoResp, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &userrpc.GetSingleUserInfoReq{
		Id: userId,
	})
	if err != nil {
		return nil, err
	}

	var userInfo types.User
	_ = copier.Copy(&userInfo, infoResp.User)

	return &types.UserInfoResp{
		UserInfo: userInfo,
	}, nil
}
