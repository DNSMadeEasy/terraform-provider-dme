package models

type SecondaryIPSet struct {
	Name string        `json:"name,omitempty"`
	IPs  []interface{} `json:"ips,omitempty"`
}

func (secondIp *SecondaryIPSet) ToMap() map[string]interface{} {
	secondIPMap := make(map[string]interface{})

	A(secondIPMap, "name", secondIp.Name)

	A(secondIPMap, "ips", secondIp.IPs)

	return secondIPMap
}
