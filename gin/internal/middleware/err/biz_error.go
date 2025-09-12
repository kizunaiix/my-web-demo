package err

type BizError interface {
	error
	StatusCode() int // HTTP状态码
	BizCode() int    // 业务错误码
}

var _ BizError = (*bizError)(nil)

type bizError struct {
	msg      string
	httpCode int
	bizCode  int
}

func New(m string, httpcode int, bizcode int) BizError {
	if httpcode == 0 {
		httpcode = 200
	}

	return &bizError{
		msg:      m,
		httpCode: httpcode,
		bizCode:  bizcode,
	}
}

func (r *bizError) Error() string { return r.msg }

func (r *bizError) StatusCode() int { return r.httpCode }

func (r *bizError) BizCode() int { return r.bizCode }

// 自定义标准错误
func ErrBadRequest(msg string) BizError  { return New(msg, 400, 10400) }
func ErrForbidden(msg string) BizError   { return New(msg, 403, 10403) }
func ErrNotFound(msg string) BizError    { return New(msg, 404, 10404) }
func ErrServerError(msg string) BizError { return New(msg, 500, 10500) }
