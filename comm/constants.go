// errors
package comm

import (
	"bytes"
	"runtime"
	"strings"
)

const (
	BASECODE = 50000
)

//backtesting内部错误定义（所有模块）
//命名规则：Err$模块名$错误名  e.g.  ErrCommInvalid
//注：模块名只使用一个单词来标识，多级的组装成某个单词；
//   错误名则单词数量约束，但尽量言简意赅
var (
	ErrNetMgrEmptyAddr      = &ErrorInfo{BASECODE + 500, "netMgr empty address!"}
	ErrNetMgrNewContextFail = &ErrorInfo{BASECODE + 501, "netMgr new context fail!"}
	ErrNetMgrNewSocketFail  = &ErrorInfo{BASECODE + 502, "netMgr new socket fail!"}
	ErrNetMgrBindFail       = &ErrorInfo{BASECODE + 503, "netMgr bind fail!"}
	ErrNetMgrStopServerFail = &ErrorInfo{BASECODE + 504, "netMgr stop server fail"}
	ErrNetMgrSendMsgFail    = &ErrorInfo{BASECODE + 505, "netMgr send msg fail"}
	ErrNetMgrRecvMsgFail    = &ErrorInfo{BASECODE + 506, "netMgr recv msg fail"}
	ErrNetMgrConnectFail    = &ErrorInfo{BASECODE + 507, "netMgr connect fail"}

	ErrNewRouterRequestFail = &ErrorInfo{BASECODE + 550, "routerRequest new request fail"}
)

type ErrorInfo struct {
	code int
	msg  string
}

func (err *ErrorInfo) Code() int {
	return err.code
}

func (err *ErrorInfo) Error() string {
	return err.msg
}

type DropboxError interface {
	GetMessage() string
	GetStack() string
	GetInner() error
	Error() string
}

type DropboxBaseError struct {
	Info    *ErrorInfo
	Stack   string
	Context string
	inner   error
}

func New(info *ErrorInfo) DropboxError {
	stack, context := StackTrace()
	return &DropboxBaseError{
		Info:    info,
		Stack:   stack,
		Context: context,
	}

}

func Wrap(inn error, info *ErrorInfo) DropboxError {
	stack, context := StackTrace()
	return &DropboxBaseError{
		Info:    info,
		Stack:   stack,
		Context: context,
		inner:   inn,
	}
}

func GetMessage(err interface{}) string {
	switch e := err.(type) {
	case DropboxError:
		dberr := DropboxError(e)
		ret := []string{}
		for dberr != nil {
			ret = append(ret, dberr.GetMessage())
			d := dberr.GetInner()
			if d == nil {
				break
			}
			var ok bool
			dberr, ok = d.(DropboxError)
			if !ok {
				ret = append(ret, d.Error())
				break
			}
		}
		return strings.Join(ret, " ")
	case runtime.Error:
		return runtime.Error(e).Error()
	default:
		return "Passed a non-error to GetMessage"
	}
}

func (e *DropboxBaseError) Error() string {
	return DefaultError(e)
}

func (e *DropboxBaseError) GetMessage() string {
	return e.Info.Error()
}

func (e *DropboxBaseError) GetCode() int {
	return e.Info.Code()
}

func (e *DropboxBaseError) GetStack() string {
	return e.Stack
}

func (e *DropboxBaseError) GetInner() error {
	return e.inner
}

func DefaultError(e DropboxError) string {
	errLines := make([]string, 1)
	var origStack string
	errLines[0] = "ERROR:"
	fillErrorInfo(e, &errLines, &origStack)
	errLines = append(errLines, "")

	errLines = append(errLines, "ORIGINAL STACK TRACE:")
	errLines = append(errLines, origStack)

	return strings.Join(errLines, "\n")
}

func fillErrorInfo(err error, errLines *[]string, origStack *string) {
	if err == nil {
		return
	}

	derr, ok := err.(DropboxError)
	if ok {
		*errLines = append(*errLines, derr.GetMessage())
		*origStack = derr.GetStack()
		fillErrorInfo(derr.GetInner(), errLines, origStack)
	} else {
		*errLines = append(*errLines, err.Error())
	}
}

func stackTrace(skip int) (current, context string) {
	buf := make([]byte, 128)
	for {
		n := runtime.Stack(buf, false)
		if n < len(buf) {
			buf = buf[:n]
			break
		}
		buf = make([]byte, len(buf)*2)
	}

	indexNewline := func(b []byte, start int) int {
		if start >= len(b) {
			return len(b)
		}
		searchBuf := b[start:]
		index := bytes.IndexByte(searchBuf, '\n')
		if index == -1 {
			return len(b)
		} else {
			return (start + index)
		}
	}

	var strippedBuf bytes.Buffer
	index := indexNewline(buf, 0)
	if index != -1 {
		strippedBuf.Write(buf[:index])
	}

	for i := 0; i < skip; i++ {
		index = indexNewline(buf, index+1)
		index = indexNewline(buf, index+1)
	}

	isDone := false
	startIndex := index
	lastIndex := index
	for !isDone {
		index = indexNewline(buf, index+1)
		if (index - lastIndex) <= 1 {
			isDone = true
		} else {
			lastIndex = index
		}
	}
	strippedBuf.Write(buf[startIndex:index])
	return strippedBuf.String(), string(buf[index:])
}

func StackTrace() (current, context string) {
	return stackTrace(3)
}
