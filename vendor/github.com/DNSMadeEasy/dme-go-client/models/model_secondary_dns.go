package models

type SecondaryDNS struct {
	Name     []interface{} `json:"names,omitempty"`
	IpsetID  string        `json:"ipSetId,omitempty"`
	FolderID string        `json:"folderId,omitempty"`
	Ids      []interface{} `json:",omitempty"`
}

func (secondDNS SecondaryDNS) ToMap() map[string]interface{} {
	secondaryDNSMap := make(map[string]interface{})

	if secondDNS.Name != nil {
		A(secondaryDNSMap, "names", secondDNS.Name)
	}

	A(secondaryDNSMap, "ipSetId", secondDNS.IpsetID)

	A(secondaryDNSMap, "folderId", secondDNS.FolderID)

	A(secondaryDNSMap, "ids", secondDNS.Ids)

	return secondaryDNSMap
}
