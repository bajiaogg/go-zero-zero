package logic

import (
	"context"

	"go-zero-zero/service/user/rpc/internal/svc"
	"go-zero-zero/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateAddressLogic {
	return &CreateAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  地址
func (l *CreateAddressLogic) CreateAddress(in *pb.CreateAddressReq) (*pb.SaveAddressResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SaveAddressResp{}, nil
}
