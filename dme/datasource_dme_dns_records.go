package dme

import (
	"fmt"

	"github.com/DNSMadeEasy/dme-go-client/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceManagedDNSRecordActions() *schema.Resource {
	return &schema.Resource{
		Read: datasourceManagedDNSRecordActionsRead,

		Schema: map[string]*schema.Schema{
			"domain_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"dynamic_dns": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"ttl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// "monitor": &schema.Schema{
			// 	Type:     schema.TypeBool,
			// 	Optional: true,
			// 	Computed: true,
			// },

			// "failover": &schema.Schema{
			// 	Type:     schema.TypeBool,
			// 	Optional: true,
			// 	Computed: true,
			// },

			// "failed": &schema.Schema{
			// 	Type:     schema.TypeBool,
			// 	Optional: true,
			// 	Computed: true,
			// },

			"gtd_location": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"caa_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"issuer_critical": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"keywords": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"title": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"redirect_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"hardlink": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"mx_level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"weight": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"priority": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func datasourceManagedDNSRecordActionsRead(d *schema.ResourceData, m interface{}) error {
	dmeClient := m.(*client.Client)
	name := d.Get("name").(string)
	recordtype := d.Get("type").(string)

	con, err := dmeClient.GetbyId("dns/managed/" + d.Get("domain_id").(string) + "/records")
	if err != nil {
		return err
	}

	data := con.S("data").Data().([]interface{})
	var flag bool
	var count int
	for _, info := range data {
		val := info.(map[string]interface{})
		if StripQuotes(val["name"].(string)) == name && StripQuotes(val["type"].(string)) == recordtype {
			flag = true
			break
		}
		count = count + 1
	}
	if flag != true {
		return fmt.Errorf("Record of specified name not found")
	}

	cont1 := con.S("data").Index(count)

	d.SetId(fmt.Sprintf("%v", cont1.S("id").String()))
	d.Set("name", StripQuotes(cont1.S("name").String()))
	str := StripQuotes(cont1.S("value").String())
	if d.Get("type").(string) == "TXT" || d.Get("type").(string) == "SPF" || d.Get("type").(string) == "CAA" {
		str = str[2 : len(str)-2]
	}

	d.Set("value", str)

	d.Set("type", StripQuotes(cont1.S("type").String()))
	d.Set("dynamic_dns", StripQuotes(cont1.S("dynamicDns").String()))
	d.Set("password", StripQuotes(cont1.S("password").String()))
	d.Set("ttl", StripQuotes(cont1.S("ttl").String()))
	d.Set("gtd_location", StripQuotes(cont1.S("gtdLocation").String()))
	d.Set("description", StripQuotes(cont1.S("description").String()))
	d.Set("keywords", StripQuotes(cont1.S("keywords").String()))
	d.Set("title", StripQuotes(cont1.S("title").String()))
	d.Set("redirect_type", StripQuotes(cont1.S("redirectType").String()))
	d.Set("hardlink", StripQuotes(cont1.S("hardLink").String()))
	d.Set("mx_level", StripQuotes(cont1.S("mxLevel").String()))
	d.Set("weight", StripQuotes(cont1.S("weight").String()))
	d.Set("priority", StripQuotes(cont1.S("priority").String()))
	d.Set("port", StripQuotes(cont1.S("port").String()))
	d.Set("caa_type", StripQuotes(cont1.S("caaType").String()))
	d.Set("issuer_critical", StripQuotes(cont1.S("issuerCritical").String()))

	return nil

}
