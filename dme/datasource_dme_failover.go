package dme

import (
	"fmt"
	"log"

	"github.com/DNSMadeEasy/dme-go-client/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceDMEFailover() *schema.Resource {
	return &schema.Resource{
		Read: datasourceDMEFailoverRead,

		Schema: map[string]*schema.Schema{
			"record_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"monitor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"system_description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"max_emails": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"sensitivity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"protocol_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"failover": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"auto_failover": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"ip1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"ip2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"ip3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"ip4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"ip5": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"contact_list": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"http_fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"http_file": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"send_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"timeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"dns_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"dns_fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"http_query_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func datasourceDMEFailoverRead(d *schema.ResourceData, m interface{}) error {
	dmeClient := m.(*client.Client)
	con, err := dmeClient.GetbyId("monitor/" + d.Get("record_id").(string))
	if err != nil {
		return err
	}
	log.Println("Inside read, container value: ", con)

	d.SetId(fmt.Sprintf("%v", con.S("recordId")))
	d.Set("monitor", StripQuotes(con.S("monitor").String()))
	d.Set("system_description", StripQuotes(con.S("systemDescription").String()))
	d.Set("max_emails", StripQuotes(con.S("maxEmails").String()))
	d.Set("sensitivity", StripQuotes(con.S("sensitivity").String()))
	d.Set("protocol_id", StripQuotes(con.S("protocolId").String()))
	d.Set("port", StripQuotes(con.S("port").String()))
	d.Set("failover", StripQuotes(con.S("failover").String()))
	d.Set("auto_failover", StripQuotes(con.S("autoFailover").String()))
	d.Set("ip1", StripQuotes(con.S("ip1").String()))
	d.Set("ip2", StripQuotes(con.S("ip2").String()))
	d.Set("ip3", StripQuotes(con.S("ip3").String()))
	d.Set("ip4", StripQuotes(con.S("ip4").String()))
	d.Set("ip5", StripQuotes(con.S("ip5").String()))
	d.Set("contact_list", StripQuotes(con.S("contactListId").String()))
	d.Set("http_fqdn", StripQuotes(con.S("httpFqdn").String()))
	d.Set("http_file", StripQuotes(con.S("httpFile").String()))
	d.Set("http_query_string", StripQuotes(con.S("httpQueryString").String()))
	d.Set("send_string", StripQuotes(con.S("sendString").String()))
	d.Set("timeout", StripQuotes(con.S("timeout").String()))
	// d.Set("dns_fqdn", StripQuotes(con.S("dnsFqdn").String()))
	// d.Set("dns_timeout", StripQuotes(con.S("dnsTimeout").String()))

	return nil

}
