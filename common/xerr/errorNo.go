package xerr

// DbError 数据库异常
const DbError uint32 = 5700 // 因为我的mysql版本是5.7

// DefaultCode 默认成功的返回
const DefaultCode uint32 = 200
const ServerCommonError uint32 = 1001

// DefaultErrorCode 业务处理失败
const DefaultErrorCode uint32 = 40004

// 特殊错误码

const TokenGenerateError uint32 = 10001 // 生成token失败
const TokenExpireError uint32 = 10002   // token已过期
const NO_AUTH uint32 = 20001            //权限不足
const OperationError uint32 = 30001     // 操作有误
const MissParameter uint32 = 40001      // 缺少必填参数
const InvalidParameter uint32 = 40002   //  非法参数
const LimitedError uint32 = 40005       // 调用频次超限

// 用户模块
