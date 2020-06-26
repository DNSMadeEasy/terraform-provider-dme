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

func TestAccTemplate_Basic(t *testing.T) {
	var template models.Template
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDMETemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDMETemplateConfig_basic("check_template_basic"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDMETemplateExists("dme_template.example", &template),
					testAccCheckDMETemplateAttributes("check_template_basic", &template),
				),
			},
		},
	})
}

func TestAccDMETemplate_Update(t *testing.T) {
	var template models.Template

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDMETemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDMETemplateConfig_basic("template_update1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDMETemplateExists("dme_template.example", &template),
					testAccCheckDMETemplateAttributes("template_update1", &template),
				),
			},
			{
				Config: testAccCheckDMETemplateConfig_basic("template_update2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDMETemplateExists("dme_template.example", &template),
					testAccCheckDMETemplateAttributes("template_update2", &template),
				),
			},
		},
	})
}

func testAccCheckDMETemplateConfig_basic(name string) string {
	return fmt.Sprintf(`
	resource "dme_template" "example" {
		name = "%s"
	}
	`, name)
}

func testAccCheckDMETemplateExists(name string, template *models.Template) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Template %s not found", name)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No template id was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		con, err := client.GetbyId("dns/template/" + rs.Primary.ID)
		if err != nil {
			return err
		}

		tp, _ := templatefromcontainer(con)

		*template = *tp
		return nil
	}
}

func templatefromcontainer(con *container.Container) (*models.Template, error) {
	template := models.Template{}

	template.Name = StripQuotes(con.S("name").String())
	template.PublicTemplate = StripQuotes(con.S("publicTemplate").String())
	return &template, nil
}

func testAccCheckDMETemplateDestroy(s *terraform.State) error {
	// time.Sleep(10 * time.Minute)
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "dme_template" {
			_, err := client.GetbyId("dns/template/" + rs.Primary.ID)
			if err == nil {
				return fmt.Errorf("Template is still exists")
			}
		} else {
			continue
		}

	}
	return nil
}

func testAccCheckDMETemplateAttributes(name string, template *models.Template) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if name != template.Name {
			return fmt.Errorf("Bad template name %s", template.Name)
		}
		return nil
	}
}
