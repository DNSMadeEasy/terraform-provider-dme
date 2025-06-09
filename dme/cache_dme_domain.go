package dme

import (
	"sync"

	"github.com/DNSMadeEasy/dme-go-client/client"
	"github.com/DNSMadeEasy/dme-go-client/container"
)

type domainCacheType struct {
	domains map[string]*container.Container
	mutex   sync.Mutex
}

var domainCache = domainCacheType{
	make(map[string]*container.Container),
	sync.Mutex{},
}

func (cache *domainCacheType) getByDomain(dmeClient *client.Client, domain string) (*container.Container, error) {
	cache.mutex.Lock()
	if _, ok := cache.domains[domain]; !ok {
		con, err := dmeClient.GetbyId("dns/managed/" + domain)
		if err != nil {
			cache.mutex.Unlock()
			return nil, err
		}
		cache.domains[domain] = con
	}
	cache.mutex.Unlock()
	return cache.domains[domain], nil
}
