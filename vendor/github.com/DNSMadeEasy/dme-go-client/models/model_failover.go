package models

type FailoverAttribute struct {
	Monitor           string `json:"monitor,omitempty"`
	SystemDescription string `json:"systemDescription,omitempty"`
	MaxEmails         string `json:"maxEmails,omitempty"`
	Sensitivity       string `json:"sensitivity,omitempty"`
	ProtocolId        string `json:"protocolId,omitempty"`
	Port              string `json:"port,omitempty"`
	Failover          string `json:"failover,omitempty"`
	AutoFailover      string `json:"autoFailover,omitempty"`
	Ip1               string `json:"ip1,omitempty"`
	Ip2               string `json:"ip2,omitempty"`
	Ip3               string `json:"ip3,omitempty"`
	Ip4               string `json:"ip4,omitempty"`
	Ip5               string `json:"ip5,omitempty"`
	ContactList       string `json:"contactListId,omitempty"`
	HttpFqdn          string `json:"httpFqdn,omitempty"`
	HttpFile          string `json:"httpFile,omitempty"`
	HttpQueryString   string `json:"httpQueryString,omitempty"`
	SendString        string `json:"sendString,omitempty"`
	Timeout           string `json:"timeout,omitempty"`
	DNSFqdn           string `json:"dnsFqdn,omitempty"`
	DNSTimeout        string `json:"dnsTimeout,omitempty"`
}

func (failover *FailoverAttribute) ToMap() map[string]interface{} {
	failoverMap := make(map[string]interface{})

	A(failoverMap, "monitor", failover.Monitor)
	A(failoverMap, "systemDescription", failover.SystemDescription)
	A(failoverMap, "maxEmails", failover.MaxEmails)
	A(failoverMap, "sensitivity", failover.Sensitivity)
	A(failoverMap, "protocolId", failover.ProtocolId)
	A(failoverMap, "port", failover.Port)
	A(failoverMap, "failover", failover.Failover)
	A(failoverMap, "autoFailover", failover.AutoFailover)
	A(failoverMap, "ip1", failover.Ip1)
	A(failoverMap, "ip2", failover.Ip2)
	A(failoverMap, "ip3", failover.Ip3)
	A(failoverMap, "ip4", failover.Ip4)
	A(failoverMap, "ip5", failover.Ip5)
	A(failoverMap, "contactListId", failover.ContactList)
	A(failoverMap, "httpFqdn", failover.HttpFqdn)
	A(failoverMap, "httpFile", failover.HttpFile)
	A(failoverMap, "httpQueryString", failover.HttpQueryString)
	A(failoverMap, "sendString", failover.SendString)
	A(failoverMap, "timeout", failover.Timeout)
	A(failoverMap, "dnsFqdn", failover.DNSFqdn)
	A(failoverMap, "dnsTimeout", failover.DNSTimeout)

	return failoverMap
}
