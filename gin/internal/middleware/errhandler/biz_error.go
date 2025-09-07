package errhandler

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

func NewBizError(m string, httpcode int, bizcode int) *bizError {
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
var (
	ErrBadRequest  = NewBizError("bad request", 400, 10400)
	ErrForbidden   = NewBizError("forbidden", 403, 10403)
	ErrNotFound    = NewBizError("resource not found", 404, 10404)
	ErrServerError = NewBizError("internal server error", 500, 10500)
)
