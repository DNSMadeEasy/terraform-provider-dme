package dme

import (
	"fmt"
	"testing"

	"github.com/DNSMadeEasy/dme-go-client/client"
	"github.com/DNSMadeEasy/dme-go-client/container"
	"github.com/DNSMadeEasy/dme-go-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccContactList_Basic(t *testing.T) {
	var contact models.ContactList
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDMEContactListDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDMEContactListConfig_basic("check@gmail.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDMEContactListExists("dme_contact_list.example", &contact),
					testAccCheckDMEContactListAttributes("check@gmail.com", &contact),
				),
			},
		},
	})
}

func TestAccDMEContactList_Update(t *testing.T) {
	var contact models.ContactList

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDMEContactListDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDMEContactListConfig_basic("check@gmail.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDMEContactListExists("dme_contact_list.example", &contact),
					testAccCheckDMEContactListAttributes("check@gmail.com", &contact),
				),
			},
			{
				Config: testAccCheckDMEContactListConfig_basic("check01@gmail.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDMEContactListExists("dme_contact_list.example", &contact),
					testAccCheckDMEContactListAttributes("check01@gmail.com", &contact),
				),
			},
		},
	})
}

func testAccCheckDMEContactListConfig_basic(email string) string {
	return fmt.Sprintf(`
	resource "dme_contact_list" "example" {
		name = "check_contactlist01"
		emails = ["%s"]
	}
	`, email)
}

func testAccCheckDMEContactListExists(name string, contact *models.ContactList) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("contact list %s not found", name)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No contact list id was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		con, err := client.GetbyId("contactList/" + rs.Primary.ID)
		if err != nil {
			return err
		}

		tp, _ := contactlistfromcontainer(con)

		*contact = *tp
		return nil
	}
}

func contactlistfromcontainer(con *container.Container) (*models.ContactList, error) {
	contactlist := models.ContactList{}

	contactlist.Name = StripQuotes(con.S("name").String())
	conEmails := con.S("emails").Data().([]interface{})
	emailList := make([]interface{}, 0)
	for _, val := range conEmails {
		tpMap := val.(map[string]interface{})
		emailList = append(emailList, tpMap["email"].(string))
	}
	contactlist.Emails = emailList

	return &contactlist, nil
}

func testAccCheckDMEContactListDestroy(s *terraform.State) error {
	// time.Sleep(10 * time.Minute)
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "dme_contact_list" {
			_, err := client.GetbyId("contactList/" + rs.Primary.ID)
			if err == nil {
				return fmt.Errorf("ContactList is still exists")
			}
		} else {
			continue
		}

	}
	return nil
}

func testAccCheckDMEContactListAttributes(email string, contact *models.ContactList) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if "check_contactlist01" != contact.Name {
			return fmt.Errorf("Bad Contact List name %s", contact.Name)
		}
		if email != contact.Emails[0] {
			return fmt.Errorf("Bad email for contact list %s", contact.Emails[0])
		}
		return nil
	}
}
