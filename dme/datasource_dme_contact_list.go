package dme

import (
	"fmt"
	"log"

	"github.com/DNSMadeEasy/dme-go-client/client"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceDMEContactList() *schema.Resource {
	return &schema.Resource{
		Read: datasourceDMEContactListRead,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"emails": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func datasourceDMEContactListRead(d *schema.ResourceData, m interface{}) error {
	dmeClient := m.(*client.Client)
	name := d.Get("name").(string)

	con, err := dmeClient.GetbyId("contactList/getAll")
	if err != nil {
		return err
	}

	temp := con.Data().([]interface{})
	var cnt int
	var flag bool
	for _, val := range temp {
		valMap := val.(map[string]interface{})
		if valMap["name"] == name {
			flag = true
			break
		}
		cnt = cnt + 1
	}
	if flag != true {
		return fmt.Errorf("Contact list of specified name not found")
	}

	log.Println("container index: ", cnt)
	dataCon := con.Index(cnt)
	log.Println("data from container :", dataCon)
	d.SetId(dataCon.S("id").String())
	d.Set("name", StripQuotes(dataCon.S("name").String()))
	conEmails := dataCon.S("emails").Data().([]interface{})
	emailList := make([]string, 0)
	for _, val := range conEmails {
		tpMap := val.(map[string]interface{})
		emailList = append(emailList, tpMap["email"].(string))
	}
	d.Set("emails", emailList)
	return nil
}
