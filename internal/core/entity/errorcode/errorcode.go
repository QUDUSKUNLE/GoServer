package errorcode

type ErrorCode string

const (
	Success    			ErrorCode = "SUCCESS"
	InvalidRequest  ErrorCode = "INVALID_REQUEST"
	DuplicateUser   ErrorCode = "DUPLICATE_USER"
	InternalError   ErrorCode = "INTERNAL_ERROR"
)

const (
	SuccessErrorMessage = "success"
	InternalErrorMessage = "internal error"
	InvalidRequestErrorMessage = "invalid request"
)
