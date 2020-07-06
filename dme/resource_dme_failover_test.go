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

func TestAccFailover_Basic(t *testing.T) {
	var failover models.FailoverAttribute
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDMEFailoverDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDMEFailoverConfig_basic("1.2.3.4"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDMEFailoverExists("dme_dns_record.first", "dme_failover.a1", &failover),
					testAccCheckDMEFailoverAttributes("1.2.3.4", &failover),
				),
			},
		},
	})
}

func TestAccDMEFailover_Update(t *testing.T) {
	var a models.FailoverAttribute

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDMEFailoverDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDMEFailoverConfig_basic("1.2.3.3"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDMEFailoverExists("dme_dns_record.first", "dme_failover.a1", &a),
					testAccCheckDMEFailoverAttributes("1.2.3.3", &a),
				),
			},
			{
				Config: testAccCheckDMEFailoverConfig_basic("1.2.3.5"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDMEFailoverExists("dme_dns_record.first", "dme_failover.a1", &a),
					testAccCheckDMEFailoverAttributes("1.2.3.5", &a),
				),
			},
		},
	})
}

func testAccCheckDMEFailoverConfig_basic(ip1 string) string {
	return fmt.Sprintf(`
	resource "dme_dns_record" "first" {
		domain_id = "6999572" 
		name = "testrecord1"
		type = "A"
		value = "1.2.2.2"
		ttl = "86400"
	  }

	  resource "dme_failover" "a1" {
		record_id = "${dme_dns_record.first.id}"
		failover = "true"
		ip1 = "%s"
		ip2 = "1.2.3.9"
		protocol_id = "3"
		port = "8080"
		sensitivity = "8"
	}
	`, ip1)
}

func testAccCheckDMEFailoverExists(recordName string, failoverName string, model *models.FailoverAttribute) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs1, err1 := s.RootModule().Resources[recordName]
		rs2, err2 := s.RootModule().Resources[failoverName]

		if !err1 {
			return fmt.Errorf("Record name %s not found", recordName)
		}

		if !err2 {
			return fmt.Errorf("Failover name %s not found", failoverName)
		}
		if rs1.Primary.ID == "" {
			return fmt.Errorf("No record id was set")
		}
		if rs2.Primary.ID == "" {
			return fmt.Errorf("No Failover id was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		resp, err := client.GetbyId("monitor/" + rs1.Primary.ID)

		if err != nil {
			return err
		}

		tp, _ := failoverfromcontainer(resp)

		*model = *tp
		return nil
	}
}

func testAccCheckDMEFailoverAttributes(ip1 string, model *models.FailoverAttribute) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if "8080" != model.Port {
			return fmt.Errorf("Bad Port name %s", model.Port)
		}

		if ip1 != model.Ip1 {
			return fmt.Errorf("Bad IP1 value %s", model.Ip1)
		}

		if "1.2.3.9" != model.Ip2 {
			return fmt.Errorf("Bad IP2 value %s", model.Ip2)
		}

		if "8" != model.Sensitivity {
			return fmt.Errorf("Bad Sensitivity value %s", model.Sensitivity)
		}

		if "3" != model.ProtocolId {
			return fmt.Errorf("Bad Protocol value %s", model.ProtocolId)
		}

		return nil
	}
}

func testAccCheckDMEFailoverDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)
	_, err1 := s.RootModule().Resources["dme_dns_record.first"]
	if !err1 {
		return fmt.Errorf("Record %s not found", "dme_dns_record_first")
	}
	// domainid := rs1.Primary.ID
	for _, rs := range s.RootModule().Resources {

		if rs.Type == "dme_dns_record" {
			resp, _ := client.GetbyId("dns/managed/6999572/records?recordName=testrecord1&type=A")

			if resp.S("totalRecords").String() != "0" {
				return fmt.Errorf("Record is still exists")
			}
		} else {
			continue
		}
	}
	return nil
}

func failoverfromcontainer(con *container.Container) (*models.FailoverAttribute, error) {

	model := models.FailoverAttribute{}

	model.Port = StripQuotes(con.S("port").String())
	model.Sensitivity = StripQuotes(con.S("sensitivity").String())
	model.ProtocolId = StripQuotes(con.S("protocolId").String())
	model.Ip2 = StripQuotes(con.S("ip2").String())
	model.Ip1 = StripQuotes(con.S("ip1").String())
	return &model, nil

}
