## for development: add to your .terraformrc:
# provider_installation {
#   dev_overrides {
#       "dnsmadeeasy.com/provider/dme" = "<PATH>" # usually $HOME/go/bin
#   }
#   direct {}
# }


terraform {
  required_providers {
    dme = {
      source = "dnsmadeeasy.com/provider/dme"
    }
  }
}

variable "api_key" {
  description = "API key for DME provider"
  type        = string
}

variable "secret_key" {
  description = "Secret key for DME provider"
  type        = string
}

## add a `secrets.auto.tfvars` file with your api secrets. 
provider "dme" {
  api_key    = var.api_key
  secret_key = var.secret_key
  ## if you want to use the sandbox api, see docs at at https://api-docs.dnsmadeeasy.com/#intro 
  # base_url = "https://api.sandbox.dnsmadeeasy.com/V2.0" 
}

resource "dme_domain" "domain1" {
  name = "domain1.com"
}

resource "dme_dns_record" "record1" {
  domain_id     = dme_domain.domain1.id
  name          = "record1"
  type          = "A"
  ttl           = "86400"
  value         = "10.20.30.40"
  # description   = "First http record"
  # keywords      = "practice record"
  # title         = "record"
  # redirect_type = "Standard - 302"
  # hardlink      = "true"
}

resource "dme_dns_record" "record2" {
  domain_id     = dme_domain.domain1.id
  name          = "record2"
  type          = "A"
  ttl           = "86400"
  value         = "10.20.30.50"
  # description   = "First http record"
  # keywords      = "practice record"
  # title         = "record"
  # redirect_type = "Standard - 302"
  # hardlink      = "true"
}

resource "dme_dns_record" "HTTPREDrecord" {
  domain_id     = "${dme_domain.domain1.id}"
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

output "domain1" {
  value = "${dme_domain.domain1}"
}
# output "record1" {
#   value = dme_dns_record.record1
# }

