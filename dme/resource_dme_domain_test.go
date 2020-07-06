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

func TestAccDomain_Basic(t *testing.T) {
	var domain models.DomainAttribute
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDMEDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDMEDomainConfig_basic("domain_test_basic1.com", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDMEDomainExists("dme_domain.example", &domain),
					testAccCheckDMEDomainAttributes("domain_test_basic1.com", "false", &domain),
				),
			},
		},
	})
}

func TestAccDMEDomain_Update(t *testing.T) {
	var domain models.DomainAttribute

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDMEDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDMEDomainConfig_basic("domain_test_update1.com", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDMEDomainExists("dme_domain.example", &domain),
					testAccCheckDMEDomainAttributes("domain_test_update1.com", "false", &domain),
				),
			},
			{
				Config: testAccCheckDMEDomainConfig_basic("domain_test_update1.com", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDMEDomainExists("dme_domain.example", &domain),
					testAccCheckDMEDomainAttributes("domain_test_update1.com", "true", &domain),
				),
			},
		},
	})
}

func testAccCheckDMEDomainConfig_basic(name string, gtd string) string {
	return fmt.Sprintf(`
	resource "dme_domain" "example" {
		name = "%s"
		gtd_enabled = "%s"
	}
	`, name, gtd)
}

func testAccCheckDMEDomainExists(name string, domain *models.DomainAttribute) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Domain %s not found", name)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No domain id was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		con, err := client.GetbyId("dns/managed/" + rs.Primary.ID)
		if err != nil {
			return err
		}

		tp, _ := domainfromcontainer(con)

		*domain = *tp
		return nil

	}
}

func domainfromcontainer(con *container.Container) (*models.DomainAttribute, error) {
	domain := models.DomainAttribute{}

	domain.Name = StripQuotes(con.S("name").String())
	domain.GtdEnabled = StripQuotes(con.S("gtdEnabled").String())
	domain.Created = StripQuotes(con.S("created").String())
	domain.Updated = StripQuotes(con.S("updated").String())

	return &domain, nil

}

func testAccCheckDMEDomainDestroy(s *terraform.State) error {
	// time.Sleep(10 * time.Minute)
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "dme_domain" {
			_, err := client.GetbyId("dns/managed/" + rs.Primary.ID)
			if err == nil {
				return fmt.Errorf("Domain is still exists")
			}
		} else {
			continue
		}

	}
	return nil
}

func testAccCheckDMEDomainAttributes(name string, gtd string, domain *models.DomainAttribute) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if name != domain.Name {
			return fmt.Errorf("Bad domain name %s", domain.Name)
		}
		if gtd != domain.GtdEnabled {
			return fmt.Errorf("Bad gtd enable flag for domain %s", domain.GtdEnabled)
		}
		return nil
	}
}
