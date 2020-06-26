package dme

import (
	"fmt"

	"github.com/DNSMadeEasy/dme-go-client/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceDmeCustomSoaRecord() *schema.Resource {
	return &schema.Resource{
		Read:          datasourceConstellixDomainRead,
		SchemaVersion: 1,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"email": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"comp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"ttl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"refresh": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"serial": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"retry": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"expire": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"negative_cache": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func datasourceConstellixDomainRead(d *schema.ResourceData, m interface{}) error {
	dmeClient := m.(*client.Client)
	name := d.Get("name").(string)
	con, err := dmeClient.GetbyId("dns/soa/")
	if err != nil {
		return err
	}
	data := con.S("data").Data().([]interface{})
	var flag bool
	var cnt int

	for _, info := range data {
		val := info.(map[string]interface{})
		if StripQuotes(val["name"].(string)) == name {
			flag = true
			break
		}
		cnt = cnt + 1
	}
	if flag != true {
		return fmt.Errorf("SOA Record of specified name not found")
	}

	dataCon := con.S("data").Index(cnt)
	d.SetId(dataCon.S("id").String())

	d.Set("name", StripQuotes(dataCon.S("name").String()))
	d.Set("email", StripQuotes(dataCon.S("email").String()))
	d.Set("comp", StripQuotes(dataCon.S("comp").String()))
	d.Set("ttl", StripQuotes(dataCon.S("ttl").String()))
	d.Set("retry", StripQuotes(dataCon.S("retry").String()))
	d.Set("refresh", StripQuotes(dataCon.S("refresh").String()))
	d.Set("expire", StripQuotes(dataCon.S("expire").String()))
	d.Set("serial", StripQuotes(dataCon.S("serial").String()))
	d.Set("negative_cache", StripQuotes(dataCon.S("negativeCache").String()))
	return nil

}
