package dme

import (
	"fmt"

	"github.com/DNSMadeEasy/dme-go-client/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceDmeFolder() *schema.Resource {
	return &schema.Resource{
		Read:          datasourceDmeFolderRead,
		SchemaVersion: 1,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"default_folder": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"domains": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"secondaries": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"folder_permissions": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"permission": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"group_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"group_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func datasourceDmeFolderRead(d *schema.ResourceData, m interface{}) error {

	dmeClient := m.(*client.Client)
	name := d.Get("name").(string)

	con, err := dmeClient.GetbyId("security/folder/")
	if err != nil {
		return err
	}

	var flag bool
	var cnt int
	data1 := con.Data().([]interface{})
	for _, info := range data1 {
		val := info.(map[string]interface{})
		if StripQuotes(val["label"].(string)) == name {
			flag = true
			break
		}
		cnt = cnt + 1
	}
	if flag != true {
		return fmt.Errorf("Folder Record of specified name not found")
	}

	dataCon := con.Index(cnt)
	dataid := StripQuotes(dataCon.S("value").String())

	cont, err := dmeClient.GetbyId("security/folder/" + dataid)
	if err != nil {
		return err
	}

	d.SetId(StripQuotes(cont.S("id").String()))
	d.Set("name", StripQuotes(cont.S("name").String()))
	d.Set("default_folder", StripQuotes(cont.S("defaultFolder").String()))

	count, err := cont.ArrayCount("folderPermissions") //container response
	if err != nil {
		return fmt.Errorf("No Permissions found")
	}
	permilist := make([]interface{}, 0)

	for i := 0; i < count; i++ {
		permiCont, err := cont.ArrayElement(i, "folderPermissions")

		if err != nil {
			return fmt.Errorf("Unable to parse the permission list")
		}

		map_permission := make(map[string]interface{})

		map_permission["permission"] = StripQuotes(permiCont.S("permission").String())
		map_permission["group_id"] = StripQuotes(permiCont.S("groupId").String())
		map_permission["group_name"] = StripQuotes(permiCont.S("groupName").String())
		permilist = append(permilist, map_permission)
	}
	d.Set("folder_permissions", permilist)

	domains := cont.S("domains").Data().([]interface{})
	listdomains := make([]string, 0)
	for _, domain := range domains {
		domainstrvalue := fmt.Sprintf("%.f", domain)
		listdomains = append(listdomains, domainstrvalue)
	}
	d.Set("domains", listdomains)

	secondariesData := cont.S("secondaries").Data().([]interface{})
	listsecondaries := make([]string, 0)
	for _, secondary := range secondariesData {
		secondarystrvalue := fmt.Sprintf("%.f", secondary)
		listsecondaries = append(listsecondaries, secondarystrvalue)
	}
	d.Set("secondaries", listsecondaries)
	return nil
}
