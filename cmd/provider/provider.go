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
			"bar": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func resourceFooCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*FooClient)
	s := c.CreateFoo()
	d.SetId(s)
	return resourceFooRead(ctx, d, m)
}

func resourceFooRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	id := d.Id()
	c := m.(*FooClient)
	bar, err := c.GetBar(id)
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set("bar", bar)
	return diags
}

func resourceFooUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*FooClient)
	id := d.Id()

	if d.HasChange("bar") {
		newBar := d.Get("bar").(int)
		err := c.SetBar(id, newBar)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return resourceFooRead(ctx, d, m)
}

func resourceFooDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	panic("delete not implemented")
}

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"hostport": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"foo_thing": resourceFoo(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	hostport := d.Get("hostport").(string)
	return NewClient(hostport), nil
}
