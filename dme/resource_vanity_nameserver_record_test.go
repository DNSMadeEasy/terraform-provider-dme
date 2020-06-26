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

func TestAccVNS_Basic(t *testing.T) {
	var vns models.Vanity
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDmeVNSDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDmeVNSConfig_basic("checkvns"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDmeVNSExists("dme_vanity_nameserver_record.vanityrecord", &vns),
					testAccCheckDmeVNSAttributes("checkvns", &vns),
				),
			},
		},
	})
}

func TestAccDmeVNS_Update(t *testing.T) {
	var vns models.Vanity

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDmeVNSDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDmeVNSConfig_basic("vnscheck"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDmeVNSExists("dme_vanity_nameserver_record.vanityrecord", &vns),
					testAccCheckDmeVNSAttributes("vnscheck", &vns),
				),
			},
			{
				Config: testAccCheckDmeVNSConfig_basic("yashcheckvns"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDmeVNSExists("dme_vanity_nameserver_record.vanityrecord", &vns),
					testAccCheckDmeVNSAttributes("yashcheckvns", &vns),
				),
			},
		},
	})
}

func testAccCheckDmeVNSConfig_basic(name string) string {
	return fmt.Sprintf(`
	resource "dme_vanity_nameserver_record" "vanityrecord" {
		name                 = "%s"
		servers              = ["abc.com.", "xyz.com.", "yash.com."]
		public_config        = false
		default_config       = false
		name_server_group_id = 1
	  }
	`, name)
}

func testAccCheckDmeVNSExists(name string, vns *models.Vanity) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("VNS record %s not found", name)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No VNS id was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.GetbyId("dns/vanity/" + rs.Primary.ID)
		if err != nil {
			return err
		}

		tp, _ := vnsfromcontainer(cont)

		*vns = *tp
		return nil

	}
}

func vnsfromcontainer(con *container.Container) (*models.Vanity, error) {

	vns := models.Vanity{}

	vns.Name = StripQuotes(con.S("name").String())

	return &vns, nil

}

func testAccCheckDmeVNSDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "dme_vanity_nameserver_record" {
			_, err := client.GetbyId("dns/vanity/" + rs.Primary.ID)
			if err == nil {
				return fmt.Errorf("Vanity Name Server is still exists")
			}
		} else {
			continue
		}

	}
	return nil
}

func testAccCheckDmeVNSAttributes(name string, vns *models.Vanity) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if name != vns.Name {
			return fmt.Errorf("Bad VNS name %s", vns.Name)
		}
		return nil
	}
}
