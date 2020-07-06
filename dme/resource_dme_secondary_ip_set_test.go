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

func TestAccSecondaryIPSet_Basic(t *testing.T) {
	var secondIP models.SecondaryIPSet
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDMESecondaryIPSetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDMESecondaryIPSetConfig_basic("12.23.56.45"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDMESecondaryIPSetExists("dme_secondary_ip_set.example", &secondIP),
					testAccCheckDMESecondaryIPSetAttributes("12.23.56.45", &secondIP),
				),
			},
		},
	})
}

func TestAccDMESecondaryIPSet_Update(t *testing.T) {
	var secondIP models.SecondaryIPSet

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDMESecondaryIPSetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDMESecondaryIPSetConfig_basic("12.23.56.45"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDMESecondaryIPSetExists("dme_secondary_ip_set.example", &secondIP),
					testAccCheckDMESecondaryIPSetAttributes("12.23.56.45", &secondIP),
				),
			},
			{
				Config: testAccCheckDMESecondaryIPSetConfig_basic("102.23.56.84"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDMESecondaryIPSetExists("dme_secondary_ip_set.example", &secondIP),
					testAccCheckDMESecondaryIPSetAttributes("102.23.56.84", &secondIP),
				),
			},
		},
	})
}

func testAccCheckDMESecondaryIPSetConfig_basic(ip string) string {
	return fmt.Sprintf(`
	resource "dme_secondary_ip_set" "example" {
		name = "check_ipset"
		ips = ["%s"]
	}
	`, ip)
}

func testAccCheckDMESecondaryIPSetExists(name string, secondIP *models.SecondaryIPSet) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Secondary IP Set %s not found", name)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No Secondary IP Set id was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		con, err := client.GetbyId("dns/secondary/ipSet/" + rs.Primary.ID)
		if err != nil {
			return err
		}

		tp, _ := secondaryIPSetfromcontainer(con)

		*secondIP = *tp
		return nil
	}
}

func secondaryIPSetfromcontainer(con *container.Container) (*models.SecondaryIPSet, error) {
	secondIP := models.SecondaryIPSet{}

	secondIP.Name = StripQuotes(con.S("name").String())
	secondIP.IPs = con.S("ips").Data().([]interface{})

	return &secondIP, nil
}

func testAccCheckDMESecondaryIPSetDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "dme_secondary_ip_set" {
			_, err := client.GetbyId("dns/secondary/ipSet/" + rs.Primary.ID)
			if err == nil {
				return fmt.Errorf("Secondary IP Set is still exists")
			}
		} else {
			continue
		}
	}
	return nil
}

func testAccCheckDMESecondaryIPSetAttributes(ip string, secondIP *models.SecondaryIPSet) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if "check_ipset" != secondIP.Name {
			return fmt.Errorf("Bad Secondary IP set name %s", secondIP.Name)
		}
		if ip != secondIP.IPs[0] {
			return fmt.Errorf("Bad IP for secondary IP set %s", secondIP.IPs[0])
		}
		return nil
	}
}
