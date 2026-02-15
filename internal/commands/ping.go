package commands

import (
	"fmt"
	"strings"
)

func cmdPing(cmd []string) string {
	lower := strings.ToLower(cmd[0])
	var res string
	if len(cmd) == 1 {
		res = "+PONG\r\n"
	} else if len(cmd) == 2 {
		msg := cmd[1]
		res = fmt.Sprintf("$%d\r\n%s\r\n", len(msg), msg)
	} else {
		res = fmt.Sprintf("-ERR wrong number of arguments for '%s' command\r\n", lower)
	}
	return res
}
