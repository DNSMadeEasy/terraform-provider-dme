---
layout: "dme"
page_title: "dme: dme_secondary_dns"
sidebar_current: "docs-dme-resource-dme_secondary_dns"
description: |-
    Manages secondary DNS within the account.
---

# dme_secondary_dns #
Manages one or more secondary DNS within the account.

# Example Usage #
```hcl
resource "dme_secondary_dns" "example" {
  name = "example.com"
  ipset_id = "${dme_secondary_ip_set.example.id}"
  folder_id = "${dme_folder.example.id}"
}

```

# Note #
It takes around 10 minutes to reflect the changes on the DNS Made Easy platform. Till time your record will be marked in creating state and won't be deleted.

## Argument Reference ##
* `name` - (Required) Name of Secondary DNS action. Name should be unique.
* `ipset_id` - (Required) The ID of secondary ip set.
* `folder_id` - (Optional) The ID of a domain folder.

## Attribute Reference ##
* `id` - Set to the dme calculated id of secondary DNS action.