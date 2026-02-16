package resp

import (
	"fmt"
	"strconv"
)

func SimpleString(s string) string {
	return "+" + s + "\r\n"
}

func Error(s string) string {
	return "-ERR" + s + "\r\n"
}

func Integer(n int) string {
	res := strconv.Itoa(n)
	return ":" + res + "\r\n"
}

func BulkString(s string) string {
	if s == "" {
		return "$-1\r\n"
	}
	res := fmt.Sprintf("$%d\r\n%s\r\n", len(s), s)
	return res
}

func NilBulkString() string {
	return "$-1\r\n"
}

func WrongNumberOfArgs(s string) string {
	return "-ERR wrong number of arguments for '" + s + "' command\r\n"
}

func UnknownCommand(s string) string {
	return "-ERR unknown command '" + s + "'\r\n"
}

func WrongType() string {
	return "-WRONGTYPE Operation against a key holding the wrong kind of value\r\n"
}

func ToBytes(s string) []byte {
	return []byte(s)
}

func ToString(b []byte) string {
	return string(b)
}
