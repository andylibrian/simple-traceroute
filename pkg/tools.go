package pkg

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
)

// LookupCache used to prevent AS-DNS requests for same hosts
type LookupCache struct {
	as     map[string]int64
	aMutex sync.RWMutex
	hosts  map[string]string
	hMutex sync.RWMutex
}

// NewLookupCache constructor for LookupCache
func NewLookupCache() *LookupCache {
	return &LookupCache{
		as:    make(map[string]int64, 1024),
		hosts: make(map[string]string, 4096),
	}
}

// LookupAS returns AS number for IP using origin.asn.cymru.com service
func (cache *LookupCache) LookupAS(ip string) int64 {
	cache.aMutex.RLock()
	v, exist := cache.as[ip]
	cache.aMutex.RUnlock()
	if exist {
		return v
	}

	ipParts := strings.Split(ip, ".")
	if len(ipParts) == 4 {
		return cache.lookupAS4(ip, ipParts)
	}

	return -1
}

func (cache *LookupCache) lookupAS4(ip string, ipParts []string) int64 {

	txts, err := net.LookupTXT(fmt.Sprintf("%s.%s.%s.%s.origin.asn.cymru.com", ipParts[3], ipParts[2], ipParts[1], ipParts[0]))
	if nil != err || nil == txts || len(txts) < 1 {
		return -1
	}

	parts := strings.Split(txts[0], " | ")
	if len(parts) < 2 {
		return -1
	}

	asnum, err := strconv.ParseInt(parts[0], 10, 64)
	if nil != err {
		return -1
	}

	cache.aMutex.Lock()
	cache.as[ip] = asnum
	cache.aMutex.Unlock()

	return asnum
}

// LookupHost returns AS number for IP using origin.asn.cymru.com service
func (cache *LookupCache) LookupHost(ip string) string {
	cache.hMutex.RLock()
	v, exist := cache.hosts[ip]
	cache.hMutex.RUnlock()
	if exist {
		return v
	}

	var result string

	addrs, _ := net.LookupAddr(ip)
	if len(addrs) > 0 {
		result = addrs[0]
	}

	cache.hMutex.Lock()
	cache.hosts[ip] = result
	cache.hMutex.Unlock()

	return result
}
