---
layout: "dme"
page_title: "dme: dme_secondary_ip_set"
sidebar_current: "docs-dme-resource-dme_secondary_ip_set"
description: |-
    Manages Secondary IP Set within the account.
---

# dme_secondary_ip_set #
Manages one or more Secondary IP Set within the account.

# Example Usage #
```hcl
resource "dme_secondary_ip_set" "one" {
  name = "example"
  ips = [
    "12.35.4.8"
  ]
}
```

## Argument Reference ##
* `name` - (Required) Name of secondary ip set action. Name should be unique.
* `ips` - (Required) List of ip addresses. 

## Attribute Reference ##
* `id` - Set to the dme calculated id of secondary ip set action.