---
layout: "dme"
page_title: "dme: dme_failover"
sidebar_current: "docs-dme-resource-dme_failover"
description: |-
    Manages failover in A record within the account.
---

# dme_failover #
Manages failover in A record in a domain within the account.

# Example Usage #
```hcl
resource "dme_failover" "record" {
  record_id     = "${dme_dns_record.record.id}"
  failover      = "true"
  ip1           = "1.2.3.4"
  ip2           = "1.2.3.6"
  protocol_id   = "3"
  port          = "8080"
  sensitivity   = "8"
}

```
## Argument Reference ##
* `monitor` - (Optional) True indicates System Monitoring is Enabled. If monitor is enabled, system_description and max_emails value is required.
* `system_description` - (Optional) The system description of the failover configuration.
* `max_emails` - (Optional) The maximum number of emails to send per failover event. Value can be less than or equal to 150.
* `sensitivity` - (Required) The number of checks placed against the primary IP before a Failover event occurs. List of Sensitivity ID’s:. Low (slower failover) = 8. Medium = 5. High = 3.
* `protocol_id` - (Required)The protocol for DNS Failover to monitor on. List of Protocol IDs:. TCP = 1, UDP = 2, HTTP = 3, DNS = 4, SMTP = 5, HTTPS = 6.
* `port` - (Required) The port for the DNS Failover service to monitor on the specified protocol.
* `failover` - (Required) True indicates DNS Failover is enabled. If failover minimum 2 Ip address values are required.
* `auto_failover` - (Optional) True indicates the failback to the primary IP address is a manual process. False indicates the failback to the primary IP is an automatic process.
* `ip1` - (Required) The primary IP address.
* `ip2` - (Required) The secondary  IP address.
* `ip3` - (Optional) The teriary  IP address.
* `ip4` - (Optional) The quaternary  IP address.
* `ip5` - (Optional) The quinary  IP address.
* `contact_list` - (Optional) The ID of the contact list for system monitoring notifications.
* `http_fqdn` - (Optional) The FQDN to monitor for HTTP or HTTPS checks.
* `http_file` - (Optional) The file to query for HTTP or HTTPS checks.
* `http_query_string` - (Optional) The string to query for HTTP or HTTPS checks. String length must be max 16 characters.
* `send_string` - (Optional) The send string value is required if protocol_id is 2. The value will be string.
* `timeout` - (Optional) The timeout value is required if the protocol_id is 2. The timeout value can be less than or equal to 7.
* `dns_fqdn` - (Optional) The string value is required if the protocol_id is 4.
* `dns_timeout` - (Optional) The timeout value is required if the protocol_id is 4. The timeout value can be less than or equal to 7.

## Attribute Reference ##
The only attribute that this resource exports is the `record_id`, which is set to the dme calculated id of the failover resource.










