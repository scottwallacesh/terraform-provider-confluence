package confluence

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSpace() *schema.Resource {
	return &schema.Resource{
		Create: resourceSpaceCreate,
		Read:   resourceSpaceRead,
		Update: resourceSpaceUpdate,
		Delete: resourceSpaceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(0, 255),
			},
			"url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceSpaceCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	contentRequest := spaceFromResourceData(d)
	contentResponse, err := client.CreateSpace(contentRequest)

	if err != nil {
		return err
	}
	d.SetId(strconv.Itoa(contentResponse.Id))
	return resourceSpaceRead(d, m)
}

func resourceSpaceRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	contentResponse, err := client.GetSpace(d.Get("key").(string))
	if err != nil {
		d.SetId("")
		return err
	}
	return updateResourceDataFromSpace(d, contentResponse, client)
}

func resourceSpaceUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	contentRequest := spaceFromResourceData(d)
	_, err := client.UpdateSpace(contentRequest)
	if err != nil {
		d.SetId("")
		return err
	}
	return resourceSpaceRead(d, m)
}

func resourceSpaceDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	err := client.DeleteSpace(d.Get("key").(string))
	if err != nil {
		return err
	}
	// d.SetId("") is automatically called assuming delete returns no errors
	return nil
}

func spaceFromResourceData(d *schema.ResourceData) *Space {
	id, _ := strconv.Atoi(d.Id())
	result := &Space{
		Id:   id,
		Key:  d.Get("key").(string),
		Name: d.Get("name").(string),
	}
	return result
}

func updateResourceDataFromSpace(d *schema.ResourceData, space *Space, client *Client) error {
	d.SetId(strconv.Itoa(space.Id))
	m := map[string]interface{}{
		"key":  space.Key,
		"name": space.Name,
		"url":  space.Links.Base + space.Links.WebUI,
	}
	for k, v := range m {
		err := d.Set(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}
