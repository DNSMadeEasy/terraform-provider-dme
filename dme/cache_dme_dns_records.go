package dme

import (
	"sync"

	"github.com/DNSMadeEasy/dme-go-client/client"
	"github.com/DNSMadeEasy/dme-go-client/container"
)

type dnsRecordsCacheType struct {
	records map[string]*container.Container
	mutex   sync.Mutex
}

var dnsRecordsCache = dnsRecordsCacheType{
	make(map[string]*container.Container),
	sync.Mutex{},
}

func (cache *dnsRecordsCacheType) getByDomain(dmeClient *client.Client, domain string) (*container.Container, error) {
	cache.mutex.Lock()
	if _, ok := cache.records[domain]; !ok {
		con, err := dmeClient.GetbyId("dns/managed/" + domain + "/records")
		if err != nil {
			cache.mutex.Unlock()
			return nil, err
		}
		cache.records[domain] = con.S("data")
	}
	cache.mutex.Unlock()
	return cache.records[domain], nil
}
