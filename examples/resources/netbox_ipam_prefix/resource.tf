resource "netbox_ipam_prefix" "prefix_test" {
  prefix = "192.168.56.0/24"
  vlan_id = netbox_ipam_vlan.vlan_test.id
  description = "Prefix created by terraform"
  site_id = netbox_ipam_vlan_group.vlan_group_test.site_id
  role_id = data.netbox_ipam_role.vlan_role_production.id
  status = "active"

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
