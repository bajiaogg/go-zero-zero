package logic

import (
	"context"

	"go-zero-zero/service/user/rpc/internal/svc"
	"go-zero-zero/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddressListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddressListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddressListLogic {
	return &AddressListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddressListLogic) AddressList(in *pb.GetAddressListReq) (*pb.GetAddressListResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetAddressListResp{}, nil
}
