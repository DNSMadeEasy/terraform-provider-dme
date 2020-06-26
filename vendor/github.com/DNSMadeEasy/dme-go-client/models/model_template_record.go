package models

import (
	"log"
)

type TemplateRecord struct {
	Name           string `json:"name,omitempty"`
	IdUpdate       string `json:"id,omitempty"`
	Value          string `json:"value,omitempty"`
	Type           string `json:"type,omitempty"`
	DynamicDNS     string `json:"dynamicDns,omitempty"`
	Password       string `json:"password,omitempty"`
	Ttl            string `json:"ttl,omitempty"`
	GtdLocation    string `json:"gtdLocation,omitempty"`
	Description    string `json:"description,omitempty"`
	Keywords       string `json:"keywords,omitempty"`
	Title          string `json:"title,omitempty"`
	RedirectType   string `json:"redirectType,omitempty"`
	HardLink       string `json:"hardLink,omitempty"`
	MxLevel        string `json:"mxLevel,omitempty"`
	Weight         string `json:"weight,omitempty"`
	Priority       string `json:"priority,omitempty"`
	Port           string `json:"port,omitempty"`
	CaaType        string `json:"caaType,omitempty"`
	IssuerCritical string `json:"issuerCritical,omitempty"`
}

func (record *TemplateRecord) ToMap() map[string]interface{} {
	recordMap := make(map[string]interface{})

	log.Println("Inside model: recordmap values: ", recordMap)
	log.Println("RECORD VAlues inside model: ", record)
	A(recordMap, "name", record.Name)
	A(recordMap, "id", record.IdUpdate)
	A(recordMap, "value", record.Value)
	A(recordMap, "type", record.Type)
	A(recordMap, "dynamicDns", record.DynamicDNS)
	A(recordMap, "password", record.Password)
	A(recordMap, "ttl", record.Ttl)
	A(recordMap, "gtdLocation", record.GtdLocation)
	A(recordMap, "description", record.Description)
	A(recordMap, "keywords", record.Keywords)
	A(recordMap, "title", record.Title)
	A(recordMap, "redirectType", record.RedirectType)
	A(recordMap, "hardLink", record.HardLink)
	A(recordMap, "mxLevel", record.MxLevel)
	A(recordMap, "weight", record.Weight)
	A(recordMap, "priority", record.Priority)
	A(recordMap, "port", record.Port)
	A(recordMap, "caaType", record.CaaType)
	A(recordMap, "issuerCritical", record.IssuerCritical)

	log.Println("Inside model:after recordmap values: ", recordMap)

	return recordMap
}
