---
layout: "dme"
page_title: "dme: dme_failover"
sidebar_current: "docs-dme-datasource-dme_failover"
description: |-
    Manages failover in A record within the account.
---

# dme_failover #
Manages failover in A record in a domain within the account.

# Example Usage #
```hcl
data "dme_failover" "first" {
	record_id	= "${dme_dns_record.first.id}"
}


```


## Argument Reference ##
* `record_id` - (Required) Record id of the record.

## Attribute Reference ##
* `monitor` - (Optional) True indicates System Monitoring is Enabled. If monitor is enabled, system_description and max_emails value is required.
* `system_description` - (Optional) The system description of the failover configuration.
* `max_emails` - (Optional) The maximum number of emails to send per failover event.Value can be less than or equal to 150.
* `sensitivity` - (Optional) The number of checks placed against the primary IP before a Failover event occurs. List of Sensitivity ID’s:. Low (slower failover) = 8. Medium = 5. High = 3.
* `protocol_id` - (Optional)The protocol for DNS Failover to monitor on. List of Protocol IDs:. TCP = 1, UDP = 2, HTTP = 3, DNS = 4, SMTP = 5, HTTPS = 6.
* `port` - (Optional) The port for the DNS Failover service to monitor on the specified protocol.
* `failover` - (Optional) True indicates DNS Failover is enabled. If failover minimum 2 Ip address values are required.
* `auto_failover` - (Optional) True indicates the failback to the primary IP address is a manual process. False indicates the failback to the primary IP is an automatic process.
* `ip1` - (Optional) The primary IP address.
* `ip2` - (Optional) The secondary  IP address.
* `ip3` - (Optional) The teriary  IP address.
* `ip4` - (Optional) The quaternary  IP address.
* `ip5` - (Optional) The quinary  IP address.
* `contact_list` - (Optional) The ID of the contact list for system monitoring notifications.
* `http_fqdn` - (Optional) The FQDN to monitor for HTTP or HTTPS checks.
* `http_file` - (Optional) The file to query for HTTP or HTTPS checks.
* `http_query_string` - (Optional) The string to query for HTTP or HTTPS checks.
* `send_string` - (Optional) The send string value for protocol_id = 2. The value will be string.
* `timeout` - (Optional) The timeout value for the protocol_id = 2. The timeout value can be less than or equal to 7.
* `dns_fqdn` - (Optional) The string value for the protocol_id = 4.
* `dns_timeout` - (Optional) The timeout value for the protocol_id = 4. The timeout value can be less than 100.











