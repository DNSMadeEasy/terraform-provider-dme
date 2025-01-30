package dme

import (
	"fmt"

	"github.com/DNSMadeEasy/dme-go-client/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "API key for HTTP call",
				DefaultFunc: schema.EnvDefaultFunc("DME_API_KEY", nil),
			},

			"secret_key": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Secret Key for HMAC",
				DefaultFunc: schema.EnvDefaultFunc("DME_SECRET_KEY", nil),
			},

			"insecure": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Allows insecure HTTTPS client",
			},

			"proxyurl": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Proxy server URL",
				Deprecated:  "Use proxy_url instead",
			},

			"proxy_url": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Proxy server URL",
			},

			"base_url": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "API base URL",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"dme_custom_soa_record":        resourceDmeCustomSoaRecord(),
			"dme_domain":                   resourceDMEDomain(),
			"dme_dns_record":               resourceManagedDNSRecordActions(),
			"dme_template":                 resourceDMETemplate(),
			"dme_vanity_nameserver_record": resourceDmeVanityNameserverRecord(),
			"dme_transfer_acl":             resourceDMEACL(),
			"dme_secondary_dns":            resourceDMESecondaryDNS(),
			"dme_secondary_ip_set":         resourceDMESecondaryIPSet(),
			"dme_failover":                 resourceDMEFailover(),
			"dme_folder_record":            resourceDmeFolder(),
			"dme_template_record":          resourceDMETemplateRecord(),
			"dme_contact_list":             resourceDMEContactList(),
		},

		DataSourcesMap: map[string]*schema.Resource{
			"dme_custom_soa_record":        datasourceDmeCustomSoaRecord(),
			"dme_domain":                   datasourceDMEDomain(),
			"dme_dns_record":               datasourceManagedDNSRecordActions(),
			"dme_template":                 datasourceDMETemplate(),
			"dme_vanity_nameserver_record": datasourceDmeVanityNameserverRecord(),
			"dme_transfer_acl":             datasourceDmeACL(),
			"dme_secondary_dns":            datasourceDMESecondaryDNS(),
			"dme_secondary_ip_set":         datasourceDMESecondaryIPSet(),
			"dme_failover":                 datasourceDMEFailover(),
			"dme_folder_record":            datasourceDmeFolder(),
			"dme_template_record":          datasourceDMETemplateRecord(),
			"dme_contact_list":             datasourceDMEContactList(),
		},

		ConfigureFunc: configureClient,
	}
}

func configureClient(d *schema.ResourceData) (interface{}, error) {
	config := config{
		apiKey:    d.Get("api_key").(string),
		secretKey: d.Get("secret_key").(string),
		insecure:  d.Get("insecure").(bool),
		proxyURL:  getProxyURL(d),
		baseURL:   d.Get("base_url").(string),
	}

	if err := config.Valid(); err != nil {
		return nil, err
	}
	cli := config.getClient()
	return cli, nil
}

func getProxyURL(d *schema.ResourceData) string {
	// First try with `proxy_url`.
	if v, ok := d.Get("proxy_url").(string); ok {
		return v
	}
	// Otherwise, try deprecated `proxyurl`
	return d.Get("proxyurl").(string)
}

func (c config) Valid() error {
	if c.apiKey == "" {
		return fmt.Errorf("API Key is required")
	}

	if c.secretKey == "" {
		return fmt.Errorf("secret key is required")
	}
	return nil
}

func (c config) getClient() interface{} {
	return client.GetClient(
		c.apiKey,
		c.secretKey,
		client.Insecure(c.insecure),
		client.ProxyURL(c.proxyURL),
		client.BaseURL(c.baseURL),
	)
}

type config struct {
	apiKey    string
	secretKey string
	insecure  bool
	proxyURL  string
	baseURL   string
}
