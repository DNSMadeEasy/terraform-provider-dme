---
layout: "dme"
page_title: "dme: dme_dns_record"
sidebar_current: "docs-dme-resource-dme_dns_record"
description: |-
    Manages records in a domain within the account.
---

# dme_dns_record #
Manages one or more records in a domain within the account.

# Example Usage #
```hcl
resource "dme_dns_record" "record" {
  domain_id     = "${dme_domain.example.id}"
  name          = "practice"
  type          = "HTTPRED"
  ttl           = "86402"
  value         = "http://www.facebook.com"
  description   = "First http record"
  keywords      = "practice record"
  title         = "record"
  redirect_type = "Standard - 302"
  hardlink      = "true"
}

```

## Argument Reference ##
* `name` - (Required) Name of record.
* `value` - (Required) Value of record.
  For A record Ipv4 address is required. For example: value: "1.2.3.4"
  For CNAME record alias name is required. For example: value: "www"
  For ANAME record FQDN is required. For example: value: "www.google.com."
  For MX record server name is required. For example: value: "document."
  For HTTPRED record URL is required. For example: value: "http://www.google.com"
  For TXT record text data is required. For example: value: "practice"
  For SPF record string value is required. For example: value: "1.2.3.4"
  For PTR record host name is required. For example: value: "mail.domainDocument."
  For NS record host name is required. For example: value: "mail.domainDocument." 
  For AAAA record IPv6 address is required. For example: value: "0::0:0:0:0:0:6"
  For SRV record host name is required. For example: value: "mail.domainDocument."
  For CAA record text data is required. For example: value: "comodoca.com"
* `type` - (Required) The record type. Values: A, AAAA, ANAME, CNAME, HTTPRED, MX, NS, PTR, SRV, TXT, CAA or SPF.
* `ttl` - (Required) The time to live or TTL of the record.
* `gtd_location` - (Optional) Global Traffic Director location. Values:DEFAULT, US_EAST, US_WEST, EUROPE, ASIA_PAC, OCREANIA.
* `dynamic_dns` - (Optional) Indicates if the record has dynamic DNS enabled or not.
* `password` - (Optional) The per record password for a dynamic DNS update.
* `description` - (Optional) For HTTP Redirection Records, A description of the HTTP Redirection Record
* `keywords` - (Optional) For HTTPRED records. Keywords associated with the HTTPRED record.
* `title` - (Optional) For HTTPRED records. The title of the HTTPRED record.
* `redirect_type` - (Optional) For HTTPRED records. Type of redirection performed. Values: Hidden Frame Masked, Standard - 301, Standard - 302.
* `hardlink` - (Optional) For HTTP Redirection Records.
* `mx_level` - (Optional) The priority for an MX record. MxLevel is required for creating mx record.
* `weight` - (Optional) The weight for an SRV Record. Weight is required for creating SRV record.
* `priority` - (Optional) The priority for an SRV Record. Priority is required for creating SRV record.
* `port` - (Optional) The port for an SRV Record. Port is required for creating SRV record.
* `caa_type` - (Optional) The type for an CAA Record. Caatype is required for creating CAA record. Caa type can be "issue", "issuewild", "iodef"
* `issuer_critical` - (Optional) The issuer critical value for an CAA Record. It is required for creating CAA record. Value will be integer less than or equla to 255.

## Attribute Reference ##
The only attribute that this resource exports is the `domain_id`, which is set to the dme calculated id of the resource.