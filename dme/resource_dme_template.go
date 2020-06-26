package dme

import (
	"log"

	"github.com/DNSMadeEasy/dme-go-client/models"

	"github.com/DNSMadeEasy/dme-go-client/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceDMETemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceDMETemplateCreate,
		Read:   resourceDMETemplateRead,
		Update: resourceDMETemplateUpdate,
		Delete: resourceDMETemplateDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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

func resourceDMETemplateCreate(d *schema.ResourceData, m interface{}) error {
	dmeClient := m.(*client.Client)
	templateAttr := models.Template{}

	if name, ok := d.GetOk("name"); ok {
		templateAttr.Name = name.(string)
	}

	// if flag, ok := d.GetOk("publicTemplate"); ok {
	// 	templateAttr.PublicTemplate = flag.(string)
	// }

	log.Println("template struct : ", templateAttr)
	con, err := dmeClient.Save(&templateAttr, "dns/template")
	if err != nil {
		return err
	}
	log.Println("container from create response :", con)
	d.SetId(con.S("id").String())
	return resourceDMETemplateRead(d, m)
}

func resourceDMETemplateUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceDMEDomainRead(d, m)
}

func resourceDMETemplateRead(d *schema.ResourceData, m interface{}) error {
	dmeClient := m.(*client.Client)
	dn := d.Id()

	con, err := dmeClient.GetbyId("dns/template/" + dn)
	if err != nil {
		return err
	}

	d.SetId(con.S("id").String())
	d.Set("name", StripQuotes(con.S("name").String()))
	d.Set("public_template", StripQuotes(con.S("publicTemplate").String()))
	ids := con.S("domainIds").Data().([]interface{})
	listIds := make([]float64, 0)
	for _, id := range ids {
		listIds = append(listIds, id.(float64))
	}

	d.Set("domain_ids", listIds)
	return nil
}

func resourceDMETemplateDelete(d *schema.ResourceData, m interface{}) error {
	dmeClient := m.(*client.Client)
	dn := d.Id()

	err := dmeClient.Delete("dns/template/" + dn)
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}
