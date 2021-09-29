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
	vmwareengine "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/vmwareengine/beta"
)

func resourceVmwareEnginePrivateCloud() *schema.Resource {
	return &schema.Resource{
		Create: resourceVmwareEnginePrivateCloudCreate,
		Read:   resourceVmwareEnginePrivateCloudRead,
		Update: resourceVmwareEnginePrivateCloudUpdate,
		Delete: resourceVmwareEnginePrivateCloudDelete,

		Importer: &schema.ResourceImporter{
			State: resourceVmwareEnginePrivateCloudImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The location for the resource",
			},

			"management_cluster": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "Input only. The management cluster for this private cloud. This parameter is required during creation of private cloud to provide details for the default cluster.",
				MaxItems:    1,
				Elem:        VmwareEnginePrivateCloudManagementClusterSchema(),
			},

			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The resource name of this private cloud. Resource names are schemeless URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names. For example: `projects/my-project/locations/us-west1-a/privateClouds/my-cloud`",
			},

			"network_config": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: "Required. Network configuration.",
				MaxItems:    1,
				Elem:        VmwareEnginePrivateCloudNetworkConfigSchema(),
			},

			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "User-provided description for this private cloud.",
			},

			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Labels are a way to attach lightweight metadata to resources for filtering and querying resource data. No more than 64 user labels can be associated with each resource. Label keys and values can be no longer than 63 characters, can only contain lowercase letters, numeric characters, underscores and dashes, where label keys must start with a letter and international characters are allowed. The empty string is a valid value. Labels are set on creation and updated like any other field. Specifically, to add a new label, you would need to provide all of the existing labels along with the new label. If you only provide a map with the new label, all of the old labels will be removed (probably not what is desired).",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"project": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      "The project for the resource",
			},

			"conditions": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Output only. The conditions that caused the current private cloud state. For example, private cloud provisioning failure description.",
				Elem:        VmwareEnginePrivateCloudConditionsSchema(),
			},

			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Output only. Creation time of this resource in RFC3339 text format.",
			},

			"delete_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Output only. Time the resource was marked as deleted, in RFC3339 text format.",
			},

			"expire_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Output only. Planned deletion time of this resource in RFC3339 text format.",
			},

			"hcx": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Output only. HCX appliance.",
				Elem:        VmwareEnginePrivateCloudHcxSchema(),
			},

			"nsx": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Output only. NSX appliance.",
				Elem:        VmwareEnginePrivateCloudNsxSchema(),
			},

			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Output only. State of the resource. Possible values: ACTIVE, CREATING, UPDATING, FAILED, DELETED",
			},

			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Output only. Last update time of this resource in RFC3339 text format.",
			},

			"vcenter": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Output only. Vcenter appliance.",
				Elem:        VmwareEnginePrivateCloudVcenterSchema(),
			},
		},
	}
}

func VmwareEnginePrivateCloudManagementClusterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Required. The user-provided identifier of the new `Cluster`.",
			},

			"node_count": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Required. Number of nodes in this cluster.",
			},

			"node_type_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Required. The canonical identifier of node types (`NodeType`) in this cluster. For example: standard-72.",
			},
		},
	}
}

func VmwareEnginePrivateCloudNetworkConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"management_cidr": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Required. Management CIDR used by VMWare management appliances.",
			},

			"network": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      "Required. The relative resource name of the consumer VPC network this private cloud is attached to. Specify the name in the following form: `projects/{project}/global/networks/{network_id}` where `{project}` can either be a project number or a project ID.",
			},

			"service_network": {
				Type:             schema.TypeString,
				Computed:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      "Output only. The relative resource name of the service VPC network this private cloud is attached to. The name is specified in the following form: `projects/{service_project_number}/global/networks/{network_id}`.",
			},
		},
	}
}

func VmwareEnginePrivateCloudConditionsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"code": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Output only. Machine-readable representation of the condition.",
			},

			"message": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Output only. Human-readable description of the condition.",
			},
		},
	}
}

func VmwareEnginePrivateCloudHcxSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"external_ip": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "External IP address of the appliance.",
			},

			"fdqn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Fully qualified domain name of the appliance.",
			},

			"internal_ip": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Internal IP address of the appliance.",
			},

			"version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Version of the appliance.",
			},
		},
	}
}

func VmwareEnginePrivateCloudNsxSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"external_ip": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "External IP address of the appliance.",
			},

			"fdqn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Fully qualified domain name of the appliance.",
			},

			"internal_ip": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Internal IP address of the appliance.",
			},

			"version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Version of the appliance.",
			},
		},
	}
}

func VmwareEnginePrivateCloudVcenterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"external_ip": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "External IP address of the appliance.",
			},

			"fdqn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Fully qualified domain name of the appliance.",
			},

			"internal_ip": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Internal IP address of the appliance.",
			},

			"version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Version of the appliance.",
			},
		},
	}
}

func resourceVmwareEnginePrivateCloudCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &vmwareengine.PrivateCloud{
		Location:          dcl.String(d.Get("location").(string)),
		ManagementCluster: expandVmwareEnginePrivateCloudManagementCluster(d.Get("management_cluster")),
		Name:              dcl.String(d.Get("name").(string)),
		NetworkConfig:     expandVmwareEnginePrivateCloudNetworkConfig(d.Get("network_config")),
		Description:       dcl.String(d.Get("description").(string)),
		Labels:            checkStringMap(d.Get("labels")),
		Project:           dcl.String(project),
	}

	id, err := replaceVarsForId(d, config, "projects/{{project}}/locations/{{location}}/privateClouds/{{name}}")
	if err != nil {
		return fmt.Errorf("error constructing id: %s", err)
	}
	d.SetId(id)
	createDirective := CreateDirective
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}
	billingProject := project
	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := NewDCLVmwareEngineClient(config, userAgent, billingProject)
	res, err := client.ApplyPrivateCloud(context.Background(), obj, createDirective...)

	if _, ok := err.(dcl.DiffAfterApplyError); ok {
		log.Printf("[DEBUG] Diff after apply returned from the DCL: %s", err)
	} else if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error creating PrivateCloud: %s", err)
	}

	log.Printf("[DEBUG] Finished creating PrivateCloud %q: %#v", d.Id(), res)

	return resourceVmwareEnginePrivateCloudRead(d, meta)
}

func resourceVmwareEnginePrivateCloudRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &vmwareengine.PrivateCloud{
		Location:          dcl.String(d.Get("location").(string)),
		ManagementCluster: expandVmwareEnginePrivateCloudManagementCluster(d.Get("management_cluster")),
		Name:              dcl.String(d.Get("name").(string)),
		NetworkConfig:     expandVmwareEnginePrivateCloudNetworkConfig(d.Get("network_config")),
		Description:       dcl.String(d.Get("description").(string)),
		Labels:            checkStringMap(d.Get("labels")),
		Project:           dcl.String(project),
	}

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}
	billingProject := project
	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := NewDCLVmwareEngineClient(config, userAgent, billingProject)
	res, err := client.GetPrivateCloud(context.Background(), obj)
	if err != nil {
		resourceName := fmt.Sprintf("VmwareEnginePrivateCloud %q", d.Id())
		return handleNotFoundDCLError(err, d, resourceName)
	}

	if err = d.Set("location", res.Location); err != nil {
		return fmt.Errorf("error setting location in state: %s", err)
	}
	if err = d.Set("management_cluster", flattenVmwareEnginePrivateCloudManagementCluster(res.ManagementCluster)); err != nil {
		return fmt.Errorf("error setting management_cluster in state: %s", err)
	}
	if err = d.Set("name", res.Name); err != nil {
		return fmt.Errorf("error setting name in state: %s", err)
	}
	if err = d.Set("network_config", flattenVmwareEnginePrivateCloudNetworkConfig(res.NetworkConfig)); err != nil {
		return fmt.Errorf("error setting network_config in state: %s", err)
	}
	if err = d.Set("description", res.Description); err != nil {
		return fmt.Errorf("error setting description in state: %s", err)
	}
	if err = d.Set("labels", res.Labels); err != nil {
		return fmt.Errorf("error setting labels in state: %s", err)
	}
	if err = d.Set("project", res.Project); err != nil {
		return fmt.Errorf("error setting project in state: %s", err)
	}
	if err = d.Set("conditions", flattenVmwareEnginePrivateCloudConditionsArray(res.Conditions)); err != nil {
		return fmt.Errorf("error setting conditions in state: %s", err)
	}
	if err = d.Set("create_time", res.CreateTime); err != nil {
		return fmt.Errorf("error setting create_time in state: %s", err)
	}
	if err = d.Set("delete_time", res.DeleteTime); err != nil {
		return fmt.Errorf("error setting delete_time in state: %s", err)
	}
	if err = d.Set("expire_time", res.ExpireTime); err != nil {
		return fmt.Errorf("error setting expire_time in state: %s", err)
	}
	if err = d.Set("hcx", flattenVmwareEnginePrivateCloudHcx(res.Hcx)); err != nil {
		return fmt.Errorf("error setting hcx in state: %s", err)
	}
	if err = d.Set("nsx", flattenVmwareEnginePrivateCloudNsx(res.Nsx)); err != nil {
		return fmt.Errorf("error setting nsx in state: %s", err)
	}
	if err = d.Set("state", res.State); err != nil {
		return fmt.Errorf("error setting state in state: %s", err)
	}
	if err = d.Set("update_time", res.UpdateTime); err != nil {
		return fmt.Errorf("error setting update_time in state: %s", err)
	}
	if err = d.Set("vcenter", flattenVmwareEnginePrivateCloudVcenter(res.Vcenter)); err != nil {
		return fmt.Errorf("error setting vcenter in state: %s", err)
	}

	return nil
}
func resourceVmwareEnginePrivateCloudUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &vmwareengine.PrivateCloud{
		Location:          dcl.String(d.Get("location").(string)),
		ManagementCluster: expandVmwareEnginePrivateCloudManagementCluster(d.Get("management_cluster")),
		Name:              dcl.String(d.Get("name").(string)),
		NetworkConfig:     expandVmwareEnginePrivateCloudNetworkConfig(d.Get("network_config")),
		Description:       dcl.String(d.Get("description").(string)),
		Labels:            checkStringMap(d.Get("labels")),
		Project:           dcl.String(project),
	}
	// Construct state hint from old values
	old := &vmwareengine.PrivateCloud{
		Location:          dcl.String(oldValue(d.GetChange("location")).(string)),
		ManagementCluster: expandVmwareEnginePrivateCloudManagementCluster(oldValue(d.GetChange("management_cluster"))),
		Name:              dcl.String(oldValue(d.GetChange("name")).(string)),
		NetworkConfig:     expandVmwareEnginePrivateCloudNetworkConfig(oldValue(d.GetChange("network_config"))),
		Description:       dcl.String(oldValue(d.GetChange("description")).(string)),
		Labels:            checkStringMap(oldValue(d.GetChange("labels"))),
		Project:           dcl.StringOrNil(oldValue(d.GetChange("project")).(string)),
	}
	directive := UpdateDirective
	directive = append(directive, dcl.WithStateHint(old))
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""
	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := NewDCLVmwareEngineClient(config, userAgent, billingProject)
	res, err := client.ApplyPrivateCloud(context.Background(), obj, directive...)

	if _, ok := err.(dcl.DiffAfterApplyError); ok {
		log.Printf("[DEBUG] Diff after apply returned from the DCL: %s", err)
	} else if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error updating PrivateCloud: %s", err)
	}

	log.Printf("[DEBUG] Finished creating PrivateCloud %q: %#v", d.Id(), res)

	return resourceVmwareEnginePrivateCloudRead(d, meta)
}

func resourceVmwareEnginePrivateCloudDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &vmwareengine.PrivateCloud{
		Location:          dcl.String(d.Get("location").(string)),
		ManagementCluster: expandVmwareEnginePrivateCloudManagementCluster(d.Get("management_cluster")),
		Name:              dcl.String(d.Get("name").(string)),
		NetworkConfig:     expandVmwareEnginePrivateCloudNetworkConfig(d.Get("network_config")),
		Description:       dcl.String(d.Get("description").(string)),
		Labels:            checkStringMap(d.Get("labels")),
		Project:           dcl.String(project),
	}

	log.Printf("[DEBUG] Deleting PrivateCloud %q", d.Id())
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}
	billingProject := project
	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := NewDCLVmwareEngineClient(config, userAgent, billingProject)
	if err := client.DeletePrivateCloud(context.Background(), obj); err != nil {
		return fmt.Errorf("Error deleting PrivateCloud: %s", err)
	}

	log.Printf("[DEBUG] Finished deleting PrivateCloud %q", d.Id())
	return nil
}

func resourceVmwareEnginePrivateCloudImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/privateClouds/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)",
		"(?P<location>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVarsForId(d, config, "projects/{{project}}/locations/{{location}}/privateClouds/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func expandVmwareEnginePrivateCloudManagementCluster(o interface{}) *vmwareengine.PrivateCloudManagementCluster {
	if o == nil {
		return vmwareengine.EmptyPrivateCloudManagementCluster
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return vmwareengine.EmptyPrivateCloudManagementCluster
	}
	obj := objArr[0].(map[string]interface{})
	return &vmwareengine.PrivateCloudManagementCluster{
		ClusterId:  dcl.String(obj["cluster_id"].(string)),
		NodeCount:  dcl.Int64(int64(obj["node_count"].(int))),
		NodeTypeId: dcl.String(obj["node_type_id"].(string)),
	}
}

func flattenVmwareEnginePrivateCloudManagementCluster(obj *vmwareengine.PrivateCloudManagementCluster) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"cluster_id":   obj.ClusterId,
		"node_count":   obj.NodeCount,
		"node_type_id": obj.NodeTypeId,
	}

	return []interface{}{transformed}

}

func expandVmwareEnginePrivateCloudNetworkConfig(o interface{}) *vmwareengine.PrivateCloudNetworkConfig {
	if o == nil {
		return vmwareengine.EmptyPrivateCloudNetworkConfig
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return vmwareengine.EmptyPrivateCloudNetworkConfig
	}
	obj := objArr[0].(map[string]interface{})
	return &vmwareengine.PrivateCloudNetworkConfig{
		ManagementCidr: dcl.String(obj["management_cidr"].(string)),
		Network:        dcl.String(obj["network"].(string)),
	}
}

func flattenVmwareEnginePrivateCloudNetworkConfig(obj *vmwareengine.PrivateCloudNetworkConfig) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"management_cidr": obj.ManagementCidr,
		"network":         obj.Network,
		"service_network": obj.ServiceNetwork,
	}

	return []interface{}{transformed}

}

func flattenVmwareEnginePrivateCloudConditionsArray(objs []vmwareengine.PrivateCloudConditions) []interface{} {
	if objs == nil {
		return nil
	}

	items := []interface{}{}
	for _, item := range objs {
		i := flattenVmwareEnginePrivateCloudConditions(&item)
		items = append(items, i)
	}

	return items
}

func flattenVmwareEnginePrivateCloudConditions(obj *vmwareengine.PrivateCloudConditions) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"code":    obj.Code,
		"message": obj.Message,
	}

	return transformed

}

func flattenVmwareEnginePrivateCloudHcx(obj *vmwareengine.PrivateCloudHcx) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"external_ip": obj.ExternalIP,
		"fdqn":        obj.Fdqn,
		"internal_ip": obj.InternalIP,
		"version":     obj.Version,
	}

	return []interface{}{transformed}

}

func flattenVmwareEnginePrivateCloudNsx(obj *vmwareengine.PrivateCloudNsx) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"external_ip": obj.ExternalIP,
		"fdqn":        obj.Fdqn,
		"internal_ip": obj.InternalIP,
		"version":     obj.Version,
	}

	return []interface{}{transformed}

}

func flattenVmwareEnginePrivateCloudVcenter(obj *vmwareengine.PrivateCloudVcenter) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"external_ip": obj.ExternalIP,
		"fdqn":        obj.Fdqn,
		"internal_ip": obj.InternalIP,
		"version":     obj.Version,
	}

	return []interface{}{transformed}

}
