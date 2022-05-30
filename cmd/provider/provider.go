package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFoo() *schema.Resource {
	return &schema.Resource{}
}

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"foo": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"email": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"foo_thing": resourceFoo(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	accessKey := d.Get("access_key").(string)
	var diags diag.Diagnostics
	if accessKey != "" {
		return &FooClient{accessKey: accessKey}, diags
	}
	return nil, nil
}
