package common

type ICodeError interface {
	ErrorMessage() string
}

type codeErrorImpl struct {
	Message string
}

func NewCodeError(message string) ICodeError {
	return &codeErrorImpl{
		Message: message,
	}
}

func (c *codeErrorImpl) ErrorMessage() string {
	return c.Message
}
