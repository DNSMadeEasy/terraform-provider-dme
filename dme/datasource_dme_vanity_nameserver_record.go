package dme

import (
	"fmt"

	"github.com/DNSMadeEasy/dme-go-client/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceDmeVanityNameserverRecord() *schema.Resource {
	return &schema.Resource{
		Read:          datasourceDmeVanityNameserverRead,
		SchemaVersion: 1,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"servers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"public_config": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"default_config": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"name_server_group_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func datasourceDmeVanityNameserverRead(d *schema.ResourceData, m interface{}) error {
	dmeClient := m.(*client.Client)
	name := d.Get("name").(string)

	con, err := dmeClient.GetbyId("dns/vanity/")
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
	// d.Set("servers", StripQuotes(dataCon.S("servers").String()))
	d.Set("public_config", StripQuotes(dataCon.S("public").String()))
	d.Set("default_config", StripQuotes(dataCon.S("default").String()))
	d.Set("name_server_group_id", StripQuotes(dataCon.S("nameServerGroupId").String()))

	servers := dataCon.S("servers").Data().([]interface{})
	listServers := make([]string, 0)
	for _, server := range servers {
		listServers = append(listServers, server.(string))
	}
	d.Set("servers", listServers)

	return nil
}
