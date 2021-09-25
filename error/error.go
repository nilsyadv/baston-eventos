package error

type CustomeError struct {
	Msg          string `json:"message"`
	Err          string `json:"error"`
	ResponseCode int    `json:"response_code"`
}

func (customerr CustomeError) Error() string {
	return customerr.Err
}

func (customerr CustomeError) Message() string {
	return customerr.Msg
}

// CreateCustomeError Create New Custome Error
func CreateCustomeError(msg string, err error, respcode int) CustomeError {
	return CustomeError{
		Msg:          msg,
		Err:          err.Error(),
		ResponseCode: respcode,
	}
}
