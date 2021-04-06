package netbox

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	netboxclient "github.com/netbox-community/go-netbox/netbox/client"
)

func dataNetboxJSONDcimRackRolesList() *schema.Resource {
	return &schema.Resource{
		Read: dataNetboxJSONDcimRackRolesListRead,

		Schema: map[string]*schema.Schema{
			"json": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataNetboxJSONDcimRackRolesListRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*netboxclient.NetBoxAPI)

	list, err := client.Dcim.DcimRackRolesList(nil, nil)
	if err != nil {
		return err
	}

	j, _ := json.Marshal(list.Payload.Results)

	d.Set("json", string(j))
	d.SetId("NetboxJSONDcimRackRolesList")

	return nil
}
