package dme

import (
	"log"
	"reflect"
	"sort"

	"github.com/DNSMadeEasy/dme-go-client/client"
	"github.com/DNSMadeEasy/dme-go-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceDMESecondaryIPSet() *schema.Resource {
	return &schema.Resource{
		Create: resourceDMESecondaryIPSetCreate,
		Read:   resourceDMESecondaryIPSetRead,
		Update: resourceDMESecondaryIPSetUpdate,
		Delete: resourceDMESecondaryIPSetDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"ips": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceDMESecondaryIPSetCreate(d *schema.ResourceData, m interface{}) error {
	dmeClient := m.(*client.Client)
	secondIP := &models.SecondaryIPSet{}

	if name, ok := d.GetOk("name"); ok {
		secondIP.Name = name.(string)
	}

	if ips, ok := d.GetOk("ips"); ok {
		secondIP.IPs = ips.([]interface{})
	}

	log.Println("model struct in create SIP :", secondIP)
	con, err := dmeClient.Save(secondIP, "dns/secondary/ipSet")
	if err != nil {
		return err
	}

	log.Println("container from response in create :", con)
	d.SetId(con.S("id").String())
	return resourceDMESecondaryIPSetRead(d, m)
}

func resourceDMESecondaryIPSetUpdate(d *schema.ResourceData, m interface{}) error {
	dmeClient := m.(*client.Client)

	secondIP := &models.SecondaryIPSet{}
	if name, ok := d.GetOk("name"); ok {
		secondIP.Name = name.(string)
	}

	if d.HasChange("ips") {
		secondIP.IPs = d.Get("ips").([]interface{})
	}

	log.Println("model struct in create SIP :", secondIP)
	dn := d.Id()
	_, err := dmeClient.Update(secondIP, "dns/secondary/ipSet/"+dn)
	if err != nil {
		return err
	}
	return resourceDMESecondaryIPSetRead(d, m)
}

func resourceDMESecondaryIPSetRead(d *schema.ResourceData, m interface{}) error {
	dmeClient := m.(*client.Client)
	dn := d.Id()

	con, err := dmeClient.GetbyId("dns/secondary/ipSet/" + dn)
	if err != nil {
		return err
	}

	d.SetId(con.S("id").String())
	d.Set("name", StripQuotes(con.S("name").String()))
	ips := con.S("ips").Data().([]interface{})
	respIPList := make([]string, 0)
	for _, ip := range ips {
		respIPList = append(respIPList, ip.(string))
	}

	listget := make([]string, 0)
	if ips, ok := d.GetOk("ips"); ok {
		listget = toListOfString(ips)
	}

	finallist := make([]string, 0)
	finallist = append(finallist, listget...)

	sort.Strings(listget)
	sort.Strings(respIPList)

	if reflect.DeepEqual(listget, respIPList) {
		d.Set("ips", d.Get("ips"))
		return nil
	}
	d.Set("ips", con.S("ips").Data().([]interface{}))
	return nil
}

func resourceDMESecondaryIPSetDelete(d *schema.ResourceData, m interface{}) error {
	dmeClient := m.(*client.Client)
	dn := d.Id()

	err := dmeClient.Delete("dns/secondary/ipSet/" + dn)
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}
