package xerr

import "fmt"

type Xerror struct {
	Errno uint32 `json:"errno"`
	Msg   string `json:"msg"`
}

func (e *Xerror) GetErrCode() uint32 {
	return e.Errno
}

func (e *Xerror) GetErrMsg() string {
	return e.Msg
}

func (e *Xerror) Error() string {
	return fmt.Sprintf("Errno:%d，ErrMsg:%s", e.Errno, e.Msg)
}

func NewDefaultError(msg string) *Xerror {
	return &Xerror{DefaultErrorCode, msg}
}

func NewErrCodeMsg(errCode uint32, msg string) *Xerror {
	return &Xerror{Errno: errCode, Msg: msg}
}

func NewErrCode(errCode uint32) *Xerror {
	return &Xerror{Errno: errCode, Msg: MapErrMsg(errCode)}
}

func NewErrMsg(errMsg string) *Xerror {
	return &Xerror{Errno: ServerCommonError, Msg: errMsg}
}
