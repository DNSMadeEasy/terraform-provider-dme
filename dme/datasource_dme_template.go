package dme

import (
	"fmt"

	"github.com/DNSMadeEasy/dme-go-client/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceDMETemplate() *schema.Resource {
	return &schema.Resource{

		Read: datasourceDMETemplateRead,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"domain_ids": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeFloat},
			},

			"public_template": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func datasourceDMETemplateRead(d *schema.ResourceData, m interface{}) error {
	dmeClient := m.(*client.Client)
	name := d.Get("name").(string)

	con, err := dmeClient.GetbyId("dns/template")
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
		return fmt.Errorf("Template of specified name not found")
	}

	dataCon := con.S("data").Index(cnt)
	d.SetId(dataCon.S("id").String())
	d.Set("name", StripQuotes(dataCon.S("name").String()))
	d.Set("public_template", StripQuotes(dataCon.S("publicTemplate").String()))

	ids := dataCon.S("domainIds").Data().([]interface{})
	listIds := make([]float64, 0)
	for _, id := range ids {
		listIds = append(listIds, id.(float64))
	}

	d.Set("domain_ids", listIds)
	return nil
}
