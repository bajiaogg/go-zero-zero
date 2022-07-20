package response

import (
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-zero/common/xerr"
	"google.golang.org/grpc/status"
	"net/http"
)

func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		r := Success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		// 错误
		errCode := xerr.DefaultErrorCode
		errMsg := xerr.MapErrMsg(errCode)

		causeErr := errors.Cause(err)
		// 自定义错误类型, api
		if e, ok := causeErr.(*xerr.Xerror); ok {
			errCode = e.GetErrCode()
			errMsg = e.GetErrMsg()
		} else {
			// grpc err 需要搭配拦截器
			if gstatus, ok := status.FromError(causeErr); ok {
				grpcCode := uint32(gstatus.Code())
				if xerr.IsCodeErr(grpcCode) {
					errCode = grpcCode
					errMsg = gstatus.Message()
				}
			}
		}
		httpx.WriteJson(w, http.StatusBadRequest, Error(errCode, errMsg))
	}
}
