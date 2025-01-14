package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"goauthentik.io/api"
)

func resourceStageCaptcha() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceStageCaptchaCreate,
		ReadContext:   resourceStageCaptchaRead,
		UpdateContext: resourceStageCaptchaUpdate,
		DeleteContext: resourceStageCaptchaDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"public_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"private_key": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
		},
	}
}

func resourceStageCaptchaSchemaToProvider(d *schema.ResourceData) *api.CaptchaStageRequest {
	r := api.CaptchaStageRequest{
		Name:       d.Get("name").(string),
		PublicKey:  d.Get("public_key").(string),
		PrivateKey: d.Get("private_key").(string),
	}
	return &r
}

func resourceStageCaptchaCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*APIClient)

	r := resourceStageCaptchaSchemaToProvider(d)

	res, hr, err := c.client.StagesApi.StagesCaptchaCreate(ctx).CaptchaStageRequest(*r).Execute()
	if err != nil {
		return httpToDiag(hr, err)
	}

	d.SetId(res.Pk)
	return resourceStageCaptchaRead(ctx, d, m)
}

func resourceStageCaptchaRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*APIClient)

	res, hr, err := c.client.StagesApi.StagesCaptchaRetrieve(ctx, d.Id()).Execute()
	if err != nil {
		return httpToDiag(hr, err)
	}

	d.Set("name", res.Name)
	d.Set("public_key", res.PublicKey)
	return diags
}

func resourceStageCaptchaUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*APIClient)

	app := resourceStageCaptchaSchemaToProvider(d)

	res, hr, err := c.client.StagesApi.StagesCaptchaUpdate(ctx, d.Id()).CaptchaStageRequest(*app).Execute()
	if err != nil {
		return httpToDiag(hr, err)
	}

	d.SetId(res.Pk)
	return resourceStageCaptchaRead(ctx, d, m)
}

func resourceStageCaptchaDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*APIClient)
	hr, err := c.client.StagesApi.StagesCaptchaDestroy(ctx, d.Id()).Execute()
	if err != nil {
		return httpToDiag(hr, err)
	}
	return diag.Diagnostics{}
}
