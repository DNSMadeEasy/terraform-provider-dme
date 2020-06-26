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

func TestAccFolder_Basic(t *testing.T) {
	var folder models.Folder
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDmeFolderDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDmeFolderConfig_basic("checkfolder"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDmeFolderExists("dme_folder_record.folderrecord", &folder),
					testAccCheckDmeFolderAttributes("checkfolder", &folder),
				),
			},
		},
	})
}

func TestAccDmeFolder_Update(t *testing.T) {
	var folder models.Folder

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDmeFolderDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDmeFolderConfig_basic("foldercheck"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDmeFolderExists("dme_folder_record.folderrecord", &folder),
					testAccCheckDmeFolderAttributes("foldercheck", &folder),
				),
			},
			{
				Config: testAccCheckDmeFolderConfig_basic("yashcheckfolder"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDmeFolderExists("dme_folder_record.folderrecord", &folder),
					testAccCheckDmeFolderAttributes("yashcheckfolder", &folder),
				),
			},
		},
	})
}

func testAccCheckDmeFolderConfig_basic(name string) string {
	return fmt.Sprintf(`
	resource "dme_folder_record" "folderrecord" {
		name        = "%s"
		domains     = ["6994874", "6994926", "6994935"]
		secondaries = ["132212", "132207", "132182"]
		folder_permissions {
		  permission = 7
		  group_id   = 159249
		  group_name = "Default"
		}
	  }
	  
	`, name)
}

func testAccCheckDmeFolderExists(name string, folder *models.Folder) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Folder record %s not found", name)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No Folder id was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.GetbyId("security/folder/" + rs.Primary.ID)
		if err != nil {
			return err
		}

		tp, _ := folderfromcontainer(cont)

		*folder = *tp
		return nil

	}
}

func folderfromcontainer(con *container.Container) (*models.Folder, error) {

	folder := models.Folder{}

	folder.Name = StripQuotes(con.S("name").String())

	return &folder, nil

}

func testAccCheckDmeFolderDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "dme_folder_record" {
			_, err := client.GetbyId("security/folder/" + rs.Primary.ID)
			if err == nil {
				return fmt.Errorf("Folder still exists")
			}
		} else {
			continue
		}

	}
	return nil
}

func testAccCheckDmeFolderAttributes(name string, folder *models.Folder) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if name != folder.Name {
			return fmt.Errorf("Bad Folder name %s", folder.Name)
		}
		return nil
	}
}
