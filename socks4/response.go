package socks4

import "fmt"

const (
	Granted = iota + 90
	RejectedOrFailed
	NoIdentd
	UserIdMismatch
)

const responseSize = 8

var responseCodeNames map[byte]string = map[byte]string{
	Granted:          "Granted",
	RejectedOrFailed: "RejectedOrFailed",
	NoIdentd:         "NoIdentd",
	UserIdMismatch:   "UserIdMismatch",
}

type Response struct {
	version      byte
	responseCode byte
	reserved     [6]byte
}

func (r Response) ResponseCode() byte {
	return r.responseCode
}

func (r Response) SetResponseCode(code byte) {
	r.responseCode = code
}

func (r Response) Bytes() []byte {
	return append([]byte{r.version, r.responseCode}, r.reserved[:]...)
}

func (r Response) String() string {
	return fmt.Sprintf("Response code : %s", responseCodeNames[r.responseCode])
}

type invalidResponseError struct {
	message string
}

func (e *invalidResponseError) Error() string {
	return e.message
}

func makeInvalidResultCodeError() *invalidResponseError {
	return &invalidResponseError{message: "Invalid result code"}
}

func NewResponse(resultCode byte) (*Response, error) {
	if ok, err := validateResultCode(resultCode); !ok {
		return nil, err
	}
	return &Response{responseCode: resultCode}, nil
}

func validateResultCode(resultCode byte) (bool, error) {
	if resultCode != Granted && resultCode != RejectedOrFailed && resultCode != NoIdentd && resultCode != UserIdMismatch {
		return false, makeInvalidResultCodeError()
	}
	return true, nil
}
