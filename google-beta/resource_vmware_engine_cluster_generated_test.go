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
	dcl "github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	vmwareengine "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/vmwareengine/beta"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"strings"
	"testing"
)

func TestAccVmwareEngineCluster_BasicCluster(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },

		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckVmwareEngineClusterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccVmwareEngineCluster_BasicCluster(context),
			},
			{
				ResourceName:      "google_vmware_engine_cluster.primary",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccVmwareEngineCluster_BasicClusterUpdate0(context),
			},
			{
				ResourceName:      "google_vmware_engine_cluster.primary",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
func TestAccVmwareEngineCluster_MinimalCluster(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },

		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckVmwareEngineClusterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccVmwareEngineCluster_MinimalCluster(context),
			},
			{
				ResourceName:      "google_vmware_engine_cluster.primary",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccVmwareEngineCluster_BasicCluster(context map[string]interface{}) string {
	return Nprintf(`
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


`, context)
}

func testAccVmwareEngineCluster_BasicClusterUpdate0(context map[string]interface{}) string {
	return Nprintf(`
resource "google_vmware_engine_cluster" "primary" {
  location      = "{{zone}}"
  node_count    = 4
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


`, context)
}

func testAccVmwareEngineCluster_MinimalCluster(context map[string]interface{}) string {
	return Nprintf(`
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


`, context)
}

func testAccCheckVmwareEngineClusterDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "rs.google_vmware_engine_cluster" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			billingProject := ""
			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			obj := &vmwareengine.Cluster{
				Location:     dcl.String(rs.Primary.Attributes["location"]),
				NodeTypeId:   dcl.String(rs.Primary.Attributes["node_type_id"]),
				PrivateCloud: dcl.String(rs.Primary.Attributes["private_cloud"]),
				Project:      dcl.StringOrNil(rs.Primary.Attributes["project"]),
				CreateTime:   dcl.StringOrNil(rs.Primary.Attributes["create_time"]),
				Management:   dcl.Bool(rs.Primary.Attributes["management"] == "true"),
				Name:         dcl.StringOrNil(rs.Primary.Attributes["name"]),
				State:        vmwareengine.ClusterStateEnumRef(rs.Primary.Attributes["state"]),
				UpdateTime:   dcl.StringOrNil(rs.Primary.Attributes["update_time"]),
			}

			client := NewDCLVmwareEngineClient(config, config.userAgent, billingProject)
			_, err := client.GetCluster(context.Background(), obj)
			if err == nil {
				return fmt.Errorf("google_vmware_engine_cluster still exists %v", obj)
			}
		}
		return nil
	}
}
