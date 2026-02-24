package commands

import (
	"strings"

	"github.com/FMR006/redis-go/internal/resp"
	"github.com/FMR006/redis-go/internal/storage"
)

func Dispatch(cmd []string, storage *storage.Storage) string {
	var res string
	if len(cmd) == 0 {
		return resp.Error("empty command")
	}
	upper := strings.ToUpper(cmd[0])
	switch upper {
	case "PING":
		res = cmdPing(cmd)
	case "ECHO":
		res = cmdEcho(cmd)
	case "SET":
		res = cmdSet(cmd, storage)
	case "GET":
		res = cmdGet(cmd, storage)
	case "RPUSH":
		res = cmdRPush(cmd, storage)
	case "LRANGE":
		res = cmdLRange(cmd, storage)
	case "LPUSH":
		res = cmdLPush(cmd, storage)
	case "LLEN":
		res = cmdLLen(cmd, storage)
	default:
		res = resp.UnknownCommand(cmd[0])
	}
	return res

}
