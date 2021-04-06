package netbox

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	netboxclient "github.com/netbox-community/go-netbox/netbox/client"
)

func dataNetboxJSONExtrasJobResultsList() *schema.Resource {
	return &schema.Resource{
		Read: dataNetboxJSONExtrasJobResultsListRead,

		Schema: map[string]*schema.Schema{
			"json": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataNetboxJSONExtrasJobResultsListRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*netboxclient.NetBoxAPI)

	list, err := client.Extras.ExtrasJobResultsList(nil, nil)
	if err != nil {
		return err
	}

	j, _ := json.Marshal(list.Payload.Results)

	d.Set("json", string(j))
	d.SetId("NetboxJSONExtrasJobResultsList")

	return nil
}
