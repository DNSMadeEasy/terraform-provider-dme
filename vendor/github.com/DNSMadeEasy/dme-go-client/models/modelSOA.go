package models

type Soa struct {
	Name          string `json:"name,omitempty"`
	Email         string `json:"email,omitempty"`
	Comp          string `json:"comp,omitempty"`
	TTL           int    `json:"ttl,omitempty"`
	Serial        int    `json:"serial,omitempty"`
	Refresh       int    `json:"refresh,omitempty"`
	Retry         int    `json:"retry,omitempty"`
	Expire        int    `json:"expire,omitempty"`
	NegativeCache int    `json:"negativeCache,omitempty"`
}

func (soa *Soa) ToMap() map[string]interface{} {
	soaMap := make(map[string]interface{})

	A(soaMap, "name", soa.Name)
	A(soaMap, "email", soa.Email)
	A(soaMap, "comp", soa.Comp)
	A(soaMap, "ttl", soa.TTL)
	A(soaMap, "serial", soa.Serial)
	A(soaMap, "refresh", soa.Refresh)
	A(soaMap, "retry", soa.Retry)
	A(soaMap, "expire", soa.Expire)
	A(soaMap, "negativeCache", soa.NegativeCache)

	return soaMap
}
