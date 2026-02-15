package commands

import (
	"fmt"
	"strings"
)

func cmdEcho(cmd []string) string {
	lower := strings.ToLower(cmd[0])
	var res string
	if len(cmd) != 2 {
		res = fmt.Sprintf("-ERR wrong number of arguments for '%s' command\r\n", lower)
		return res
	}
	res = fmt.Sprintf("$%d\r\n%s\r\n", len(cmd[1]), cmd[1])
	return res
}
