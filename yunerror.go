package pb

import "fmt"

/*
Error定义规则
0： 成功
100-199： 请求相关错误，一般是由于请求格式的输入错误导致。
200-299： 授权相关错误，包括访问令牌、业务授权等
300-399： 预先定义的服务内部错误
1000： 系统内部错误
>1000：开发者自定义错误，但描述统一以系统内部错误返回
*/

type EnumYunErr uint32

const YunErrBegin EnumYunErr = 0

const (
	//系统错误
	YUNERR_SUCCESS                                EnumYunErr = 0                  //成功（不可加上起始值）
	YUNERR_INVALID_PARAMETER                      EnumYunErr = YunErrBegin + 100  //请求参数无效
	YUNERR_UNSUPPORTED_METHOD                     EnumYunErr = YunErrBegin + 101  //未知的方法
	YUNERR_URLBODYTOOLARGE_METHOD                 EnumYunErr = YunErrBegin + 102  //请求body过长
	YUNERR_INVALID_CONTENTLENGTH_METHOD           EnumYunErr = YunErrBegin + 103  //无效的ContentLength
	YUNERR_QID_ALREADY_EXISTS                     EnumYunErr = YunErrBegin + 104  //qid已存在
	YUNERR_TOKEN_INVALID                          EnumYunErr = YunErrBegin + 201  //无效的访问令牌
	YUNERR_TOKEN_EXPIRED                          EnumYunErr = YunErrBegin + 202  //访问令牌过期
	YUNERR_INCORRECT_SIGNATURE                    EnumYunErr = YunErrBegin + 203  //无效的签名
	YUNERR_UNAUTHORIZED_METHOD                    EnumYunErr = YunErrBegin + 210  //未授权的方法
	YUNERR_UNAUTHORIZED_PARAMETER                 EnumYunErr = YunErrBegin + 211  //未授权的参数
	YUNERR_OBJECT_NOT_FOUND                       EnumYunErr = YunErrBegin + 301  //指定的对象不存在
	YUNERR_OBJECT_EXISTS                          EnumYunErr = YunErrBegin + 302  //指定的对象已存在
	YUNERR_OBJECT_DATA_NOT_FOUND                  EnumYunErr = YunErrBegin + 303  //指定的对象数据不存在
	YUNERR_SERVICE_BUSY                           EnumYunErr = YunErrBegin + 304  //总线返回服务忙
	YUNERR_INTERNAL_SERVICE_ACCESS_ERROR_UNREPORT EnumYunErr = YunErrBegin + 499  //内部服务访问错误，不上报
	YUNERR_SERVICE_UNKNOWN_ERROR                  EnumYunErr = YunErrBegin + 501  //服务未知错误
	YUNERR_INVALID_OPERATION                      EnumYunErr = YunErrBegin + 502  //无效的操作方法
	YUNERR_DATABASE_ERROR                         EnumYunErr = YunErrBegin + 503  //数据库操作出错，请重试
	YUNERR_SERVICE_TEMPORARILY_UNAVAILABLE        EnumYunErr = YunErrBegin + 504  //服务暂不可用
	YUNERR_REQUEST_TIMEOUT                        EnumYunErr = YunErrBegin + 505  //应用总线返回请求超时
	YUNERR_INTERNAL_SERVICE_ACCESS_ERROR          EnumYunErr = YunErrBegin + 510  //内部服务访问错误
	YUNERR_APPBUS_LOGIN                           EnumYunErr = YunErrBegin + 513  //总线登陆失败
	YUNERR_SEND_TIMEOUT                           EnumYunErr = YunErrBegin + 901  //发送超时（总线SDK）
	YUNERR_SYSTEM_INTERNAL_ERROR                  EnumYunErr = YunErrBegin + 1000 //系统内部错误

	//开发者自定义错误
	//警告：一般情况下，开发者不需要自定义错误，除非极特殊的情况
	//并且定义的值必须大于（YunErrBegin+1000），且不需要在EnumYunErr_desc做错误描述映射
	YUNERR_USER_TEST = YunErrBegin + 1001
)

var EnumYunErr_desc = map[EnumYunErr]string{
	YUNERR_SUCCESS:                                "Success",
	YUNERR_INVALID_PARAMETER:                      "Invalid parameter",
	YUNERR_UNSUPPORTED_METHOD:                     "Unsupported method",
	YUNERR_URLBODYTOOLARGE_METHOD:                 "Request body too large",
	YUNERR_INVALID_CONTENTLENGTH_METHOD:           "ContentLength invalid",
	YUNERR_QID_ALREADY_EXISTS:                     "Qid already exists",
	YUNERR_TOKEN_INVALID:                          "Token invalid or no longer valid",
	YUNERR_TOKEN_EXPIRED:                          "Token expired",
	YUNERR_INCORRECT_SIGNATURE:                    "Incorrect signature",
	YUNERR_UNAUTHORIZED_METHOD:                    "Unauthorized method",
	YUNERR_UNAUTHORIZED_PARAMETER:                 "Unauthorized parameter",
	YUNERR_INTERNAL_SERVICE_ACCESS_ERROR_UNREPORT: "Internal service access error unreport",
	YUNERR_SERVICE_UNKNOWN_ERROR:                  "Service unknown error",
	YUNERR_SERVICE_TEMPORARILY_UNAVAILABLE:        "Service temporarily unavailable",
	YUNERR_OBJECT_NOT_FOUND:                       "Specified object cannot be found",
	YUNERR_OBJECT_EXISTS:                          "Specified object already exists",
	YUNERR_OBJECT_DATA_NOT_FOUND:                  "Specified object data cannot be found",
	YUNERR_INVALID_OPERATION:                      "Invalid operation",
	YUNERR_DATABASE_ERROR:                         "A database error occurred. Please try again",
	YUNERR_SYSTEM_INTERNAL_ERROR:                  "System internal error",
	YUNERR_SERVICE_BUSY:                           "Service busy now",
	YUNERR_INTERNAL_SERVICE_ACCESS_ERROR:          "Internal service access error",
}

func (x EnumYunErr) Code() uint32 {
	return uint32(x)
}

func (x EnumYunErr) Desc() string {
	var key EnumYunErr
	switch {
	//用户自定义错误码
	case x >= YUNERR_SERVICE_UNKNOWN_ERROR:
		key = YUNERR_SERVICE_UNKNOWN_ERROR
	//旧版本错误码
	case (x != YUNERR_SUCCESS && x < YunErrBegin):
		key = YUNERR_SERVICE_UNKNOWN_ERROR
	//其他新版本错误码
	default:
		key = x
	}
	desc, ok := EnumYunErr_desc[key]
	if !ok && key != YUNERR_SERVICE_UNKNOWN_ERROR {
		desc, _ = EnumYunErr_desc[YUNERR_SERVICE_UNKNOWN_ERROR]
	}
	return desc
}

func (x EnumYunErr) String() string {
	return fmt.Sprintf("{\"code\":%d,\"desc\":\"%s\"}", x.Code(), x.Desc())
}

func (x EnumYunErr) Error() string {
	return x.String()
}
