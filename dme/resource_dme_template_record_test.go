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

func TestAccTemplateRecords_Basic(t *testing.T) {
	var record models.TemplateRecord
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDMETemplateRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDMETemplateRecordConfig_basic("86400"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDMETemplateRecordExists("dme_template.template1", "dme_template_record.a1", &record),
					testAccCheckDMETemplateRecordAttributes("86400", &record),
				),
			},
		},
	})
}

func TestAccDMETemplateRecord_Update(t *testing.T) {
	var a models.TemplateRecord

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDMETemplateRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDMETemplateRecordConfig_basic("86600"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDMETemplateRecordExists("dme_template.template1", "dme_template_record.a1", &a),
					testAccCheckDMETemplateRecordAttributes("86600", &a),
				),
			},
			{
				Config: testAccCheckDMETemplateRecordConfig_basic("86500"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDMETemplateRecordExists("dme_template.template1", "dme_template_record.a1", &a),
					testAccCheckDMETemplateRecordAttributes("86500", &a),
				),
			},
		},
	})
}

func testAccCheckDMETemplateRecordConfig_basic(ttl string) string {
	return fmt.Sprintf(`
	resource "dme_template" "template1" {
		name = "testrecord.com"
	}

	resource "dme_template_record" "a1"{
		template_id = "${dme_template.template1.id}"
		name = "temprecord"
		ttl = "%s"
		type = "A"
		value = "1.2.3.4"
	}
	`, ttl)
}

func testAccCheckDMETemplateRecordExists(templateName string, name string, model *models.TemplateRecord) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs1, err1 := s.RootModule().Resources[templateName]
		rs2, err2 := s.RootModule().Resources[name]

		if !err1 {
			return fmt.Errorf("Template %s not found", templateName)
		}

		if !err2 {
			return fmt.Errorf("Record %s not found", name)
		}
		if rs1.Primary.ID == "" {
			return fmt.Errorf("No Template id was set")
		}
		if rs2.Primary.ID == "" {
			return fmt.Errorf("No Record id was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		resp, err := client.GetbyId("dns/template/" + rs1.Primary.ID + "/records/")

		if err != nil {
			return err
		}

		con1 := resp.S("data").Index(0)
		tp, _ := templaterecordfromcontainer(con1)

		*model = *tp
		return nil
	}
}

func testAccCheckDMETemplateRecordAttributes(ttl string, model *models.TemplateRecord) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if "temprecord" != model.Name {
			return fmt.Errorf("Bad A record name %s", model.Name)
		}

		if ttl != model.Ttl {
			return fmt.Errorf("Bad A record ttl %s", model.Ttl)
		}

		return nil
	}
}

func testAccCheckDMETemplateRecordDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)
	rs1, err1 := s.RootModule().Resources["dme_template.template1"]
	if !err1 {
		return fmt.Errorf("Template %s not found", "dme_template.template1")
	}
	domainid := rs1.Primary.ID
	for _, rs := range s.RootModule().Resources {
		log.Println("inside destroy: ", rs.Type)
		if rs.Type == "dme_template_record" {
			resp, _ := client.GetbyId("dns/template/" + domainid + "/records/")

			if resp.S("totalRecords").String() != "0" {
				return fmt.Errorf("Record is still exists")
			}
		} else {
			continue
		}
	}
	return nil
}

func templaterecordfromcontainer(con *container.Container) (*models.TemplateRecord, error) {

	model := models.TemplateRecord{}

	model.Name = StripQuotes(con.S("name").String())
	model.Ttl = StripQuotes(con.S("ttl").String())
	model.Type = StripQuotes(con.S("type").String())

	return &model, nil

}
