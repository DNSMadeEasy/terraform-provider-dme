package dme

import (
	"fmt"
	"log"
	"testing"

	"github.com/DNSMadeEasy/dme-go-client/client"
	"github.com/DNSMadeEasy/dme-go-client/container"
	"github.com/DNSMadeEasy/dme-go-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccACL_Basic(t *testing.T) {
	var acl models.ACLAttribute
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDMEACLDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDMEACLConfig_basic("transferacl", "1.2.3.4"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDMEACLExists("dme_transfer_acl.example", &acl),
					testAccCheckDMEACLAttributes("transferacl", "1.2.3.4", &acl),
				),
			},
		},
	})
}

func TestAccDMEACL_Update(t *testing.T) {
	var acl models.ACLAttribute

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDMEACLDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDMEACLConfig_basic("transferacl", "1.2.3.4"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDMEACLExists("dme_transfer_acl.example", &acl),
					testAccCheckDMEACLAttributes("transferacl", "1.2.3.4", &acl),
				),
			},
			{
				Config: testAccCheckDMEACLConfig_basic("transferacl", "1.2.3.6"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDMEACLExists("dme_transfer_acl.example", &acl),
					testAccCheckDMEACLAttributes("transferacl", "1.2.3.6", &acl),
				),
			},
		},
	})
}

func testAccCheckDMEACLConfig_basic(name string, ips string) string {
	return fmt.Sprintf(`
	resource "dme_transfer_acl" "example" {
		name = "%s"
		ips = ["%s"]
	}
	`, name, ips)
}

func testAccCheckDMEACLExists(name string, acl *models.ACLAttribute) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("ACL %s not found", name)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No ACL id was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		con, err := client.GetbyId("dns/transferAcl/" + rs.Primary.ID)
		if err != nil {
			return err
		}

		tp, _ := aclfromcontainer(con)

		*acl = *tp
		return nil

	}
}

func aclfromcontainer(con *container.Container) (*models.ACLAttribute, error) {
	acl := models.ACLAttribute{}

	acl.Name = StripQuotes(con.S("name").String())
	ips := con.S("ips").Data().([]interface{})
	listips := make([]string, 0)
	listips = append(listips, ips[0].(string))
	acl.Ips = listips

	return &acl, nil

}

func testAccCheckDMEACLDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "dme_transfer_acl" {
			_, err := client.GetbyId("dns/transferAcl/" + rs.Primary.ID)
			if err == nil {
				return fmt.Errorf("ACL is still exists")
			}
		} else {
			continue
		}

	}
	return nil
}

func testAccCheckDMEACLAttributes(name string, ips string, acl *models.ACLAttribute) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if name != acl.Name {
			return fmt.Errorf("Bad ACL name %s", acl.Name)
		}
		str := acl.Ips[0]
		log.Println("acl.ips value: ", str)
		if ips != str {
			return fmt.Errorf("Bad IP value for ACL %s", acl.Ips)
		}
		return nil
	}
}
