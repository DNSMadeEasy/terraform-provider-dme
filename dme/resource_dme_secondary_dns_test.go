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

func TestAccSecondaryDNS_Basic(t *testing.T) {
	var secondDNS models.SecondaryDNS
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDMESecondaryDNSDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDMESecondaryDNSConfig_basic("12345"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDMESecondaryDNSExists("dme_secondary_dns.example", &secondDNS),
					testAccCheckDMESecondaryDNSAttributes("12345", &secondDNS),
				),
			},
		},
	})
}

func TestAccDMESecondaryDNS_Update(t *testing.T) {
	var secondDNS models.SecondaryDNS

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDMESecondaryDNSDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDMESecondaryDNSConfig_basic("12345"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDMESecondaryDNSExists("dme_secondary_dns.example", &secondDNS),
					testAccCheckDMESecondaryDNSAttributes("12345", &secondDNS),
				),
			},
			{
				Config: testAccCheckDMESecondaryDNSConfig_basic("22214"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDMESecondaryDNSExists("dme_secondary_dns.example", &secondDNS),
					testAccCheckDMESecondaryDNSAttributes("22214", &secondDNS),
				),
			},
		},
	})
}

func testAccCheckDMESecondaryDNSConfig_basic(ipset string) string {
	return fmt.Sprintf(`
	resource "dme_secondary_dns" "example" {
		name = "check_dns.com"
		ipset_id = "%s"
	}
	`, ipset)
}

func testAccCheckDMESecondaryDNSExists(name string, secondDNS *models.SecondaryDNS) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Secondary DNS %s not found", name)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No Secondary DNS id was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		con, err := client.GetbyId("dns/secondary/" + rs.Primary.ID)
		if err != nil {
			return err
		}

		tp, _ := secondDNSfromcontainer(con)

		*secondDNS = *tp
		return nil

	}
}

func secondDNSfromcontainer(con *container.Container) (*models.SecondaryDNS, error) {
	secondDNS := models.SecondaryDNS{}

	name := StripQuotes(con.S("name").String())
	secondDNS.Name = toListOfInterface(name)
	secondDNS.IpsetID = StripQuotes(con.S("ipSetId").String())

	return &secondDNS, nil
}

func testAccCheckDMESecondaryDNSDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "dme_secondary_dns" {
			_, err := client.GetbyId("dns/secondary/" + rs.Primary.ID)
			if err == nil {
				return fmt.Errorf("Secondary DNS is still exists")
			}
		} else {
			continue
		}

	}
	return nil
}

func testAccCheckDMESecondaryDNSAttributes(ipset string, secondDNS *models.SecondaryDNS) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if "check_dns.com" != secondDNS.Name[0] {
			return fmt.Errorf("Bad secondary DNS name %s", "temp.com")
		}
		if ipset != secondDNS.IpsetID {
			return fmt.Errorf("Bad Ip set ID for Secondary DNS %v", secondDNS.IpsetID)
		}
		return nil
	}
}
