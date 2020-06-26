---
layout: "dme"
page_title: "dme: dme_template"
sidebar_current: "docs-dme-datasource-dme_template"
description: |-
  Data source for template action
---

# dme_template #
Data source for Template action

## Example Usage ##

```hcl
data "dme_template" "example" {
  name = "example"
}

```

## Argument Reference ##
* `name` - (Required) Name of domain action. Name should be unique.

## Attribute Reference ##
* `name` - (Required) Name of domain action. Name should be unique.
* `domain_ids` - ids of the domain to which this template is associated.
* `public_template` - True represents a system defined public template rather than a user defined account specific template.
* `id` - Set to the dme calculated id of domain action.