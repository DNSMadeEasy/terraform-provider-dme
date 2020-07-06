---
layout: "dme"
page_title: "DME: dme_custom_soa_record"
sidebar_current: "docs-dme-resource-dme_custom_soa_record"
description: |-
    Manages Custom SOA Records for the account.
---
# dme_custom_soa_record #
Manages Custom SOA Records for the account.
# Example Usage #
```hcl
resource "dme_custom_soa_record" "soacheck" {
  name           = "soarecord"
  email          = "crestsoa.com."
  comp           = "crestsoa.com."
  ttl            = 23000
  negative_cache = 400
  refresh        = 14400
  retry          = 300
  serial         = 2009010110
  expire         = 86450
}

```

## Argument Reference ##
* `name` - (Required) SOA Record name
* `email` - (Required) Contact email address.
* `comp` - (Required) Primary name server. 
* `ttl` - (Required) TTL of the SOA Record (in seconds). TTl value should be greater than or equal to 21600
* `refresh` - (Required) Zone refresh time (in seconds). Refresh value should be greater than or equal to 14400
* `serial` - (Required) Starting zone serial number
* `retry` - (Required) Failed Refresh retry time (in seconds). Retry value should be greater than or equal to 300
* `expire` - (Required) Expire time of zone the (in seconds). Expire value should be greater than or equal to 86400
* `negative_cache` - (Required) Record not found cache (in seconds). Negative Cache value should be greater than or equal to 180

## Attribute Reference ##
No attributes are exported