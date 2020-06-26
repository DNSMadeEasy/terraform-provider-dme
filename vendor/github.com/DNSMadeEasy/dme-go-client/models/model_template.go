package models

type Template struct {
	Name           string        `json:"name,omitempty"`
	DomainID       []interface{} `json:"domainIds,omitempty"`
	PublicTemplate string        `json:"publicTemplate,omitempty"`
}

func (tem *Template) ToMap() map[string]interface{} {
	templatMap := make(map[string]interface{})

	A(templatMap, "name", tem.Name)

	A(templatMap, "publicTemplate", tem.PublicTemplate)

	return templatMap
}
