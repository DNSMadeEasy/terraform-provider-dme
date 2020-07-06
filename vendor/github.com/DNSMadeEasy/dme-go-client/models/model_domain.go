package models

type DomainAttribute struct {
	Name          string `json:"name,omitempty"`
	GtdEnabled    string `json:"gtdEnabled,omitempty"`
	SOAID         string `json:"soaId,omitempty"`
	TemplateID    string `json:"templateId,omitempty"`
	VanityID      string `json:"vanityId,omitempty"`
	TransferAClID string `json:"transferAclId,omitempty"`
	FolderID      string `json:"folderId,omitempty"`
	Updated       string `json:"updated,omitempty"`
	Created       string `json:"created,omitempty"`
}

func (domain *DomainAttribute) ToMap() map[string]interface{} {
	domainMap := make(map[string]interface{})

	A(domainMap, "name", domain.Name)

	A(domainMap, "gtdEnabled", domain.GtdEnabled)

	A(domainMap, "soaId", domain.SOAID)

	A(domainMap, "templateId", domain.TemplateID)

	A(domainMap, "vanityId", domain.VanityID)

	A(domainMap, "transferAclId", domain.TransferAClID)

	A(domainMap, "folderId", domain.FolderID)

	A(domainMap, "updated", domain.Updated)

	A(domainMap, "created", domain.Created)

	return domainMap
}
