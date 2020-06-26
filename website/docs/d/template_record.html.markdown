---
layout: "dme"
page_title: "dme: dme_template_record"
sidebar_current: "docs-dme-datasource-dme_template_record"
description: |-
    Manages records in a template within the account.
---

# dme_template_record #
Manages one or more records in a template within the account.

# Example Usage #
```hcl
data "dme_template_record" "record" {
  template_id   = "${dme_template.example.id}"
  name          = "practice"
  type          = "A"
}

```


## Argument Reference ##
* `name` - (Required) Name of record.
* `type` - (Required) The record type. Values: A, AAAA, ANAME, CNAME, HTTPRED, MX, NS, PTR, SRV, TXT, CAA or SPF.
* `template_id` - (Required) Template id of the record.

## Attribute Reference ##
* `value` - (Optional) Value of record.
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
* `ttl` - (Optional) The time to live or TTL of the record.
* `gtd_location` - (Optional) Global Traffic Director location. Values:DEFAULT, US_EAST, US_WEST, EUROPE, ASIA_PAC, OCREANIA.
* `dynamic_dns` - (Optional) Indicates if the record has dynamic DNS enabled or not.
* `password` - (Optional) The per record password for a dynamic DNS update.
* `description` - (Optional) For HTTP Redirection Records, A description of the HTTP Redirection Record
* `keywords` - (Optional) For HTTPRED records. Keywords associated with the HTTPRED record.
* `title` - (Optional) For HTTPRED records. The title of the HTTPRED record.
* `redirect_type` - (Optional) For HTTPRED records. Type of redirection performed. Values: Hidden Frame Masked, Standard - 301, Standard - 302.
* `hardlink` - (Optional) For HTTP Redirection Records.
* `mx_level` - (Optional) The priority for an MX record.
* `weight` - (Optional) The weight for an SRV Record.
* `priority` - (Optional) The priority for an SRV Record.
* `port` - (Optional) The port for an SRV Record. 
* `caa_type` - (Optional) The type for an CAA Record. Caa type can be "issue", "issuewild", "iodef"
* `issuer_critical` - (Optional) The issuer critical value for an CAA Record. Value will be integer less than or equals to 255.
