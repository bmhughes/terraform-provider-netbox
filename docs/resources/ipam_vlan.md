---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "netbox_ipam_vlan Resource - terraform-provider-netbox"
subcategory: ""
description: |-
  Manage a vlan (ipam module) within Netbox.
---

# netbox_ipam_vlan (Resource)

Manage a vlan (ipam module) within Netbox.

## Example Usage

```terraform
resource "netbox_ipam_vlan" "vlan_test" {
  vlan_id = 100
  name = "TestVlan"
  site_id = netbox_ipam_vlan_group.vlan_group_test.site_id
  description = "VLAN created by terraform"
  vlan_group_id = netbox_ipam_vlan_group.vlan_group_test.id
  tenant_id = netbox_tenancy_tenant.tenant_test.id
  role_id = data.netbox_ipam_role.vlan_role_production.id

  tag {
    name = "tag1"
    slug = "tag1"
  }

  custom_field {
    name = "cf_boolean"
    type = "boolean"
    value = "true"
  }

  custom_field {
    name = "cf_date"
    type = "date"
    value = "2020-12-25"
  }

  custom_field {
    name = "cf_text"
    type = "text"
    value = "some text"
  }

  custom_field {
    name = "cf_integer"
    type = "integer"
    value = "10"
  }

  custom_field {
    name = "cf_selection"
    type = "selection"
    value = "1"
  }

  custom_field {
    name = "cf_url"
    type = "url"
    value = "https://github.com"
  }

  custom_field {
    name = "cf_multiple_selection"
    type = "multiple"
    value = "0,1"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name for this vlan (ipam module).
- `vlan_id` (Number) The ID of the vlan (vlan tag).

### Optional

- `custom_field` (Block Set) Existing custom fields to associate to this vlan (ipam module). (see [below for nested schema](#nestedblock--custom_field))
- `description` (String) The description of this vlan (ipam module).
- `role_id` (Number) ID of the role attached to this vlan (ipam module).
- `site_id` (Number) ID of the site where this vlan (ipam module) is located.
- `status` (String) The description of this vlan (ipam module).
- `tag` (Block Set) Existing tag to associate to this vlan (ipam module). (see [below for nested schema](#nestedblock--tag))
- `tenant_id` (Number) ID of the tenant where this vlan (ipam module) is attached.
- `vlan_group_id` (Number) ID of the group where this vlan (ipam module) belongs to.

### Read-Only

- `content_type` (String) The content type of this vlan (ipam module).
- `id` (String) The ID of this resource.

<a id="nestedblock--custom_field"></a>
### Nested Schema for `custom_field`

Required:

- `name` (String) Name of the existing custom field.
- `type` (String) Type of the existing custom field (text, integer, boolean, url, selection, multiple).
- `value` (String) Value of the existing custom field.


<a id="nestedblock--tag"></a>
### Nested Schema for `tag`

Required:

- `name` (String) Name of the existing tag.
- `slug` (String) Slug of the existing tag.


