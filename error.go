package errorwrapper

import "fmt"

type causer interface {
	Cause() error
}

type BasicError struct {
	msg   string
	cause error

	code string
}

var DefaultSeparator string = "\n"

func New(msg string) error { return &BasicError{msg: msg} }

func NewWithError(err error) error { return &BasicError{msg: err.Error(), cause: err} }

func NewWithCode(msg string, code string) error { return &BasicError{msg: msg, code: code} }

func WrapMessage(err error, msg string) error {
	if err == nil {
		return nil
	}
	if e, ok := err.(interface{ Code() string }); ok {
		return &BasicError{msg: msg, cause: err, code: e.Code()}
	}
	return &BasicError{msg: msg, cause: err}
}

func WrapMessageWithCode(err error, msg string, code string) error {
	if err == nil {
		return nil
	}
	return &BasicError{msg: msg, cause: err, code: code}
}

func (e *BasicError) String() string {
	if e == nil {
		return ""
	}
	if e.cause == nil {
		return e.msg
	}
	if cause, ok := e.cause.(fmt.Stringer); ok {
		return e.msg + DefaultSeparator + cause.String()
	}
	return e.msg + DefaultSeparator + e.cause.Error()
}

func (e *BasicError) Error() string { return e.msg }

func (e *BasicError) Cause() error { return e.cause }

func (e *BasicError) Code() string { return e.code }

func Cause(err error) error {
	if causer, ok := err.(causer); ok && causer.Cause() != nil {
		return Cause(causer.Cause())
	}
	return err
}

func UnWrap(err error) error {
	if causer, ok := err.(causer); ok {
		return causer.Cause()
	}
	return err
}

func Code(err error) string {
	if coder, ok := err.(interface{ Code() string }); ok {
		return coder.Code()
	}
	return ""
}

func String(err error) string {
	if stringer, ok := err.(fmt.Stringer); ok {
		return stringer.String()
	}
	return err.Error()
}
