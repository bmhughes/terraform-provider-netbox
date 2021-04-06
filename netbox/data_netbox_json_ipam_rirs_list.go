package netbox

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	netboxclient "github.com/netbox-community/go-netbox/netbox/client"
)

func dataNetboxJSONIpamRirsList() *schema.Resource {
	return &schema.Resource{
		Read: dataNetboxJSONIpamRirsListRead,

		Schema: map[string]*schema.Schema{
			"json": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataNetboxJSONIpamRirsListRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*netboxclient.NetBoxAPI)

	list, err := client.Ipam.IpamRirsList(nil, nil)
	if err != nil {
		return err
	}

	j, _ := json.Marshal(list.Payload.Results)

	d.Set("json", string(j))
	d.SetId("NetboxJSONIpamRirsList")

	return nil
}
