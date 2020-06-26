---
layout: "dme"
page_title: "DME: dme_custom_soa_record"
sidebar_current: "docs-dme-datasource-dme_custom_soa_record"
description: |-
    Manages Custom SOA Records for the account.
---
# dme_custom_soa_record #
Manages Custom SOA Records for the account.

# Example Usage #
```hcl
data "dme_custom_soa_record" "soacheck" {
  name = "soarecord"
}

```

## Argument Reference ##
* `name` - (Required) SOA Record name

## Attribute Reference ##
* `name` - (Optional) SOA Record name
* `email` - (Optional) Contact email address.
* `comp` - (Optional) Primary name server. 
* `ttl` - (Optional) TTL of the SOA Record (in seconds). TTl value should be greater than or equal to 21600
* `refresh` - (Optional) Zone refresh time (in seconds). Refresh value should be greater than or equal to 14400
* `serial` - (Optional) Starting zone serial number
* `retry` - (Optional) Failed Refresh retry time (in seconds). Retry value should be greater than or equal to 300
* `expire` - (Optional) Expire time of zone the (in seconds). Expire value should be greater than or equal to 86400
* `negative_cache` - (Optional) Record not found cache (in seconds). Negative Cache value should be greater than or equal to 180