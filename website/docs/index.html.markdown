---
layout: "dme"
page_title: "Provider: DNSMadeEasy"
sidebar_current: "docs-dme-index"
description: |-
  The DNS made easy provider is used to manage various DNS Objects supported by DNS Made Easy platform. The provider needs to be configured with the proper credentials before it can be used.
---
DME Provider
------------
DME is a leading DNS service provider with a feature rich DNS services which includes, various kinds of dns records such as Aname record, Cname record, HTTPredirection, MX record. The DME provider is used to manage various DNS Objects supported by DNS Made Easy platform. The provider needs to be configured with the proper credentials before it can be used.

Authentication
--------------
The Provider supports authentication with DME platform using API-key and SECRET-key. 

 1. Authentication with user-id and password.  
 example:  

----------
 ```hcl
provider "dme" {
  # dme Api key
  api_key    = "apikey"
  # dme secret key
  secret_key = "secretkey"
  insecure  = true
  proxyurl = "https://proxy_server:proxy_port"
}
 ```

Example Usage
------------
```hcl
provider "dme" {
  # dme Api key
  api_key    = "apikey"
  # dme secret key
  secret_key = "secretkey"
  insecure  = true
  proxyurl = "https://proxy_server:proxy_port"
}

resource "dme_domain" "domain1" {
  name = "domain1.com"
}
```

Argument Reference
------------------
Following arguments are supported with DNS Made Easy terraform provider.

 * `api_key` - (Required) API key of a user which has the access to perform CRUD operations on all the DNS objects of DNS Made Easy platform.
 * `secret_key` - (Required) Secret key of a user which has the access to perform CRUD operations on all the DNS objects of DNS Made Easy platform.
 * `insecure` - (Optional) This determines whether to use insecure HTTP connection or not. Default value is `true`.  
 * `proxyurl` - (Optional) A proxy server URL when configured, all the requests to DNS Made Easy platform will be passed through the proxy-server configured.