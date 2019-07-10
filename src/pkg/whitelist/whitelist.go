package whitelist

import (
	"regexp"
	"sync"
)

var (
	lock sync.RWMutex
	ips  = map[string]*regexp.Regexp{}
)

func Setup(list []string) error {
	lock.Lock()
	defer lock.Unlock()

	for _, ip := range list {
		re, err := regexp.Compile(ip)
		if err != nil {
			return err
		}
		ips[ip] = re
	}
	return nil
}

//VerifyIP check the ip is a legal or not
func VerifyIP(ip string) bool {
	lock.RLock()
	defer lock.RUnlock()
	for _, r := range ips {
		if r.MatchString(ip) {
			return true
		}
	}
	return false
}
