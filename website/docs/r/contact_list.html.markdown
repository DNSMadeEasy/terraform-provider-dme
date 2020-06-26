---
layout: "dme"
page_title: "dme: dme_contact_list"
sidebar_current: "docs-dme-resource-dme_contact_list"
description: |-
    Manages contact list within the account.
---

# dme_contact_list #
Manages one or more contact list within the account.

# Example Usage #
```hcl
resource "dme_contact_list" "first" {
  name = "ashu01"
  emails = ["abc@gmail.com", "cde@gmail.com"]
}
```

## Argument Reference ##
* `name` - (Required) Name of contact list action. Name should be unique.
* `emails` - (Required) List of emails to associate with the contact list. 

## Attribute Reference ##
* `id` - Set to the dme calculated id of contact list.