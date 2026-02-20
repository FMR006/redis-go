package commands

import (
	"strconv"
	"strings"
	"time"

	"github.com/FMR006/redis-go/internal/resp"
	"github.com/FMR006/redis-go/internal/storage"
)

func cmdRPush(cmd []string, storage *storage.Storage) string {
	lower := strings.ToLower(cmd[0])
	if len(cmd) != 3 {
		if len(cmd) < 3 {
			return resp.WrongNumberOfArgs(lower)
		}

		for i := 2; i < len(cmd)-1; i++ {
			key := cmd[1]
			value := cmd[i]
			_, ok := storage.RPush(key, value, time.Time{})
			if !ok {
				return resp.WrongType()
			}
		}
		key := cmd[1]
		line := len(cmd)
		i, ok := storage.RPush(key, cmd[line-1], time.Time{})
		if !ok {
			return resp.WrongType()
		}
		return resp.Integer(i)
	}
	key := cmd[1]
	value := cmd[2]

	i, ok := storage.RPush(key, value, time.Time{})
	if !ok {
		return resp.WrongType()
	}

	return resp.Integer(i)
}

func cmdLRange(cmd []string, storage *storage.Storage) string {
	lower := strings.ToLower(cmd[0])
	if len(cmd) != 4 {
		return resp.WrongNumberOfArgs(lower)
	}
	key := cmd[1]
	start, err := strconv.Atoi(cmd[2])
	if err != nil {
		return resp.WrongType()
	}
	stop, err := strconv.Atoi(cmd[3])
	if err != nil {
		return resp.WrongType()
	}

	values, ok := storage.LRange(key, start, stop)
	if !ok {
		return resp.WrongType()
	}

	return resp.Array(values)
}
