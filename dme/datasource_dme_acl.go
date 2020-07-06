package dme

import (
	"fmt"
	"log"

	"github.com/DNSMadeEasy/dme-go-client/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceDmeACL() *schema.Resource {
	return &schema.Resource{
		Read:          datasourceDMEACLRead,
		SchemaVersion: 1,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"ips": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func datasourceDMEACLRead(d *schema.ResourceData, m interface{}) error {
	dmeClient := m.(*client.Client)
	name := d.Get("name").(string)
	con, err := dmeClient.GetbyId("dns/transferAcl/")
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
		return fmt.Errorf("ACL of specified name not found")
	}
	dataCon := con.S("data").Index(cnt)

	log.Println("container value:", dataCon)
	d.SetId(dataCon.S("id").String())

	d.Set("name", StripQuotes(dataCon.S("name").String()))

	ips := dataCon.S("ips").Data().([]interface{})
	listips := make([]string, 0)

	for _, id := range ips {
		listips = append(listips, id.(string))
	}
	d.Set("ips", listips)

	return nil
}
