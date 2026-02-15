package commands

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

func dispatch(cmd []string, storage map[string]string, expireAt map[string]time.Time, mu *sync.RWMutex) string {
	var res string
	if len(cmd) == 0 {
		return "-ERR empty command\r\n"
	}
	upper := strings.ToUpper(cmd[0])
	switch upper {
	case "PING":
		res = cmdPing(cmd)
	case "ECHO":
		res = cmdEcho(cmd)
	case "SET":
		res = cmdSet(cmd, storage, expireAt, mu)
	case "GET":
		res = cmdGet(cmd, storage, expireAt, mu)

	default:
		res = fmt.Sprintf("-ERR unknown command '%s'\r\n", cmd[0])
	}
	return res

}
