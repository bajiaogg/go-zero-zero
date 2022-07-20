package logic

import (
	"context"
	"github.com/pkg/errors"
	"go-zero-zero/common/stringx"
	tool "go-zero-zero/common/tools"
	"go-zero-zero/service/user/model"
	"go-zero-zero/service/user/rpc/userrpc"

	"go-zero-zero/service/user/rpc/internal/svc"
	"go-zero-zero/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.RegisterReq) (*pb.RegisterResp, error) {
	var (
		user *model.User
		err  error
	)

	switch in.RegMode {
	case model.RegisterByPwd:
		user, err = l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	case model.RegisterBySmsCode:
		user, err = l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	default:
		return &pb.RegisterResp{}, nil
	}

	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(errors.New("查询失败"), "mobile:%s, err:%v", in.Mobile, err)
	}
	if user != nil {
		return nil, errors.Wrapf(errors.New("用户已经存在"), "用户已经存在 mobile:%s", in.Mobile)
	}

	newUser := new(model.User)
	newUser.Mobile = in.Mobile
	if len(in.Password) > 0 {
		newUser.Password = tool.MD5(tool.MD5(in.Password))
	}
	if len(newUser.Nickname) == 0 {
		newUser.Nickname = "用户" + stringx.RandomString(8)
	}
	// 提前生成一批头像，然后随机获取一个
	// todo: 事务
	insertResult, err := l.svcCtx.UserModel.Insert(l.ctx, newUser)
	if err != nil {
		return nil, errors.Wrapf(errors.New("保存用户记录失败"), "mobile:%s", in.Mobile)
	}
	// todo: 注册成功发送通知

	lastId, err := insertResult.LastInsertId()

	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := generateTokenLogic.GenerateToken(&userrpc.GenerateTokenReq{
		UserId: lastId,
	})
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "GenerateToken userId : %d", lastId)
	}

	return &userrpc.RegisterResp{
		AccessToken:  tokenResp.AccessToken,
		AccessExpire: tokenResp.AccessExpire,
		RefreshAfter: tokenResp.RefreshAfter,
	}, nil
}
