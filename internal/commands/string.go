package commands

import (
	"strconv"
	"strings"
	"time"

	"github.com/FMR006/redis-go/internal/resp"
	"github.com/FMR006/redis-go/internal/storage"
)

func cmdSet(cmd []string, storage *storage.Storage) string {
	lower := strings.ToLower(cmd[0])
	var res string

	switch len(cmd) {
	case 3:
		storage.Set(cmd[1], cmd[2], time.Time{})
		res = resp.SimpleString("OK")

	case 5:
		res = setExPx(cmd, storage)
	default:
		res = resp.WrongNumberOfArgs(lower)
	}
	return res

}

func setExPx(cmd []string, storage *storage.Storage) string {
	var res string
	upper := strings.ToUpper(cmd[3])
	switch upper {
	case "EX":
		i, err := strconv.Atoi(cmd[4])

		if err != nil {
			res = resp.WrongType()
			return res
		}
		t := time.Now().Add(time.Second * time.Duration(i))
		storage.Set(cmd[1], cmd[2], t)
		res = resp.SimpleString("OK")
		return res
	case "PX":
		i, err := strconv.Atoi(cmd[4])

		if err != nil {
			res = resp.WrongType()
			return res
		}
		t := time.Now().Add(time.Millisecond * time.Duration(i))
		storage.Set(cmd[1], cmd[2], t)
		res = resp.SimpleString("OK")
	default:
		res = resp.WrongNumberOfArgs(cmd[0])
	}
	return res
}

func cmdGet(cmd []string, storage *storage.Storage) string {
	lower := strings.ToLower(cmd[0])
	var res string
	if len(cmd) != 2 {
		res = resp.WrongNumberOfArgs(lower)
		return res
	}
	val, flag := storage.Get(cmd[1])
	if !flag {
		res = resp.NilBulkString()
		return res
	}
	res = val
	return res
}
