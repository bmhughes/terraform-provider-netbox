---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "netbox_virtualization_interface Resource - terraform-provider-netbox"
subcategory: ""
description: |-
  Manage an interface (virtualization module) resource within Netbox.
---

# netbox_virtualization_interface (Resource)

Manage an interface (virtualization module) resource within Netbox.

## Example Usage

```terraform
resource "netbox_virtualization_interface" "interface_test" {
  name = "default"
  virtualmachine_id = netbox_virtualization_vm.vm_test.id
  mac_address = "AA:AA:AA:AA:AA:AA"
  mtu = 1500
  description = "Interface de test"

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

- `name` (String) Description for this interface (virtualization module)
- `virtualmachine_id` (Number) ID of the VM where this interface (virtualization module) is attached to.

### Optional

- `custom_field` (Block Set) Existing custom fields to associate to this interface (virtualization module). (see [below for nested schema](#nestedblock--custom_field))
- `description` (String) Description for this interface (virtualization module).
- `enabled` (Boolean) true or false (true by default)
- `mac_address` (String) Mac address for this interface (virtualization module)
- `mode` (String) The mode among access, tagged, tagged-all.
- `mtu` (Number) The MTU between 1 and 65536 for this interface (virtualization module).
- `tag` (Block Set) Existing tag to associate to this interface (virtualization module). (see [below for nested schema](#nestedblock--tag))
- `tagged_vlans` (Set of Number) List of vlan id tagged for this interface (virtualization module)
- `untagged_vlan` (Number) Vlan ID untagged for this interface (virtualization module).

### Read-Only

- `content_type` (String) The content type of this interface (virtualization module).
- `id` (String) The ID of this resource.
- `type` (String) Type of interface among virtualization.vminterface for VM or dcim.interface for device

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


