package service

import "errors"

var (
	// An error that is thrown when the response body is empty https://github.com/go-swagger/go-swagger/issues/1929
	ServiceEmptyBodyErrMessage = "&{0 [] } (*models.RPCStatus) is not supported by the TextConsumer, can be resolved by supporting TextUnmarshaler interface"
	ErrEmptyResponseBody       = errors.New("the request response body was empty, which likely means the request was unauthorized. check your authentication credentials and try again")
)

func handleRequestError(err error) error {
	switch err.Error() {
	case ServiceEmptyBodyErrMessage:
		return ErrEmptyResponseBody
	default:
		return err
	}
}
