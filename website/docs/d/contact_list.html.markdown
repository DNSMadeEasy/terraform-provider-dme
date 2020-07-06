---
layout: "dme"
page_title: "dme: dme_contact_list"
sidebar_current: "docs-dme-datasource-dme_contact_list"
description: |-
  Data source for Contact List
---

# dme_contact_list #
Data source for Contact List

## Example Usage ##

```hcl
data "dme_contact_list" "example" {
  name = "example"
}

```

## Argument Reference ##
* `name` - (Required) Name of Contact List action. Name should be unique.

## Attribute Reference ##
* `name` - (Required) Name of contact list action. Name should be unique.
* `emails` - List of emails assigned in the contact list.
* `id` - Set to the dme calculated id of Contact list action.