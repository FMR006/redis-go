package resp

import (
	"bufio"
	"errors"
	"strconv"
)

func ReaderLine(r *bufio.Reader) (string, error) {
	var buf []byte
	for {
		b, err := r.ReadByte()
		if err != nil {
			return "", err
		}
		switch b {
		case '\r':
			next, err := r.ReadByte()
			if err != nil {
				return "", err
			}
			if next != '\n' {
				return "", errors.New("protocol error: expected \\n after \\r")
			}
			return string(buf), nil
		default:
			buf = append(buf, b)
		}
	}
}

func ReadExactly(r *bufio.Reader, n int) (string, error) {
	var buf []byte
	for i := 0; i < n; i++ {
		b, err := r.ReadByte()
		if err != nil {
			return "", err
		}
		buf = append(buf, b)
	}
	return string(buf), nil
}

func ReadBulkString(r *bufio.Reader) (string, error) {
	str, err := ReaderLine(r)

	if err != nil {
		return "error with reader line", err
	}
	if str[0] != '$' {
		return "", errors.New("invalid type of appeal")
	}

	size, err4 := strconv.Atoi(str[1:])
	if err4 != nil {
		return "error with atoi", err4
	}

	res, err1 := ReadExactly(r, size)

	if err1 != nil {
		return "error with read exactly", err1
	}
	check, err2 := r.ReadByte()
	if err2 != nil {
		return "error with read byte 1", err2
	}
	if check != '\r' {
		return "", errors.New("invalid format")
	}
	next, err3 := r.ReadByte()

	if err3 != nil {
		return "error with read byte 2", err3
	}

	if next != '\n' {
		return "", errors.New("invalid format with next")
	}
	return res, nil
}

func ReadArray(r *bufio.Reader) ([]string, error) {
	res := make([]string, 0)
	line, err := ReaderLine(r)
	if err != nil {
		return []string{"error with ReaderLine"}, err
	}
	if line[0] != '*' {
		return []string{}, errors.New("invalid format of appeal")
	}

	count, err1 := strconv.Atoi(line[1:])
	if err1 != nil {
		return []string{"error with atoi"}, err1
	}

	for i := 0; i < count; i++ {
		bulk, err := ReadBulkString(r)
		if err != nil {
			return []string{"error with ReadBulkString"}, err
		}
		res = append(res, bulk)
	}
	return res, nil
}
