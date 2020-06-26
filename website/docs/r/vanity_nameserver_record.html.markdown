---
layout: "dme"
page_title: "DME: dme_vanity_nameserver_record"
sidebar_current: "docs-dme-resource-dme_vanity_nameserver_record"
description: |-
    Manages Custom Vanity Name Server Records for the account.
---
# dme_vanity_nameserver_record #
Manages Custom Vanity Name Server Records for the account.
# Example Usage #
```hcl
resource "dme_vanity_nameserver_record" "vanityrecord" {
  name                 = "newvnsrecord"
  servers              = ["abc.com.", "xyz.com.", "yash.com."]
  public_config        = false
  default_config       = false
  name_server_group_id = 1
}


```

## Argument Reference ##
* `name` - (Required) SOA Record name
* `servers` - (Required) The vanity host names defined in the configuration.
* `public_config` - (Optional) True represents a system defined rather than user defined vanity configuration. Default is false.
* `default_config` - (Optional) Indicates if the vanity configuration is the system default. Default is false.
* `name_server_group_id` - (Optional) The name server group the configuration is assigned

## Attribute Reference ##
No attributes are exported