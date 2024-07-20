package socks4

import (
	"encoding/binary"
	"fmt"
	"net"
)

const (
	Connect byte = iota + 1
	Bind
)

var commandNames map[byte]string = map[byte]string{Connect: "Connect", Bind: "Bind"}

const ProtocolVersion byte = 4

const minRequestLength = 9

type Request struct {
	version            byte
	command            byte
	destinationAddress *net.TCPAddr
	userID             string
}

func (r Request) Version() byte {
	return r.version
}

func (r Request) Command() byte {
	return r.command
}

func (r Request) DestinationAddress() *net.TCPAddr {
	return r.destinationAddress
}

func (r Request) UserID() string {
	return r.userID
}

func (r Request) String() string {
	return fmt.Sprintf("Version : %d Command %s Address : %s UserID %s",
		r.version, commandNames[r.command], r.destinationAddress.String(), r.userID)
}

type invalidRequestError struct {
	message string
}

func (error *invalidRequestError) Error() string {
	return error.message
}

func makeInvalidLengthError() *invalidRequestError {
	return &invalidRequestError{message: "Invalid request length."}
}

func makeInvalidVersionError() *invalidRequestError {
	return &invalidRequestError{message: "Invalid SOCKS version."}
}

func makeInvalidCommandError() *invalidRequestError {
	return &invalidRequestError{message: "Invalid SOCKS command."}
}

func NewRequest(raw []byte) (*Request, error) {
	if ok, err := validateRequest(raw); !ok {
		return nil, err
	}
	dest := makeTCPAddr(raw)
	var userID string = ""
	if len(raw) > minRequestLength {
		userID = string(raw[9:len(raw)])
	}
	return &Request{version: raw[0], command: raw[1], destinationAddress: dest,
		userID: userID}, nil
}
func validateRequest(raw []byte) (bool, error) {
	if len(raw) < minRequestLength {
		return false, makeInvalidLengthError()
	}
	if raw[0] != ProtocolVersion {
		return false, makeInvalidVersionError()
	}
	if (raw[1] != Connect) && (raw[1] != Bind) {
		return false, makeInvalidCommandError()
	}
	return true, nil
}

func makeTCPAddr(raw []byte) *net.TCPAddr {
	return &net.TCPAddr{IP: net.IP{raw[4], raw[5], raw[6], raw[7]},
		Port: int(binary.BigEndian.Uint16(raw[2:4])),
		Zone: "",
	}
}
