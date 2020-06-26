package models

type Folder struct {
	Name              string        `json:"name,omitempty"`
	DefaultFolder     bool          `json:"defaultFolder,omitempty"`
	Domains           []string      `json:"domains,omitempty"`
	Secondaries       []string      `json:"secondaries,omitempty"`
	FolderPermissions []interface{} `json:"folderPermissions,omitempty"`
}

func (folder *Folder) ToMap() map[string]interface{} {
	folderMap := make(map[string]interface{})

	A(folderMap, "name", folder.Name)
	A(folderMap, "defaultFolder", folder.DefaultFolder)
	A(folderMap, "domains", folder.Domains)
	A(folderMap, "secondaries", folder.Secondaries)
	A(folderMap, "folderPermissions", folder.FolderPermissions)

	return folderMap
}
