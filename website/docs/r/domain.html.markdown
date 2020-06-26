---
layout: "dme"
page_title: "dme: dme_domain"
sidebar_current: "docs-dme-resource-dme_domain"
description: |-
    Manages domains within the account.
---

# dme_domain #
Manages one or more domains within the account.

# Example Usage #
```hcl
resource "dme_domain" "example" {
  name = "example.com"
  gtd_enabled = "false"
  soa_id = "${dme_custom_soa_record.example.id}"
  template_id = "${dme_template.example.id}"
  vanity_id = "${dme_vanity_nameserver_record.example.id}"
  transfer_acl_id = "${dme_transfer_acl.example.id}"
  folder_id = "${dme_folder.example.id}"
}

```

# Note #
It takes around 10 minutes to reflect the changes on the DNS Made Easy platform. Till time your record will be marked in creating state and won't be deleted.

## Argument Reference ##
* `name` - (Required) Name of domain action. Name should be unique.
* `gtd_enabled` - (Optional) Indicator of whether or not this domain uses the Global Traffic Director service.
* `soa_id` - (Optional) The ID of a custom SOA record.
* `template_id` - (Optional) The ID of a template applied to the domain.
* `vanity_id` - (Optional) The ID of a vanity DNS configuration.
* `transfer_acl_id` - (Optional) The ID of an applied transfer ACL.
* `folder_id` - (Optional) The ID of a domain folder.

## Attribute Reference ##
* `updated` - The number of seconds since the domain
was last updated in Epoch time. Not configurable by the user.
* `created` - The number of seconds since the domain
was last created in Epoch time. Not configurable by the user.
* `id` - Set to the dme calculated id of domain action.