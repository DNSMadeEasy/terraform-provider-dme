package models

import "log"

type ACLAttribute struct {
	Name string   `json:"name,omitempty"`
	Ips  []string `json:"ips,omitempty"`
}

func (acl *ACLAttribute) ToMap() map[string]interface{} {
	aclMap := make(map[string]interface{})
	A(aclMap, "name", acl.Name)
	A(aclMap, "ips", acl.Ips)
	log.Println("inside model: B value: ", aclMap)
	return aclMap
}
