package netbox

import (
	"context"
	"regexp"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	netboxclient "github.com/smutel/go-netbox/netbox/client"
	"github.com/smutel/go-netbox/netbox/client/dcim"
)

func dataNetboxDcimPlatform() *schema.Resource {
	return &schema.Resource{
		Description: "Get info about platform (dcim module) from netbox.",
		ReadContext: dataNetboxDcimPlatformRead,

		Schema: map[string]*schema.Schema{
			"content_type": {
				Description: "Content type of this platform (dcim module).",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"slug": {
				Description: "Slug of this platform (dcim module).",
				Type:        schema.TypeString,
				Required:    true,
				ValidateFunc: validation.StringMatch(
					regexp.MustCompile("^[-a-zA-Z0-9_]{1,50}$"),
					"Must be like ^[-a-zA-Z0-9_]{1,50}$"),
			},
		},
	}
}

func dataNetboxDcimPlatformRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*netboxclient.NetBoxAPI)

	slug := d.Get("slug").(string)

	resource := dcim.NewDcimPlatformsListParams().WithSlug(&slug)

	list, err := client.Dcim.DcimPlatformsList(resource, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	if *list.Payload.Count < 1 {
		return diag.Errorf("Your query returned no results. " + 
			"Please change your search criteria and try again.")
	} else if *list.Payload.Count > 1 {
		return diag.Errorf("Your query returned more than one result. " +
			"Please try a more specific search criteria.")
	}

	r := list.Payload.Results[0]
	d.SetId(strconv.FormatInt(r.ID, 10))
	d.Set("content_type", convertURIContentType(r.URL))

	return nil
}
