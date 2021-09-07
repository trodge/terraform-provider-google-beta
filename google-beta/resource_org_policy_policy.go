// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: DCL     ***
//
// ----------------------------------------------------------------------------
//
//     This file is managed by Magic Modules (https://github.com/GoogleCloudPlatform/magic-modules)
//     and is based on the DCL (https://github.com/GoogleCloudPlatform/declarative-resource-client-library).
//     Changes will need to be made to the DCL or Magic Modules instead of here.
//
//     We are not currently able to accept contributions to this file. If changes
//     are required, please file an issue at https://github.com/hashicorp/terraform-provider-google/issues/new/choose
//
// ----------------------------------------------------------------------------

package google

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	dcl "github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	orgpolicy "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/orgpolicy/beta"
)

func resourceOrgPolicyPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceOrgPolicyPolicyCreate,
		Read:   resourceOrgPolicyPolicyRead,
		Update: resourceOrgPolicyPolicyUpdate,
		Delete: resourceOrgPolicyPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceOrgPolicyPolicyImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: ``,
			},

			"parent": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      ``,
			},

			"spec": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        OrgPolicyPolicySpecSchema(),
			},
		},
	}
}

func OrgPolicyPolicySpecSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"inherit_from_parent": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: ``,
			},

			"reset": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: ``,
			},

			"rules": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: ``,
				Elem:        OrgPolicyPolicySpecRulesSchema(),
			},

			"etag": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: ``,
			},

			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: ``,
			},
		},
	}
}

func OrgPolicyPolicySpecRulesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"allow_all": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: ``,
			},

			"condition": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        OrgPolicyPolicySpecRulesConditionSchema(),
			},

			"deny_all": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: ``,
			},

			"enforce": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: ``,
			},

			"values": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        OrgPolicyPolicySpecRulesValuesSchema(),
			},
		},
	}
}

func OrgPolicyPolicySpecRulesConditionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: ``,
			},

			"expression": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: ``,
			},

			"location": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: ``,
			},

			"title": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: ``,
			},
		},
	}
}

func OrgPolicyPolicySpecRulesValuesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"allowed_values": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"denied_values": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceOrgPolicyPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := &orgpolicy.Policy{
		Name:   dcl.String(d.Get("name").(string)),
		Parent: dcl.String(d.Get("parent").(string)),
		Spec:   expandOrgPolicyPolicySpec(d.Get("spec")),
	}

	id, err := obj.ID()
	if err != nil {
		return err
	}
	d.SetId(id)
	createDirective := CreateDirective
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}
	billingProject := ""
	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := NewDCLOrgPolicyClient(config, userAgent, billingProject)
	res, err := client.ApplyPolicy(context.Background(), obj, createDirective...)

	if _, ok := err.(dcl.DiffAfterApplyError); ok {
		log.Printf("[DEBUG] Diff after apply returned from the DCL: %s", err)
	} else if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error creating Policy: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Policy %q: %#v", d.Id(), res)

	return resourceOrgPolicyPolicyRead(d, meta)
}

func resourceOrgPolicyPolicyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := &orgpolicy.Policy{
		Name:   dcl.String(d.Get("name").(string)),
		Parent: dcl.String(d.Get("parent").(string)),
		Spec:   expandOrgPolicyPolicySpec(d.Get("spec")),
	}

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}
	billingProject := ""
	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := NewDCLOrgPolicyClient(config, userAgent, billingProject)
	res, err := client.GetPolicy(context.Background(), obj)
	if err != nil {
		// Resource not found
		d.SetId("")
		return err
	}

	if err = d.Set("name", res.Name); err != nil {
		return fmt.Errorf("error setting name in state: %s", err)
	}
	if err = d.Set("parent", res.Parent); err != nil {
		return fmt.Errorf("error setting parent in state: %s", err)
	}
	if err = d.Set("spec", flattenOrgPolicyPolicySpec(res.Spec)); err != nil {
		return fmt.Errorf("error setting spec in state: %s", err)
	}

	return nil
}
func resourceOrgPolicyPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := &orgpolicy.Policy{
		Name:   dcl.String(d.Get("name").(string)),
		Parent: dcl.String(d.Get("parent").(string)),
		Spec:   expandOrgPolicyPolicySpec(d.Get("spec")),
	}
	directive := UpdateDirective
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""
	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := NewDCLOrgPolicyClient(config, userAgent, billingProject)
	res, err := client.ApplyPolicy(context.Background(), obj, directive...)

	if _, ok := err.(dcl.DiffAfterApplyError); ok {
		log.Printf("[DEBUG] Diff after apply returned from the DCL: %s", err)
	} else if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error updating Policy: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Policy %q: %#v", d.Id(), res)

	return resourceOrgPolicyPolicyRead(d, meta)
}

func resourceOrgPolicyPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := &orgpolicy.Policy{
		Name:   dcl.String(d.Get("name").(string)),
		Parent: dcl.String(d.Get("parent").(string)),
		Spec:   expandOrgPolicyPolicySpec(d.Get("spec")),
	}

	log.Printf("[DEBUG] Deleting Policy %q", d.Id())
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}
	billingProject := ""
	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := NewDCLOrgPolicyClient(config, userAgent, billingProject)
	if err := client.DeletePolicy(context.Background(), obj); err != nil {
		return fmt.Errorf("Error deleting Policy: %s", err)
	}

	log.Printf("[DEBUG] Finished deleting Policy %q", d.Id())
	return nil
}

func resourceOrgPolicyPolicyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)

	if err := resourceOrgPolicyPolicyCustomImport(d, config); err != nil {
		return nil, fmt.Errorf("error encountered in import: %v", err)
	}

	return []*schema.ResourceData{d}, nil
}

func expandOrgPolicyPolicySpec(o interface{}) *orgpolicy.PolicySpec {
	if o == nil {
		return orgpolicy.EmptyPolicySpec
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return orgpolicy.EmptyPolicySpec
	}
	obj := objArr[0].(map[string]interface{})
	return &orgpolicy.PolicySpec{
		InheritFromParent: dcl.Bool(obj["inherit_from_parent"].(bool)),
		Reset:             dcl.Bool(obj["reset"].(bool)),
		Rules:             expandOrgPolicyPolicySpecRulesArray(obj["rules"]),
	}
}

func flattenOrgPolicyPolicySpec(obj *orgpolicy.PolicySpec) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"inherit_from_parent": obj.InheritFromParent,
		"reset":               obj.Reset,
		"rules":               flattenOrgPolicyPolicySpecRulesArray(obj.Rules),
		"etag":                obj.Etag,
		"update_time":         obj.UpdateTime,
	}

	return []interface{}{transformed}

}
func expandOrgPolicyPolicySpecRulesArray(o interface{}) []orgpolicy.PolicySpecRules {
	if o == nil {
		return nil
	}

	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}

	items := make([]orgpolicy.PolicySpecRules, 0, len(objs))
	for _, item := range objs {
		i := expandOrgPolicyPolicySpecRules(item)
		items = append(items, *i)
	}

	return items
}

func expandOrgPolicyPolicySpecRules(o interface{}) *orgpolicy.PolicySpecRules {
	if o == nil {
		return orgpolicy.EmptyPolicySpecRules
	}

	obj := o.(map[string]interface{})
	return &orgpolicy.PolicySpecRules{
		AllowAll:  dcl.Bool(obj["allow_all"].(bool)),
		Condition: expandOrgPolicyPolicySpecRulesCondition(obj["condition"]),
		DenyAll:   dcl.Bool(obj["deny_all"].(bool)),
		Enforce:   dcl.Bool(obj["enforce"].(bool)),
		Values:    expandOrgPolicyPolicySpecRulesValues(obj["values"]),
	}
}

func flattenOrgPolicyPolicySpecRulesArray(objs []orgpolicy.PolicySpecRules) []interface{} {
	if objs == nil {
		return nil
	}

	items := []interface{}{}
	for _, item := range objs {
		i := flattenOrgPolicyPolicySpecRules(&item)
		items = append(items, i)
	}

	return items
}

func flattenOrgPolicyPolicySpecRules(obj *orgpolicy.PolicySpecRules) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"allow_all": obj.AllowAll,
		"condition": flattenOrgPolicyPolicySpecRulesCondition(obj.Condition),
		"deny_all":  obj.DenyAll,
		"enforce":   obj.Enforce,
		"values":    flattenOrgPolicyPolicySpecRulesValues(obj.Values),
	}

	return transformed

}

func expandOrgPolicyPolicySpecRulesCondition(o interface{}) *orgpolicy.PolicySpecRulesCondition {
	if o == nil {
		return orgpolicy.EmptyPolicySpecRulesCondition
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return orgpolicy.EmptyPolicySpecRulesCondition
	}
	obj := objArr[0].(map[string]interface{})
	return &orgpolicy.PolicySpecRulesCondition{
		Description: dcl.String(obj["description"].(string)),
		Expression:  dcl.String(obj["expression"].(string)),
		Location:    dcl.String(obj["location"].(string)),
		Title:       dcl.String(obj["title"].(string)),
	}
}

func flattenOrgPolicyPolicySpecRulesCondition(obj *orgpolicy.PolicySpecRulesCondition) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"description": obj.Description,
		"expression":  obj.Expression,
		"location":    obj.Location,
		"title":       obj.Title,
	}

	return []interface{}{transformed}

}

func expandOrgPolicyPolicySpecRulesValues(o interface{}) *orgpolicy.PolicySpecRulesValues {
	if o == nil {
		return orgpolicy.EmptyPolicySpecRulesValues
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return orgpolicy.EmptyPolicySpecRulesValues
	}
	obj := objArr[0].(map[string]interface{})
	return &orgpolicy.PolicySpecRulesValues{
		AllowedValues: expandStringArray(obj["allowed_values"]),
		DeniedValues:  expandStringArray(obj["denied_values"]),
	}
}

func flattenOrgPolicyPolicySpecRulesValues(obj *orgpolicy.PolicySpecRulesValues) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"allowed_values": obj.AllowedValues,
		"denied_values":  obj.DeniedValues,
	}

	return []interface{}{transformed}

}
