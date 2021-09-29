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

func resourceVmwareEngineCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceVmwareEngineClusterCreate,
		Read:   resourceVmwareEngineClusterRead,
		Update: resourceVmwareEngineClusterUpdate,
		Delete: resourceVmwareEngineClusterDelete,

		Importer: &schema.ResourceImporter{
			State: resourceVmwareEngineClusterImport,
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

			"node_count": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Required. Number of nodes in this cluster.",
			},

			"node_type_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Required. The canonical identifier of node types (`NodeType`) in this cluster. For example: standard-72.",
			},

			"private_cloud": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      "The privateCloud for the resource",
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

			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Output only. Creation time of this resource in RFC3339 text format.",
			},

			"management": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Output only. True if the cluster is a management cluster; false otherwise. There can only be one management cluster in a private cloud and it has to be the first one.",
			},

			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Output only. The resource name of this cluster. Resource names are schemeless URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names. For example: `projects/my-project/locations/us-west1-a/privateClouds/my-cloud/clusters/my-cluster`",
			},

			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Output only. State of the resource. Possible values: STATE_UNSPECIFIED, ACTIVE, CREATING, UPDATING, FAILED, DELETED",
			},

			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Output only. Last update time of this resource in RFC3339 text format.",
			},
		},
	}
}

func resourceVmwareEngineClusterCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &vmwareengine.Cluster{
		Location:     dcl.String(d.Get("location").(string)),
		NodeCount:    dcl.Int64(int64(d.Get("node_count").(int))),
		NodeTypeId:   dcl.String(d.Get("node_type_id").(string)),
		PrivateCloud: dcl.String(d.Get("private_cloud").(string)),
		Labels:       checkStringMap(d.Get("labels")),
		Project:      dcl.String(project),
	}

	id, err := replaceVarsForId(d, config, "projects/{{project}}/locations/{{location}}/privateClouds/{{private_cloud}}/clusters/{{name}}")
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
	res, err := client.ApplyCluster(context.Background(), obj, createDirective...)

	if _, ok := err.(dcl.DiffAfterApplyError); ok {
		log.Printf("[DEBUG] Diff after apply returned from the DCL: %s", err)
	} else if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error creating Cluster: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Cluster %q: %#v", d.Id(), res)

	if err = d.Set("name", res.Name); err != nil {
		return fmt.Errorf("error setting name in state: %s", err)
	}
	// Id has a server-generated value, set again after creation
	id, err = replaceVarsForId(d, config, "projects/{{project}}/locations/{{location}}/privateClouds/{{private_cloud}}/clusters/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return resourceVmwareEngineClusterRead(d, meta)
}

func resourceVmwareEngineClusterRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &vmwareengine.Cluster{
		Location:     dcl.String(d.Get("location").(string)),
		NodeCount:    dcl.Int64(int64(d.Get("node_count").(int))),
		NodeTypeId:   dcl.String(d.Get("node_type_id").(string)),
		PrivateCloud: dcl.String(d.Get("private_cloud").(string)),
		Labels:       checkStringMap(d.Get("labels")),
		Project:      dcl.String(project),
		Name:         dcl.StringOrNil(d.Get("name").(string)),
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
	res, err := client.GetCluster(context.Background(), obj)
	if err != nil {
		resourceName := fmt.Sprintf("VmwareEngineCluster %q", d.Id())
		return handleNotFoundDCLError(err, d, resourceName)
	}

	if err = d.Set("location", res.Location); err != nil {
		return fmt.Errorf("error setting location in state: %s", err)
	}
	if err = d.Set("node_count", res.NodeCount); err != nil {
		return fmt.Errorf("error setting node_count in state: %s", err)
	}
	if err = d.Set("node_type_id", res.NodeTypeId); err != nil {
		return fmt.Errorf("error setting node_type_id in state: %s", err)
	}
	if err = d.Set("private_cloud", res.PrivateCloud); err != nil {
		return fmt.Errorf("error setting private_cloud in state: %s", err)
	}
	if err = d.Set("labels", res.Labels); err != nil {
		return fmt.Errorf("error setting labels in state: %s", err)
	}
	if err = d.Set("project", res.Project); err != nil {
		return fmt.Errorf("error setting project in state: %s", err)
	}
	if err = d.Set("create_time", res.CreateTime); err != nil {
		return fmt.Errorf("error setting create_time in state: %s", err)
	}
	if err = d.Set("management", res.Management); err != nil {
		return fmt.Errorf("error setting management in state: %s", err)
	}
	if err = d.Set("name", res.Name); err != nil {
		return fmt.Errorf("error setting name in state: %s", err)
	}
	if err = d.Set("state", res.State); err != nil {
		return fmt.Errorf("error setting state in state: %s", err)
	}
	if err = d.Set("update_time", res.UpdateTime); err != nil {
		return fmt.Errorf("error setting update_time in state: %s", err)
	}

	return nil
}
func resourceVmwareEngineClusterUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &vmwareengine.Cluster{
		Location:     dcl.String(d.Get("location").(string)),
		NodeCount:    dcl.Int64(int64(d.Get("node_count").(int))),
		NodeTypeId:   dcl.String(d.Get("node_type_id").(string)),
		PrivateCloud: dcl.String(d.Get("private_cloud").(string)),
		Labels:       checkStringMap(d.Get("labels")),
		Project:      dcl.String(project),
		Name:         dcl.StringOrNil(d.Get("name").(string)),
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
	client := NewDCLVmwareEngineClient(config, userAgent, billingProject)
	res, err := client.ApplyCluster(context.Background(), obj, directive...)

	if _, ok := err.(dcl.DiffAfterApplyError); ok {
		log.Printf("[DEBUG] Diff after apply returned from the DCL: %s", err)
	} else if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error updating Cluster: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Cluster %q: %#v", d.Id(), res)

	return resourceVmwareEngineClusterRead(d, meta)
}

func resourceVmwareEngineClusterDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &vmwareengine.Cluster{
		Location:     dcl.String(d.Get("location").(string)),
		NodeCount:    dcl.Int64(int64(d.Get("node_count").(int))),
		NodeTypeId:   dcl.String(d.Get("node_type_id").(string)),
		PrivateCloud: dcl.String(d.Get("private_cloud").(string)),
		Labels:       checkStringMap(d.Get("labels")),
		Project:      dcl.String(project),
		Name:         dcl.StringOrNil(d.Get("name").(string)),
	}

	log.Printf("[DEBUG] Deleting Cluster %q", d.Id())
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
	if err := client.DeleteCluster(context.Background(), obj); err != nil {
		return fmt.Errorf("Error deleting Cluster: %s", err)
	}

	log.Printf("[DEBUG] Finished deleting Cluster %q", d.Id())
	return nil
}

func resourceVmwareEngineClusterImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/privateClouds/(?P<private_cloud>[^/]+)/clusters/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<private_cloud>[^/]+)/(?P<name>[^/]+)",
		"(?P<location>[^/]+)/(?P<private_cloud>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVarsForId(d, config, "projects/{{project}}/locations/{{location}}/privateClouds/{{private_cloud}}/clusters/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}
