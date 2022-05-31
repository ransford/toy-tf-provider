package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFoo() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFooCreate,
		ReadContext:   resourceFooRead,
		UpdateContext: resourceFooUpdate,
		DeleteContext: resourceFooDelete,
		Schema: map[string]*schema.Schema{
			"beep": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceFooCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*FooClient)
	s := c.CreateFoo()
	d.SetId(s)
	resourceFooRead(ctx, d, m) // TODO is this necessary?
	return diags
}

func resourceFooRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	id := d.Id()
	c := m.(*FooClient)
	beep := c.GetFoo(id)
	d.Set("beep", beep)
	return diags
}

func resourceFooUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*FooClient)
	id := d.Id()

	if d.HasChange("beep") {
		newBeep := d.Get("beep").(string)
		c.SetFoo(id, newBeep)
	}

	return resourceFooRead(ctx, d, m)
}

func resourceFooDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	panic("delete not implemented")
}

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"access_key": {
				Type:     schema.TypeString,
				Required: true,
			},
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
