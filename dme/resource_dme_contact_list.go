package dme

import (
	"log"
	"reflect"
	"sort"

	"github.com/DNSMadeEasy/dme-go-client/models"

	"github.com/DNSMadeEasy/dme-go-client/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceDMEContactList() *schema.Resource {
	return &schema.Resource{
		Create: resourceDMEContactListCreate,
		Read:   resourceDMEContactListRead,
		Update: resourceDMEContactListUpdate,
		Delete: resourceDMEContactListDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"emails": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceDMEContactListCreate(d *schema.ResourceData, m interface{}) error {
	dmeClient := m.(*client.Client)
	contactList := &models.ContactList{}

	if name, ok := d.GetOk("name"); ok {
		contactList.Name = name.(string)
	}

	if emails, ok := d.GetOk("emails"); ok {
		emailList := emails.([]interface{})

		finalList := make([]interface{}, 0)
		for _, email := range emailList {
			tempMap := make(map[string]interface{})
			tempMap["email"] = email

			finalList = append(finalList, tempMap)
		}
		log.Println("Final List for create :", finalList)
		contactList.Emails = finalList
	}

	con, err := dmeClient.Save(contactList, "contactList")
	if err != nil {
		return err
	}

	d.SetId(con.S("id").String())
	return resourceDMEContactListRead(d, m)
}

func resourceDMEContactListUpdate(d *schema.ResourceData, m interface{}) error {
	dmeClient := m.(*client.Client)
	contactList := &models.ContactList{}
	dn := d.Id()

	if name, ok := d.GetOk("name"); ok {
		contactList.Name = name.(string)
	}

	if emails, ok := d.GetOk("emails"); ok {
		emailList := emails.([]interface{})

		finalList := make([]interface{}, 0)
		for _, email := range emailList {
			tempMap := make(map[string]interface{})
			tempMap["email"] = email

			finalList = append(finalList, tempMap)
		}
		log.Println("Final List for create :", finalList)
		contactList.Emails = finalList
	}

	_, err := dmeClient.Update(contactList, "contactList/"+dn)
	if err != nil {
		return err
	}
	return resourceDMEContactListRead(d, m)
}

func resourceDMEContactListRead(d *schema.ResourceData, m interface{}) error {
	dmeClient := m.(*client.Client)
	dn := d.Id()

	con, err := dmeClient.GetbyId("contactList/" + dn)
	if err != nil {
		return err
	}

	d.SetId(con.S("id").String())
	d.Set("name", StripQuotes(con.S("name").String()))

	conEmails := con.S("emails").Data().([]interface{})
	emailList := make([]string, 0)
	for _, val := range conEmails {
		tpMap := val.(map[string]interface{})
		emailList = append(emailList, tpMap["email"].(string))
	}

	tfList := make([]string, 0)
	if emails, ok := d.GetOk("emails"); ok {
		tp := toListOfString(emails)
		tfList = append(tfList, tp...)
	}

	sort.Strings(emailList)
	sort.Strings(tfList)

	if reflect.DeepEqual(emailList, tfList) {
		d.Set("ips", tfList)
		return nil
	}
	d.Set("emails", emailList)
	return nil
}

func resourceDMEContactListDelete(d *schema.ResourceData, m interface{}) error {
	dmeClient := m.(*client.Client)
	dn := d.Id()

	err := dmeClient.Delete("contactList/" + dn)
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}
