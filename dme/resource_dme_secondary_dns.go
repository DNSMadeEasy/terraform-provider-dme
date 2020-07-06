package dme

import (
	"log"

	"github.com/DNSMadeEasy/dme-go-client/client"
	"github.com/DNSMadeEasy/dme-go-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceDMESecondaryDNS() *schema.Resource {
	return &schema.Resource{
		Create: resourceDMESecondaryDNSCreate,
		Read:   resourceDMESecondaryDNSRead,
		Update: resourceDMESecondaryDNSUpdate,
		Delete: resourceDMESecondaryDNSDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"ipset_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"folder_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceDMESecondaryDNSCreate(d *schema.ResourceData, m interface{}) error {
	dmeClient := m.(*client.Client)
	dnsAttr := &models.SecondaryDNS{}

	if name, ok := d.GetOk("name"); ok {
		nameList := toListOfInterface(name)
		dnsAttr.Name = nameList
	}

	if ips, ok := d.GetOk("ipset_id"); ok {
		dnsAttr.IpsetID = ips.(string)
	}

	if folder, ok := d.GetOk("folder_id"); ok {
		dnsAttr.FolderID = folder.(string)
	}

	log.Println("struct for dns :", dnsAttr)
	con, err := dmeClient.Save(dnsAttr, "dns/secondary")
	if err != nil {
		return err
	}
	log.Println("container in the response of create :", con)
	d.SetId(con.S("id").String())
	return resourceDMESecondaryDNSRead(d, m)
}

func resourceDMESecondaryDNSUpdate(d *schema.ResourceData, m interface{}) error {
	dmeClient := m.(*client.Client)
	dnsAttr := &models.SecondaryDNS{}

	if ips, ok := d.GetOk("ipset_id"); ok {
		dnsAttr.IpsetID = ips.(string)
	}

	if d.HasChange("folder_id") {
		dnsAttr.FolderID = d.Get("folder_id").(string)
	}

	dn := d.Id()
	idList := toListOfInterface(dn)
	dnsAttr.Ids = idList
	_, err := dmeClient.Update(dnsAttr, "dns/secondary/")
	if err != nil {
		return err
	}
	return resourceDMESecondaryDNSRead(d, m)
}

func resourceDMESecondaryDNSRead(d *schema.ResourceData, m interface{}) error {
	dmeClient := m.(*client.Client)
	dn := d.Id()

	con, err := dmeClient.GetbyId("dns/secondary/" + dn)
	if err != nil {
		return err
	}

	d.SetId(con.S("id").String())
	d.Set("ipset_id", con.S("ipSetId").String())
	d.Set("folder_id", con.S("folderId").String())
	return nil
}

func resourceDMESecondaryDNSDelete(d *schema.ResourceData, m interface{}) error {
	dmeClient := m.(*client.Client)
	dn := d.Id()

	err := dmeClient.Delete("dns/secondary/" + dn)
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}
