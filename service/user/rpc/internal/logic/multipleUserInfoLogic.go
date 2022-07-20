package logic

import (
	"context"

	"go-zero-zero/service/user/rpc/internal/svc"
	"go-zero-zero/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type MultipleUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMultipleUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MultipleUserInfoLogic {
	return &MultipleUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MultipleUserInfoLogic) MultipleUserInfo(in *pb.GetMultipleUserInfoReq) (*pb.GetMultipleUserInfoResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetMultipleUserInfoResp{}, nil
}
