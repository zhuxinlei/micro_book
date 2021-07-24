package common

// 全局的 错误号 类型，用于API调用之间传递
type ServerErrorxCode int

// 全局的 错误号 的具体定义
const (
	InsertBookErrorCode = 2000 //插入book出错
	GetBookErrorCode    = 2001 //获取book信息出错
	GetBooksErrorCode    = 2002 //批量获取book信息出错

)

// 内部的错误map，用来对应 错误号和错误信息
var ErrCodeMap = map[ServerErrorxCode]string{

	InsertBookErrorCode: "insert into Books error",
	GetBookErrorCode:    "get Book info error",
	GetBooksErrorCode:    "get Book infos error",
}

// Sentinel Error： 即全局定义的Static错误变量
// 注意，这里的全局error是没有保存堆栈信息的，所以需要在初始调用处使用 errors.Wrap
var (
	InsertBookError = NewServerErrorx(InsertBookErrorCode)
	GetBookError = NewServerErrorx(GetBookErrorCode)
	GetBooksError = NewServerErrorx(GetBooksErrorCode)
)

func NewServerErrorx(code ServerErrorxCode) *ServerErrorx {
	return &ServerErrorx{
		Code:    code,
		Message: ErrCodeMap[code],
	}
}

// error的具体实现
type ServerErrorx struct {
	// 对外使用 - 错误码
	Code ServerErrorxCode
	// 对外使用 - 错误信息
	Message string
}

func (e *ServerErrorx) Error() string {
	return e.Message
}
