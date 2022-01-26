package transfer

import constants "clean-architect/constant"

// base transfer
type Base struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// error transfer

type CustomError struct {
	Base
}

func (e *CustomError) Error() string {
	return e.Code + ": " + e.Message
}

func NewCustomError(code string) *CustomError {
	return &CustomError{
		Base{
			Code: code,
		},
	}
}

func NewCustomErrorMsg(code string, msg string) *CustomError {
	return &CustomError{
		Base{
			Code:    code,
			Message: msg,
		},
	}
}

//success transfer

type Success struct {
	Base
	Data interface{} `json:"data"`
}

func NewSuccess(data interface{}) *Success {
	sucess := Success{}
	sucess.Code = constants.SUCCESS
	sucess.Data = data
	return &sucess
}
