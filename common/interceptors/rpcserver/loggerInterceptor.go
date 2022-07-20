package rpcserver

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-zero/common/xerr"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func LoggerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	resp, err = handler(ctx, req)
	if err != nil {
		causeErr := errors.Cause(err)
		// 自定义错误类型
		if e, ok := causeErr.(*xerr.Xerror); ok {
			err = status.Error(codes.Code(e.GetErrCode()), e.GetErrMsg())
		}
		logx.WithContext(ctx).Errorf("【rpc-srv-err】 %+v", err)
	}
	return resp, err
}
