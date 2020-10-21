module github.com/terraform-providers/terraform-provider-dme

go 1.13

//replace github.com/DNSMadeEasy/dme-go-client => /home/eperry/dev/tf-dnsme/dme-go-client

replace github.com/DNSMadeEasy/dme-go-client => github.com/eperry/dme-go-client v1.1.1

require (
	4d63.com/tz v1.2.0 // indirect
	github.com/DNSMadeEasy/dme-go-client v1.0.0
	github.com/hashicorp/terraform-plugin-sdk v1.14.0
	github.com/hashicorp/terraform-website v1.0.0 // indirect
)
