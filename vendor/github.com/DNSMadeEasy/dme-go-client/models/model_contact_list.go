package models

type ContactList struct {
	Name   string        `json:"name,omitempty"`
	Emails []interface{} `json:"emails,omitempty"`
}

func (contactlist *ContactList) ToMap() map[string]interface{} {
	contactlistMap := make(map[string]interface{})

	A(contactlistMap, "name", contactlist.Name)

	A(contactlistMap, "emails", contactlist.Emails)

	return contactlistMap
}
