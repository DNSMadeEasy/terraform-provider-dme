---
layout: "dme"
page_title: "dme: dme_template"
sidebar_current: "docs-dme-resource-dme_template"
description: |-
    Manages templates within the account.
---

# dme_template #
Manages one or more templates within the account.

# Example Usage #
```hcl
resource "dme_template" "first" {
  name = "ashutosh"
}

```

## Argument Reference ##
* `name` - (Required) Name of template action. Name should be unique.

## Attribute Reference ##
* `domain_ids` - ids of the domain to which this template is associated.
* `public_template` - True represents a system defined public template rather than a user defined account specific template.
* `id` - Set to the dme calculated id of domain action.