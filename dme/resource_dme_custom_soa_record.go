package dme

import (
	"fmt"
	"log"

	"github.com/DNSMadeEasy/dme-go-client/client"
	"github.com/DNSMadeEasy/dme-go-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceDmeCustomSoaRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceDmeSOACreate,
		Read:   resourceDmeSOARead,
		Update: resourceDmeSOAUpdate,
		Delete: resourceDmeSOADelete,

		SchemaVersion: 1,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"email": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"comp": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},

			"refresh": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},

			"serial": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},

			"retry": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},

			"expire": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},

			"negative_cache": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func resourceDmeSOACreate(d *schema.ResourceData, m interface{}) error {
	dmeConnect := m.(*client.Client)

	soaAttr := &models.Soa{}

	if value, ok := d.GetOk("name"); ok {
		soaAttr.Name = value.(string)
	}

	if value, ok := d.GetOk("email"); ok {
		soaAttr.Email = value.(string)
	}

	if value, ok := d.GetOk("comp"); ok {
		soaAttr.Comp = value.(string)
	}

	if value, ok := d.GetOk("ttl"); ok {
		soaAttr.TTL = value.(int)
	}

	if value, ok := d.GetOk("serial"); ok {
		soaAttr.Serial = value.(int)
	}

	if value, ok := d.GetOk("refresh"); ok {
		soaAttr.Refresh = value.(int)
	}

	if value, ok := d.GetOk("expire"); ok {
		soaAttr.Expire = value.(int)
	}

	if value, ok := d.GetOk("retry"); ok {
		soaAttr.Retry = value.(int)
	}

	if value, ok := d.GetOk("negative_cache"); ok {
		soaAttr.NegativeCache = value.(int)
	}

	cont, err := dmeConnect.Save(soaAttr, "dns/soa")

	if err != nil {
		log.Println("Error returned: ", err)
		return err
	}

	log.Println("Value of container: ", cont)
	id := cont.S("id")
	log.Println("Id value: ", id)
	d.SetId(fmt.Sprintf("%v", id))
	return resourceDmeSOARead(d, m)
}

func resourceDmeSOARead(d *schema.ResourceData, m interface{}) error {
	dmeClient := m.(*client.Client)
	dn := d.Id()

	con, err := dmeClient.GetbyId("dns/soa/" + dn)
	if err != nil {
		return err
	}

	d.Set("name", StripQuotes(con.S("name").String()))
	d.Set("email", StripQuotes(con.S("email").String()))
	d.Set("comp", StripQuotes(con.S("comp").String()))
	d.Set("ttl", StripQuotes(con.S("ttl").String()))
	d.Set("retry", StripQuotes(con.S("retry").String()))
	d.Set("refresh", StripQuotes(con.S("refresh").String()))
	d.Set("expire", StripQuotes(con.S("expire").String()))
	d.Set("serial", StripQuotes(con.S("serial").String()))
	d.Set("negative_cache", StripQuotes(con.S("negativeCache").String()))
	return nil
}

func resourceDmeSOAUpdate(d *schema.ResourceData, m interface{}) error {
	dmeConnect := m.(*client.Client)

	soaAttr := &models.Soa{}

	if value, ok := d.GetOk("name"); ok {
		soaAttr.Name = value.(string)
	}

	if value, ok := d.GetOk("email"); ok {
		soaAttr.Email = value.(string)
	}

	if value, ok := d.GetOk("comp"); ok {
		soaAttr.Comp = value.(string)
	}

	if value, ok := d.GetOk("ttl"); ok {
		soaAttr.TTL = value.(int)
	}

	if value, ok := d.GetOk("serial"); ok {
		soaAttr.Serial = value.(int)
	}

	if value, ok := d.GetOk("refresh"); ok {
		soaAttr.Refresh = value.(int)
	}

	if value, ok := d.GetOk("expire"); ok {
		soaAttr.Expire = value.(int)
	}

	if value, ok := d.GetOk("retry"); ok {
		soaAttr.Retry = value.(int)
	}

	if value, ok := d.GetOk("negative_cache"); ok {
		soaAttr.NegativeCache = value.(int)
	}
	log.Println("SOA structure is :", soaAttr)
	dn := d.Id()

	_, err := dmeConnect.Update(soaAttr, "dns/soa/"+dn)
	if err != nil {
		return err
	}
	return resourceDmeSOARead(d, m)
}

func resourceDmeSOADelete(d *schema.ResourceData, m interface{}) error {
	dmeClient := m.(*client.Client)
	dn := d.Id()

	err := dmeClient.Delete("dns/soa/" + dn)
	if err != nil {
		return nil
	}

	d.SetId("")
	return nil
}
