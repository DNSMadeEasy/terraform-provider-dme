---
layout: "dme"
page_title: "dme: dme_secondary_ip_set"
sidebar_current: "docs-dme-datasource-dme_secondary_ip_set"
description: |-
  Data source for Secondary IP Set action
---

# dme_secondary_ip_set #
Data source for Secondary IP Set action

## Example Usage ##

```hcl
data "dme_secondary_ip_set" "example" {
  name = "example"
}

```

## Argument Reference ##
* `name` - (Required) Name of secondary ip set action. Name should be unique.

## Attribute Reference ##
* `name` - (Required) Name of secondary ip set action. Name should be unique.
* `ips` - List of ip addresses assigned in the secondary ip set.
* `id` - Set to the dme calculated id of secondary Ip set action.