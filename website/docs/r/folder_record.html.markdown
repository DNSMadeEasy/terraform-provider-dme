---
layout: "dme"
page_title: "DME: dme_folder_record"
sidebar_current: "docs-dme-resource-dme_folder_record"
description: |-
    Manages Custom Folder Records for the account.
---
# dme_folder_record #
Manages Custom Folder Records for the account.
# Example Usage #
```hcl

resource "dme_folder_record" "folderrecord" {
  name           = "folder"
  domains        = ["6994874", "6994926", "6994935"]
  secondaries    = ["132212", "132207", "132182"]
  default_folder = false
  folder_permissions {
    permission = 7
    group_id   = 159249
    group_name = "Default"
  }
}

```

## Argument Reference ##
* `name` - (Required) Folder Record name
* `domains` - (Optional) A list of the primary domain IDs assigned to the folder
* `secondaries` - (Optional) A list of the secondary domain ID's assigned to the folder.
* `default_folder` - (Optional) Indicator of the folder being marked as the Default folder. Default value is false
* `folder_permissions` - (Optional) Permissions for the folder.
* `folder_permissions.group_id` - (Optional) This is static value for assigning group to the folder.
* `folder_permissions.group_name` - (Optional) This is static value for assigning group to the folder.
* `folder_permissions.permission` - (Optional) Assigning permissions for the folder. 4 for Read Only, 5 for Read and Create/Delete, 6 for Read and Edit and 7 for Read-Edit-Create/Delete all permissions.

## Attribute Reference ##
No attributes are exported