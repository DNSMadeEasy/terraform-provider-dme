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

func TestAccDomainRecords_Basic(t *testing.T) {
	var record models.ManagedDNSRecordActions
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDMERecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDMERecordConfig_basic("86400"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDMERecordExists("dme_domain.domain1", "dme_dns_record.a1", &record),
					testAccCheckDMERecordAttributes("86400", &record),
				),
			},
		},
	})
}

func TestAccDMERecord_Update(t *testing.T) {
	var a models.ManagedDNSRecordActions

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDMERecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDMERecordConfig_basic("86400"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDMERecordExists("dme_domain.domain1", "dme_dns_record.a1", &a),
					testAccCheckDMERecordAttributes("86500", &a),
				),
			},
			{
				Config: testAccCheckDMERecordConfig_basic("86500"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDMERecordExists("dme_domain.domain1", "dme_dns_record.a1", &a),
					testAccCheckDMERecordAttributes("86500", &a),
				),
			},
		},
	})
}

func testAccCheckDMERecordConfig_basic(ttl string) string {
	return fmt.Sprintf(`
	resource "dme_domain" "domain1" {
		name = "practicerecord2.com"
	}

	resource "dme_dns_record" "a1"{
		domain_id = "${dme_domain.domain1.id}"
		name = "temprecord"
		ttl = "%s"
		type = "A"
		value = "1.2.3.4"
	}
	`, ttl)
}

func testAccCheckDMERecordExists(domainName string, name string, model *models.ManagedDNSRecordActions) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs1, err1 := s.RootModule().Resources[domainName]
		rs2, err2 := s.RootModule().Resources[name]

		if !err1 {
			return fmt.Errorf("Domain %s not found", domainName)
		}

		if !err2 {
			return fmt.Errorf("Record %s not found", name)
		}
		if rs1.Primary.ID == "" {
			return fmt.Errorf("No Domain id was set")
		}
		if rs2.Primary.ID == "" {
			return fmt.Errorf("No Record id was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		resp, err := client.GetbyId("dns/managed/" + rs1.Primary.ID + "/records?recordName=temprecord&type=A")

		if err != nil {
			return err
		}

		con1 := resp.S("data").Index(0)
		tp, _ := recordfromcontainer(con1)

		*model = *tp
		return nil
	}
}

func testAccCheckDMERecordAttributes(ttl string, model *models.ManagedDNSRecordActions) resource.TestCheckFunc {
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

func testAccCheckDMERecordDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)
	rs1, err1 := s.RootModule().Resources["dme_domain.domain1"]
	if !err1 {
		return fmt.Errorf("Domain %s not found", "dme_domain.domain1")
	}
	domainid := rs1.Primary.ID
	for _, rs := range s.RootModule().Resources {

		if rs.Type == "dme_dns_record" {
			resp, _ := client.GetbyId("dns/managed/" + domainid + "/records?recordName=temprecord&type=A")

			if resp.S("totalRecords").String() != "0" {
				return fmt.Errorf("Record is still exists")
			}
		} else {
			continue
		}
	}
	return nil
}

func recordfromcontainer(con *container.Container) (*models.ManagedDNSRecordActions, error) {

	model := models.ManagedDNSRecordActions{}

	model.Name = StripQuotes(con.S("name").String())
	model.Ttl = StripQuotes(con.S("ttl").String())
	model.Type = StripQuotes(con.S("type").String())

	return &model, nil

}
