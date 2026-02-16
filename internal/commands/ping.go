package commands

import (
	"strings"

	"github.com/FMR006/redis-go/internal/resp"
)

func cmdPing(cmd []string) string {
	lower := strings.ToLower(cmd[0])
	var res string
	if len(cmd) == 1 {
		res = resp.BulkString("PONG")
	} else if len(cmd) == 2 {
		msg := cmd[1]
		res = resp.BulkString(msg)
	} else {
		res = resp.WrongNumberOfArgs(lower)
	}
	return res
}
