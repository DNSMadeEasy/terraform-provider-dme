package dme

import (
	"fmt"
	"log"
	"strconv"

	"github.com/DNSMadeEasy/dme-go-client/client"
	"github.com/DNSMadeEasy/dme-go-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceDmeFolder() *schema.Resource {
	return &schema.Resource{
		Create: resourceDmeFolderCreate,
		Read:   resourceDmeFolderRead,
		Update: resourceDmeFolderUpdate,
		Delete: resourceDmeFolderDelete,

		SchemaVersion: 1,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"default_folder": &schema.Schema{
				Type:     schema.TypeBool,
				Default:  false,
				Optional: true,
			},

			"domains": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"secondaries": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"folder_permissions": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"permission": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  7,
						},
						"group_id": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"group_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
				Optional: true,
			},
		},
	}
}

func resourceDmeFolderCreate(d *schema.ResourceData, m interface{}) error {
	dmeConnect := m.(*client.Client)

	folderAttr := &models.Folder{}

	if value, ok := d.GetOk("name"); ok {
		folderAttr.Name = value.(string)
	}

	if value, ok := d.GetOk("default_folder"); ok {
		folderAttr.DefaultFolder = value.(bool)
	}

	if value, ok := d.GetOk("domains"); ok {
		folderAttr.Domains = toListOfString(value)
	}

	if value, ok := d.GetOk("secondaries"); ok {
		folderAttr.Secondaries = toListOfString(value)
	}

	permissionlistrr := make([]interface{}, 0, 1)
	if val, ok := d.GetOk("folder_permissions"); ok {
		tp := val.(*schema.Set).List()
		map1 := make(map[string]interface{})
		inner := tp[0].(map[string]interface{})

		map1["permission"], _ = strconv.Atoi(fmt.Sprintf("%v", inner["permission"]))
		map1["groupId"], _ = strconv.Atoi(fmt.Sprintf("%v", inner["group_id"]))
		map1["groupName"] = fmt.Sprintf("%v", inner["group_name"])
		permissionlistrr = append(permissionlistrr, map1)

		folderAttr.FolderPermissions = permissionlistrr
	}

	cont, err := dmeConnect.Save(folderAttr, "security/folder/")

	if err != nil {
		log.Println("Error returned: ", err)
		return err
	}

	log.Println("Value of container: ", cont)
	id := cont.S("id")
	log.Println("Id value: ", id)
	d.SetId(fmt.Sprintf("%v", id))
	return resourceDmeFolderRead(d, m)
}

func resourceDmeFolderRead(d *schema.ResourceData, m interface{}) error {
	dmeClient := m.(*client.Client)
	dn := d.Id()

	con, err := dmeClient.GetbyId("security/folder/" + dn)
	if err != nil {
		return err
	}
	d.Set("name", StripQuotes(con.S("name").String()))
	d.Set("default_folder", StripQuotes(con.S("defaultFolder").String()))
	d.Set("domains", StripQuotes(con.S("domains").String()))
	d.Set("secondaries", StripQuotes(con.S("secondaries").String()))

	count, err := con.ArrayCount("folderPermissions") //container response
	if err != nil {
		return fmt.Errorf("No Permissions found")
	}
	permilist := make([]interface{}, 0)

	for i := 0; i < count; i++ {
		permiCont, err := con.ArrayElement(i, "folderPermissions")

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
	return nil
}

func resourceDmeFolderUpdate(d *schema.ResourceData, m interface{}) error {
	dmeConnect := m.(*client.Client)

	folderAttr := &models.Folder{}

	if value, ok := d.GetOk("name"); ok {
		folderAttr.Name = value.(string)
	}

	if value, ok := d.GetOk("default_folder"); ok {
		folderAttr.DefaultFolder = value.(bool)
	}

	if value, ok := d.GetOk("domains"); ok {
		folderAttr.Domains = toListOfString(value)
	}

	if value, ok := d.GetOk("secondaries"); ok {
		folderAttr.Secondaries = toListOfString(value)
	}

	permissionlistrr := make([]interface{}, 0, 1)
	if val, ok := d.GetOk("folder_permissions"); ok {
		tp := val.(*schema.Set).List()
		map1 := make(map[string]interface{})
		inner := tp[0].(map[string]interface{})

		map1["permission"], _ = strconv.Atoi(fmt.Sprintf("%v", inner["permission"]))
		map1["groupId"], _ = strconv.Atoi(fmt.Sprintf("%v", inner["group_id"]))
		map1["groupName"] = fmt.Sprintf("%v", inner["group_name"])
		permissionlistrr = append(permissionlistrr, map1)

		folderAttr.FolderPermissions = permissionlistrr
	}

	log.Println("Folder structure is :", folderAttr)
	dn := d.Id()

	_, err := dmeConnect.Update(folderAttr, "security/folder/"+dn)
	if err != nil {
		return err
	}
	return resourceDmeFolderRead(d, m)

}

func resourceDmeFolderDelete(d *schema.ResourceData, m interface{}) error {
	dmeClient := m.(*client.Client)
	dn := d.Id()

	err := dmeClient.Delete("security/folder/" + dn)
	if err != nil {
		return nil
	}

	d.SetId("")
	return nil
}
