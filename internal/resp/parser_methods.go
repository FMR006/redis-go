package main

import (
	"bufio"
	"errors"
	"strconv"
)

func readerLine(r *bufio.Reader) (string, error) {
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

func readExactly(r *bufio.Reader, n int) (string, error) {
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

func readBulkString(r *bufio.Reader) (string, error) {
	str, err := readerLine(r)

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

	res, err1 := readExactly(r, size)

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

func readArray(r *bufio.Reader) ([]string, error) {
	res := make([]string, 0)
	line, err := readerLine(r)
	if err != nil {
		return []string{"error with readerLine"}, err
	}
	if line[0] != '*' {
		return []string{}, errors.New("invalid format of appeal")
	}

	count, err1 := strconv.Atoi(line[1:])
	if err1 != nil {
		return []string{"error with atoi"}, err1
	}

	for i := 0; i < count; i++ {
		bulk, err := readBulkString(r)
		if err != nil {
			return []string{"eror with readBulkString"}, err
		}
		res = append(res, bulk)
	}
	return res, nil
}
