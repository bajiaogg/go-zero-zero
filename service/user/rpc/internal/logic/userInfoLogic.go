package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	tool "go-zero-zero/common/tools"
	"go-zero-zero/common/xerr"
	"go-zero-zero/service/user/model"
	"go-zero-zero/service/user/rpc/userrpc"

	"go-zero-zero/service/user/rpc/internal/svc"
	"go-zero-zero/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var ErrUserNoExistsError = xerr.NewErrMsg("用户不存在")

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *pb.GetSingleUserInfoReq) (*pb.GetSingleUserInfoResp, error) {
	userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "GetUserInfo find user db err , id:%d , err:%v", in.Id, err)
	}
	if userInfo == nil {
		return nil, errors.Wrapf(ErrUserNoExistsError, "id:%d", in.Id)
	}
	userInfo.Mobile = tool.MobileEncrypt(userInfo.Mobile)
	var respUser userrpc.User
	_ = copier.Copy(&respUser, userInfo)

	return &userrpc.GetSingleUserInfoResp{
		User: &respUser,
	}, nil
}
