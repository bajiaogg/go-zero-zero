package logic

import (
	"context"
	"github.com/pkg/errors"
	tool "go-zero-zero/common/tools"
	"go-zero-zero/common/xerr"
	"go-zero-zero/service/user/model"
	"go-zero-zero/service/user/rpc/userrpc"

	"go-zero-zero/service/user/rpc/internal/svc"
	"go-zero-zero/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var ErrGenerateTokenError = xerr.NewErrMsg("生成token失败")
var ErrUsernamePwdError = xerr.NewErrMsg("账号或密码不正确")

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
	var userId int64
	var err error
	switch in.AuthType {
	case model.LoginWithPwd:
		userId, err = l.loginByMobile(in.AuthKey, in.AuthValue)
	default:
		return nil, xerr.NewErrCode(xerr.ServerCommonError)
	}
	if err != nil {
		return nil, err
	}

	//2、Generate the token
	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := generateTokenLogic.GenerateToken(&userrpc.GenerateTokenReq{
		UserId: userId,
	})
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "GenerateToken userId : %d", userId)
	}

	// todo: 1、登录日志 2、把登录放到队列，处理其他奖励活动
	return &userrpc.LoginResp{
		AccessToken:  tokenResp.AccessToken,
		AccessExpire: tokenResp.AccessExpire,
		RefreshAfter: tokenResp.RefreshAfter,
	}, nil
}

func (l *LoginLogic) loginByMobile(mobile, password string) (int64, error) {

	user, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, mobile)
	if err != nil && err != model.ErrNotFound {
		return 0, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "根据手机号查询用户信息失败，mobile:%s,err:%v", mobile, err)
	}
	if user == nil {
		return 0, errors.Wrapf(ErrUserNoExistsError, "mobile:%s", mobile)
	}

	if !(tool.MD5(tool.MD5(password)) == user.Password) {
		return 0, errors.Wrap(ErrUsernamePwdError, "密码匹配出错")
	}

	return user.Id, nil
}
