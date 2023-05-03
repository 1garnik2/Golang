package main

import "fmt"

const (
	OkMsg        = "OK"
	CancelledMsg = "CANCELLED"
)

// error codes
const (
	Ok = iota
	Cancelled
	Unknown
)

// ErrorMessageToCode returns a gRPC error code depending on error message.
func ErrorMessageToCode(msg string) int {
	switch msg {
	case OkMsg:
		return Ok
	case CancelledMsg:
		return Cancelled
	}

	return Unknown
}
func main() {
	fmt.Println(ErrorMessageToCode("OK"))
	fmt.Println(ErrorMessageToCode("CANCELLED"))
	fmt.Println(ErrorMessageToCode("UNKNOWN"))
}
