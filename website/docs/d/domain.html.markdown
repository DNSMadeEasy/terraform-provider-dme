---
layout: "dme"
page_title: "dme: dme_domain"
sidebar_current: "docs-dme-datasource-dme_domain"
description: |-
  Data source for domain action
---

# dme_domain #
Data source for Domain action

## Example Usage ##

```hcl
data "dme_domain" "example" {
  name        = "example.com"
}

```

## Argument Reference ##
* `name` - (Required) Name of domain action. Name should be unique.

## Attribute Reference ##
* `name` - (Required) Name of domain action. Name should be unique.
* `gtd_enabled` - (Optional) Indicator of whether or not this domain uses the Global Traffic Director service.
* `soa_id` - (Optional) The ID of a custom SOA record.
* `template_id` - (Optional) The ID of a template applied to the domain.
* `vanity_id` - (Optional) The ID of a vanity DNS configuration.
* `transfer_acl_id` - (Optional) The ID of an applied transfer ACL.
* `folder_id` - (Optional) The ID of a domain folder.
* `updated` - (Optional) The number of seconds since the domain
was last updated in Epoch time.
* `created` - (Optional) The number of seconds since the domain
was last created in Epoch time.