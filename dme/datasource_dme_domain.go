package dme

import (
	"fmt"

	"github.com/DNSMadeEasy/dme-go-client/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceDMEDomain() *schema.Resource {
	return &schema.Resource{
		Read: datasourceDMEDomainRead,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"gtd_enabled": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"soa_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"template_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"vanity_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"transfer_acl_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"folder_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"updated": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"created": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func datasourceDMEDomainRead(d *schema.ResourceData, m interface{}) error {
	dmeClient := m.(*client.Client)
	name := d.Get("name").(string)

	con, err := dmeClient.GetbyId("dns/managed/")
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
		return fmt.Errorf("Domain of specified name not found")
	}

	dataCon := con.S("data").Index(cnt)
	d.SetId(dataCon.S("id").String())
	d.Set("name", StripQuotes(dataCon.S("name").String()))
	d.Set("gtd_enabled", StripQuotes(dataCon.S("gtdEnabled").String()))
	d.Set("soa_id", StripQuotes(dataCon.S("soaId").String()))
	d.Set("template_id", StripQuotes(dataCon.S("templateId").String()))
	d.Set("vanity_id", StripQuotes(dataCon.S("vanityId").String()))
	d.Set("transfer_acl_id", StripQuotes(dataCon.S("transferAclId").String()))
	d.Set("folder_id", StripQuotes(dataCon.S("folderId").String()))
	d.Set("updated", StripQuotes(dataCon.S("updated").String()))
	d.Set("created", StripQuotes(dataCon.S("created").String()))
	return nil
}
