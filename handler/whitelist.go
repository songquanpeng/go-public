package handler

import (
	"go-public/common"
	"sync"
)

var ipStore = struct {
	sync.RWMutex
	ips map[string]bool
}{
	ips: make(map[string]bool),
}

var whitelistCacheThreshold = 4

func isInWhitelist(ip string) bool {
	// If whitelist is empty, just return true
	if len(common.ServerConfig.Whitelist) == 0 {
		return true
	}
	// If ip is localhost, just return true
	if ip == "::1" {
		return true
	}
	// If whitelist is too short, we don't cache it
	if len(common.ServerConfig.Whitelist) < whitelistCacheThreshold {
		for _, v := range common.ServerConfig.Whitelist {
			if v == ip {
				return true
			}
		}
		return false
	}
	// Check the cache
	ipStore.RLock()
	if _, ok := ipStore.ips[ip]; ok {
		ipStore.RUnlock()
		return true
	}
	ipStore.RUnlock()
	// Check the whitelist
	for _, v := range common.ServerConfig.Whitelist {
		if v == ip {
			ipStore.Lock()
			ipStore.ips[ip] = true
			ipStore.Unlock()
			return true
		}
	}
	return false
}
