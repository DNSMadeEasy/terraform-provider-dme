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

func TestAccSoa_Basic(t *testing.T) {
	var soa models.Soa
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDmeSOADestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDmeSOAConfig_basic("checksoa.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDmeSOAExists("dme_custom_soa_record.soacheck", &soa),
					testAccCheckDmeSOAAttributes("checksoa.com", &soa),
				),
			},
		},
	})
}

func TestAccDmeSOA_Update(t *testing.T) {
	var soa models.Soa

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDmeSOADestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDmeSOAConfig_basic("soacheck.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDmeSOAExists("dme_custom_soa_record.soacheck", &soa),
					testAccCheckDmeSOAAttributes("soacheck.com", &soa),
				),
			},
			{
				Config: testAccCheckDmeSOAConfig_basic("yashchecksoa.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDmeSOAExists("dme_custom_soa_record.soacheck", &soa),
					testAccCheckDmeSOAAttributes("yashchecksoa.com", &soa),
				),
			},
		},
	})
}

func testAccCheckDmeSOAConfig_basic(name string) string {
	return fmt.Sprintf(`
	resource "dme_custom_soa_record" "soacheck" {
		name           = "%s"
		email          = "yashshah2crest.com."
		comp           = "yashshah2crest.com."
		ttl            = 23000
		negative_cache = 400
		refresh        = 14400
		retry          = 300
		serial         = 2009010110
		expire         = 86450
	  }
	`, name)
}

func testAccCheckDmeSOAExists(name string, soa *models.Soa) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("SOA record %s not found", name)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No SOA id was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.GetbyId("dns/soa/" + rs.Primary.ID)
		if err != nil {
			return err
		}

		tp, _ := soafromcontainer(cont)

		*soa = *tp
		return nil

	}
}

func soafromcontainer(con *container.Container) (*models.Soa, error) {

	soa := models.Soa{}

	soa.Comp = StripQuotes(con.S("comp").String())
	soa.Name = StripQuotes(con.S("name").String())
	soa.Email = StripQuotes(con.S("email").String())

	return &soa, nil

}

func testAccCheckDmeSOADestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "dme_custom_soa_record" {
			_, err := client.GetbyId("dns/soa/" + rs.Primary.ID)
			if err == nil {
				return fmt.Errorf("SOA is still exists")
			}
		} else {
			continue
		}

	}
	return nil
}

func testAccCheckDmeSOAAttributes(name string, soa *models.Soa) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if name != soa.Name {
			return fmt.Errorf("Bad SOA name %s", soa.Name)
		}
		return nil
	}
}
