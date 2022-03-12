package responses

type Responses interface {
	InternalServerError() appError
	DataNotFound() appError
	Success() appSuccess
}

//business error
type appError struct {
	Code    int
	Message string
	Data    interface{}
}

type appSuccess struct {
	Code    int
	Message string
	Data    interface{}
}

func (e appError) Error() string {
	return e.Message
}

func InternalServerError() appError {
	return appError{
		Code:    50000,
		Message: "internal server error",
		Data:    nil,
	}
}

func DataNotFound() appError {
	return appError{
		Code:    40400,
		Message: "data not found",
		Data:    nil,
	}
}

func Success() appSuccess {
	return appSuccess{
		Code:    20000,
		Message: "success",
		Data:    nil,
	}
}
