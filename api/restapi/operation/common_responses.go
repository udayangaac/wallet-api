package operation

// ErrorResp a model for common error response.
type ErrorResp struct {
	Error string `json:"error"`
}

// NewErrorResponse create an instance of common error response.
func NewErrorResponse(err error) ErrorResp {
	return ErrorResp{
		Error: err.Error(),
	}
}

// SuccessResp a model for common success response.
type SuccessResp struct {
	Msg string `json:"msg"`
}

// NewSuccessResponse create an instance of common success response.
func NewSuccessResponse(msg string) SuccessResp {
	return SuccessResp{
		Msg: msg,
	}
}
