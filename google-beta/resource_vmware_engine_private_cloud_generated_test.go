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

func TestAccVmwareEnginePrivateCloud_BasicPrivateCloud(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_name":  getTestProjectFromEnv(),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },

		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckVmwareEnginePrivateCloudDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccVmwareEnginePrivateCloud_BasicPrivateCloud(context),
			},
			{
				ResourceName:      "google_vmware_engine_private_cloud.primary",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccVmwareEnginePrivateCloud_BasicPrivateCloudUpdate0(context),
			},
			{
				ResourceName:      "google_vmware_engine_private_cloud.primary",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
func TestAccVmwareEnginePrivateCloud_MinimalPrivateCloud(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_name":  getTestProjectFromEnv(),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },

		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckVmwareEnginePrivateCloudDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccVmwareEnginePrivateCloud_MinimalPrivateCloud(context),
			},
			{
				ResourceName:      "google_vmware_engine_private_cloud.primary",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccVmwareEnginePrivateCloud_BasicPrivateCloud(context map[string]interface{}) string {
	return Nprintf(`
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
    network         = "projects/%{project_name}/global/networks/default"
  }

  description = "A sample private cloud"
  project     = "%{project_name}"
}


`, context)
}

func testAccVmwareEnginePrivateCloud_BasicPrivateCloudUpdate0(context map[string]interface{}) string {
	return Nprintf(`
resource "google_vmware_engine_private_cloud" "primary" {
  location = "{{zone}}"

  management_cluster {
    cluster_id   = "{{management}}"
    node_count   = 3
    node_type_id = "standard-72"
  }

  name = "{{cloud}}"

  network_config {
    management_cidr = "192.168.0.0/24"
    network         = "projects/%{project_name}/global/networks/default"
  }

  description = "An updated sample private cloud"
  project     = "%{project_name}"
}


`, context)
}

func testAccVmwareEnginePrivateCloud_MinimalPrivateCloud(context map[string]interface{}) string {
	return Nprintf(`
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
    network         = "projects/%{project_name}/global/networks/default"
  }

  project = "%{project_name}"
}


`, context)
}

func testAccCheckVmwareEnginePrivateCloudDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "rs.google_vmware_engine_private_cloud" {
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

			obj := &vmwareengine.PrivateCloud{
				Location:    dcl.String(rs.Primary.Attributes["location"]),
				Name:        dcl.String(rs.Primary.Attributes["name"]),
				Description: dcl.String(rs.Primary.Attributes["description"]),
				Project:     dcl.StringOrNil(rs.Primary.Attributes["project"]),
				CreateTime:  dcl.StringOrNil(rs.Primary.Attributes["create_time"]),
				DeleteTime:  dcl.StringOrNil(rs.Primary.Attributes["delete_time"]),
				ExpireTime:  dcl.StringOrNil(rs.Primary.Attributes["expire_time"]),
				State:       vmwareengine.PrivateCloudStateEnumRef(rs.Primary.Attributes["state"]),
				UpdateTime:  dcl.StringOrNil(rs.Primary.Attributes["update_time"]),
			}

			client := NewDCLVmwareEngineClient(config, config.userAgent, billingProject)
			_, err := client.GetPrivateCloud(context.Background(), obj)
			if err == nil {
				return fmt.Errorf("google_vmware_engine_private_cloud still exists %v", obj)
			}
		}
		return nil
	}
}
