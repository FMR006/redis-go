package commands

import (
	"strings"

	"github.com/FMR006/redis-go/internal/resp"
)

func cmdEcho(cmd []string) string {
	lower := strings.ToLower(cmd[0])
	var res string
	if len(cmd) != 2 {
		res = resp.WrongNumberOfArgs(lower)
		return res
	}
	res = resp.BulkString(cmd[1])
	return res
}
