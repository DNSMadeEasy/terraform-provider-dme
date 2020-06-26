---
layout: "dme"
page_title: "dme: dme_transfer_acl"
sidebar_current: "docs-dme-resource-dme_transfer_acl"
description: |-
    Manages ACL (Access Control List) within the account.
---

# dme_transfer_acl #
Manages one or more Access Control Lists within the account.

# Example Usage #
```hcl
resource "dme_transfer_acl" "example" {
  name 	= "transferacl"
  ips	= ["1.2.3.4", "2.3.4.5"]
}

```

## Argument Reference ##
* `name` - (Required) ACL Identifiable name.
* `ips` - (Required) The list of IP addresses defined in the ACL.

## Attribute Reference ##
No attributed is exported for creating this resource. 

