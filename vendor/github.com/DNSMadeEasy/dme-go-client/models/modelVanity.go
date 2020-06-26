package models

type Vanity struct {
	Name              string   `json:"name,omitempty"`
	Servers           []string `json:"servers,omitempty"`
	Public            bool     `json:"public,omitempty"`
	Default           bool     `json:"default,omitempty"`
	NameServerGroupID int      `json:"nameServerGroupId,omitempty"`
}

func (vanity *Vanity) ToMap() map[string]interface{} {
	vanityMap := make(map[string]interface{})
	A(vanityMap, "name", vanity.Name)
	A(vanityMap, "servers", vanity.Servers)
	A(vanityMap, "public", vanity.Public)
	A(vanityMap, "default", vanity.Default)
	A(vanityMap, "nameServerGroupId", vanity.NameServerGroupID)

	return vanityMap
}
