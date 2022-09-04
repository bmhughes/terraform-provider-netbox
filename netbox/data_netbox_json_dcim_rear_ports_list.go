package netbox

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	netboxclient "github.com/smutel/go-netbox/netbox/client"
	"github.com/smutel/go-netbox/netbox/client/dcim"
)

func dataNetboxJSONDcimRearPortsList() *schema.Resource {
	return &schema.Resource{
		Description: "Get json output from the dcim_rear_ports_list Netbox endpoint.",
		ReadContext: dataNetboxJSONDcimRearPortsListRead,

		Schema: map[string]*schema.Schema{
			"limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "The max number of returned results. If 0 is specified, all records will be returned.",
			},
			"json": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "JSON output of the list of objects for this Netbox endpoint.",
			},
		},
	}
}

func dataNetboxJSONDcimRearPortsListRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*netboxclient.NetBoxAPI)

	params := dcim.NewDcimRearPortsListParams()
	limit := int64(d.Get("limit").(int))
	params.Limit = &limit

	list, err := client.Dcim.DcimRearPortsList(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	j, _ := json.Marshal(list.Payload.Results)

	d.Set("json", string(j))
	d.SetId("NetboxJSONDcimRearPortsList")

	return nil
}
