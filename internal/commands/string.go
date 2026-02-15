package commands

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

func cmdSet(cmd []string, storage map[string]string, expireAt map[string]time.Time, mu *sync.RWMutex) string {
	lower := strings.ToLower(cmd[0])
	var res string

	switch len(cmd) {
	case 3:
		mu.Lock()
		storage[cmd[1]] = cmd[2]
		delete(expireAt, cmd[1])
		mu.Unlock()
		res = "+OK\r\n"

	case 5:
		res = setExPx(cmd, storage, expireAt, mu)
	default:
		res = fmt.Sprintf("-ERR wrong number of arguments for '%s' command\r\n", lower)
	}
	return res

}

func setExPx(cmd []string, storage map[string]string, expireAt map[string]time.Time, mu *sync.RWMutex) string {
	var res string
	upper := strings.ToUpper(cmd[3])
	lower := strings.ToLower(cmd[3])
	switch upper {
	case "EX":
		i, err := strconv.Atoi(cmd[4])

		if err != nil {
			res := fmt.Sprintf("-ERR wrong argument for '%s' flag\r\n", lower)
			return res
		}
		mu.Lock()
		storage[cmd[1]] = cmd[2]
		expireAt[cmd[1]] = time.Now().Add(time.Second * time.Duration(i))
		mu.Unlock()
		res = "+OK\r\n"
		return res
	case "PX":
		i, err := strconv.Atoi(cmd[4])

		if err != nil {
			res = fmt.Sprintf("-ERR wrong argument for '%s' flag\r\n", lower)
			return res
		}
		mu.Lock()
		storage[cmd[1]] = cmd[2]
		expireAt[cmd[1]] = time.Now().Add(time.Millisecond * time.Duration(i))
		mu.Unlock()
		res = "+OK\r\n"
	default:
		res = fmt.Sprintf("-ERR wrong number of arguments for '%s' command\r\n", cmd[0])
	}
	return res
}

func cmdGet(cmd []string, storage map[string]string, expeireAt map[string]time.Time, mu *sync.RWMutex) string {
	lower := strings.ToLower(cmd[0])
	var res string
	if len(cmd) != 2 {
		res = fmt.Sprintf("-ERR wrong number of arguments for '%s' command\r\n", lower)
		return res
	}
	ok := checkExpierd(cmd, storage, expeireAt, mu)
	if ok == false {
		res = "$-1\r\n"
		return res
	}
	mu.RLock()
	val := storage[cmd[1]]
	mu.RUnlock()
	res = fmt.Sprintf("$%d\r\n%s\r\n", len(val), val)
	return res
}
