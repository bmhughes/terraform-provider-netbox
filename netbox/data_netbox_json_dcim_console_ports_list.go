package netbox

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	netboxclient "github.com/netbox-community/go-netbox/netbox/client"
)

func dataNetboxJSONDcimConsolePortsList() *schema.Resource {
	return &schema.Resource{
		Read: dataNetboxJSONDcimConsolePortsListRead,

		Schema: map[string]*schema.Schema{
			"json": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataNetboxJSONDcimConsolePortsListRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*netboxclient.NetBoxAPI)

	list, err := client.Dcim.DcimConsolePortsList(nil, nil)
	if err != nil {
		return err
	}

	j, _ := json.Marshal(list.Payload.Results)

	d.Set("json", string(j))
	d.SetId("NetboxJSONDcimConsolePortsList")

	return nil
}
