package netbox

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	netboxclient "github.com/smutel/go-netbox/netbox/client"
	"github.com/smutel/go-netbox/netbox/client/extras"
)

func dataNetboxJSONExtrasTagsList() *schema.Resource {
	return &schema.Resource{
		Description: "Get json output from the extras_tags_list Netbox endpoint.",
		Read:        dataNetboxJSONExtrasTagsListRead,

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

func dataNetboxJSONExtrasTagsListRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*netboxclient.NetBoxAPI)

	params := extras.NewExtrasTagsListParams()
	limit := int64(d.Get("limit").(int))
	params.Limit = &limit

	list, err := client.Extras.ExtrasTagsList(params, nil)
	if err != nil {
		return err
	}

	tmp := list.Payload.Results
	resultLength := int64(len(tmp))
	desiredLength := *list.Payload.Count
	if limit > 0 && limit < desiredLength {
		desiredLength = limit
	}
	if limit == 0 || resultLength < desiredLength {
		limit = resultLength
	}
	offset := limit
	params.Offset = &offset
	for int64(len(tmp)) < desiredLength {
		offset = int64(len(tmp))
		if limit > desiredLength - offset {
			limit = desiredLength - offset
		}
		list, err = client.Extras.ExtrasTagsList(params, nil)
		if err != nil {
			return err
		}
		tmp = append(tmp, list.Payload.Results...)
	}

	j, _ := json.Marshal(tmp)

	d.Set("json", string(j))
	d.SetId("NetboxJSONExtrasTagsList")

	return nil
}
