---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: DCL     ***
#
# ----------------------------------------------------------------------------
#
#     This file is managed by Magic Modules (https:#github.com/GoogleCloudPlatform/magic-modules)
#     and is based on the DCL (https:#github.com/GoogleCloudPlatform/declarative-resource-client-library).
#     Changes will need to be made to the DCL or Magic Modules instead of here.
#
#     We are not currently able to accept contributions to this file. If changes
#     are required, please file an issue at https:#github.com/hashicorp/terraform-provider-google/issues/new/choose
#
# ----------------------------------------------------------------------------
subcategory: "VmwareEngine"
layout: "google"
page_title: "Google: google_vmware_engine_private_cloud"
sidebar_current: "docs-google-vmware-engine-private-cloud"
description: |-
Beta only
---

# google\_vmware\_engine\_private\_cloud

Beta only
## Example Usage - basic_private_cloud
A basic example of a vmware engine private cloud
```hcl
resource "google_vmware_engine_private_cloud" "primary" {
  location = "{{zone}}"

  management_cluster {
    cluster_id   = "{{management}}"
    node_count   = 4
    node_type_id = "standard-72"
  }

  name = "{{cloud}}"

  network_config {
    management_cidr = "192.168.0.0/24"
    network         = "projects/my-project-name/global/networks/default"
  }

  description = "A sample private cloud"
  project     = "my-project-name"
}


```
## Example Usage - minimal_private_cloud
A minimal example of a vmware engine private cloud
```hcl
resource "google_vmware_engine_private_cloud" "primary" {
  location = "{{zone}}"

  management_cluster {
    cluster_id   = "{{management}}"
    node_count   = 3
    node_type_id = "standard-72"
  }

  name = "{{cloud}}"

  network_config {
    management_cidr = "192.168.1.0/24"
    network         = "projects/my-project-name/global/networks/default"
  }

  project = "my-project-name"
}


```

## Argument Reference

The following arguments are supported:

* `location` -
  (Required)
  The location for the resource
  
* `management_cluster` -
  (Required)
  Input only. The management cluster for this private cloud. This parameter is required during creation of private cloud to provide details for the default cluster.
  
* `name` -
  (Required)
  The resource name of this private cloud. Resource names are schemeless URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names. For example: `projects/my-project/locations/us-west1-a/privateClouds/my-cloud`
  
* `network_config` -
  (Required)
  Required. Network configuration.
  


The `management_cluster` block supports:
    
* `cluster_id` -
  (Required)
  Required. The user-provided identifier of the new `Cluster`.
    
* `node_count` -
  (Required)
  Required. Number of nodes in this cluster.
    
* `node_type_id` -
  (Required)
  Required. The canonical identifier of node types (`NodeType`) in this cluster. For example: standard-72.
    
The `network_config` block supports:
    
* `management_cidr` -
  (Required)
  Required. Management CIDR used by VMWare management appliances.
    
* `network` -
  (Required)
  Required. The relative resource name of the consumer VPC network this private cloud is attached to. Specify the name in the following form: `projects/{project}/global/networks/{network_id}` where `{project}` can either be a project number or a project ID.
    
* `service_network` -
  Output only. The relative resource name of the service VPC network this private cloud is attached to. The name is specified in the following form: `projects/{service_project_number}/global/networks/{network_id}`.
    
- - -

* `description` -
  (Optional)
  User-provided description for this private cloud.
  
* `labels` -
  (Optional)
  Labels are a way to attach lightweight metadata to resources for filtering and querying resource data. No more than 64 user labels can be associated with each resource. Label keys and values can be no longer than 63 characters, can only contain lowercase letters, numeric characters, underscores and dashes, where label keys must start with a letter and international characters are allowed. The empty string is a valid value. Labels are set on creation and updated like any other field. Specifically, to add a new label, you would need to provide all of the existing labels along with the new label. If you only provide a map with the new label, all of the old labels will be removed (probably not what is desired).
  
* `project` -
  (Optional)
  The project for the resource
  


## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `projects/{{project}}/locations/{{location}}/privateClouds/{{name}}`

* `conditions` -
  Output only. The conditions that caused the current private cloud state. For example, private cloud provisioning failure description.
  
* `create_time` -
  Output only. Creation time of this resource in RFC3339 text format.
  
* `delete_time` -
  Output only. Time the resource was marked as deleted, in RFC3339 text format.
  
* `expire_time` -
  Output only. Planned deletion time of this resource in RFC3339 text format.
  
* `hcx` -
  Output only. HCX appliance.
  
* `nsx` -
  Output only. NSX appliance.
  
* `state` -
  Output only. State of the resource. Possible values: ACTIVE, CREATING, UPDATING, FAILED, DELETED
  
* `update_time` -
  Output only. Last update time of this resource in RFC3339 text format.
  
* `vcenter` -
  Output only. Vcenter appliance.
  
## Timeouts

This resource provides the following
[Timeouts](/docs/configuration/resources.html#timeouts) configuration options:

- `create` - Default is 10 minutes.
- `update` - Default is 10 minutes.
- `delete` - Default is 10 minutes.

## Import

PrivateCloud can be imported using any of these accepted formats:

```
$ terraform import google_vmware_engine_private_cloud.default projects/{{project}}/locations/{{location}}/privateClouds/{{name}}
$ terraform import google_vmware_engine_private_cloud.default {{project}}/{{location}}/{{name}}
$ terraform import google_vmware_engine_private_cloud.default {{location}}/{{name}}
```



