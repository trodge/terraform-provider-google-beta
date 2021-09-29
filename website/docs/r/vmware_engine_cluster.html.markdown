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
page_title: "Google: google_vmware_engine_cluster"
sidebar_current: "docs-google-vmware-engine-cluster"
description: |-
Beta only
---

# google\_vmware\_engine\_cluster

Beta only
## Example Usage - basic_cluster
This test creates a private cloud, creates a cluster, then updates the cloud and cluster.
```hcl
resource "google_vmware_engine_cluster" "primary" {
  location      = "{{zone}}"
  node_count    = 3
  node_type_id  = "standard-72"
  private_cloud = google_vmware_engine_private_cloud.basic.name
  project       = "{{project}}"
}

resource "google_vmware_engine_private_cloud" "basic" {
  location = "{{zone}}"

  management_cluster {
    cluster_id   = "{{management}}"
    node_count   = 4
    node_type_id = "standard-72"
  }

  name = "{{cloud}}"

  network_config {
    management_cidr = "192.168.0.0/24"
    network         = "projects/{{project}}/global/networks/default"
  }

  description = "A sample private cloud"
  project     = "{{project}}"
}


```
## Example Usage - minimal_cluster
A minimal example of a private cloud and cluster
```hcl
resource "google_vmware_engine_cluster" "primary" {
  location      = "{{zone}}"
  node_count    = 3
  node_type_id  = "standard-72"
  private_cloud = google_vmware_engine_private_cloud.minimal.name
  project       = "{{project}}"
}

resource "google_vmware_engine_private_cloud" "minimal" {
  location = "{{zone}}"

  management_cluster {
    cluster_id   = "{{management}}"
    node_count   = 3
    node_type_id = "standard-72"
  }

  name = "{{cloud}}"

  network_config {
    management_cidr = "192.168.1.0/24"
    network         = "projects/{{project}}/global/networks/default"
  }

  project = "{{project}}"
}


```

## Argument Reference

The following arguments are supported:

* `location` -
  (Required)
  The location for the resource
  
* `node_count` -
  (Required)
  Required. Number of nodes in this cluster.
  
* `node_type_id` -
  (Required)
  Required. The canonical identifier of node types (`NodeType`) in this cluster. For example: standard-72.
  
* `private_cloud` -
  (Required)
  The privateCloud for the resource
  


- - -

* `labels` -
  (Optional)
  Labels are a way to attach lightweight metadata to resources for filtering and querying resource data. No more than 64 user labels can be associated with each resource. Label keys and values can be no longer than 63 characters, can only contain lowercase letters, numeric characters, underscores and dashes, where label keys must start with a letter and international characters are allowed. The empty string is a valid value. Labels are set on creation and updated like any other field. Specifically, to add a new label, you would need to provide all of the existing labels along with the new label. If you only provide a map with the new label, all of the old labels will be removed (probably not what is desired).
  
* `project` -
  (Optional)
  The project for the resource
  


## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `projects/{{project}}/locations/{{location}}/privateClouds/{{private_cloud}}/clusters/{{name}}`

* `create_time` -
  Output only. Creation time of this resource in RFC3339 text format.
  
* `management` -
  Output only. True if the cluster is a management cluster; false otherwise. There can only be one management cluster in a private cloud and it has to be the first one.
  
* `name` -
  Output only. The resource name of this cluster. Resource names are schemeless URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names. For example: `projects/my-project/locations/us-west1-a/privateClouds/my-cloud/clusters/my-cluster`
  
* `state` -
  Output only. State of the resource. Possible values: STATE_UNSPECIFIED, ACTIVE, CREATING, UPDATING, FAILED, DELETED
  
* `update_time` -
  Output only. Last update time of this resource in RFC3339 text format.
  
## Timeouts

This resource provides the following
[Timeouts](/docs/configuration/resources.html#timeouts) configuration options:

- `create` - Default is 10 minutes.
- `update` - Default is 10 minutes.
- `delete` - Default is 10 minutes.

## Import

Cluster can be imported using any of these accepted formats:

```
$ terraform import google_vmware_engine_cluster.default projects/{{project}}/locations/{{location}}/privateClouds/{{private_cloud}}/clusters/{{name}}
$ terraform import google_vmware_engine_cluster.default {{project}}/{{location}}/{{private_cloud}}/{{name}}
$ terraform import google_vmware_engine_cluster.default {{location}}/{{private_cloud}}/{{name}}
```



