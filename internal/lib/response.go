package lib

const (
	CODE_FAIL    = -1
	CODE_WARN    = 0
	CODE_SUCCESS = 1
)

var ResponseClient = NewDefaultResponse()

// Response : 统一返回的类型
type Response struct {
	Code    int `json:"code"`
	Message any `json:"message"`
	Data    any `json:"data"`
}

// NewDefaultResponse: 构建默认的统一返回对象
func NewDefaultResponse() *Response {
	return &Response{
		Code:    CODE_WARN,
		Message: "nothing",
		Data:    nil,
	}
}

func (r *Response) Error(message any) *Response {
	r.Code = CODE_FAIL
	r.Message = message

	return r
}

func (r *Response) Warn(message any) *Response {
	r.Message = message
	return r
}

func (r *Response) Success(message string, data any) *Response {
	r.Code = CODE_SUCCESS
	if len(message) != 0 {
		r.Message = message
	}
	r.Data = data
	return r
}
