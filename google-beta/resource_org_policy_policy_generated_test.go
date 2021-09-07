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
	orgpolicy "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/orgpolicy/beta"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"strings"
	"testing"
)

func TestAccOrgPolicyPolicy_OrganizationPolicy(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        getTestOrgFromEnv(t),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOrgPolicyPolicyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccOrgPolicyPolicy_OrganizationPolicy(context),
			},
			{
				ResourceName:            "google_org_policy_policy.primary",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name"},
			},
			{
				Config: testAccOrgPolicyPolicy_OrganizationPolicyUpdate0(context),
			},
			{
				ResourceName:            "google_org_policy_policy.primary",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name"},
			},
		},
	})
}

func testAccOrgPolicyPolicy_OrganizationPolicy(context map[string]interface{}) string {
	return Nprintf(`
resource "google_org_policy_policy" "primary" {
  name   = "organizations/%{org_id}/policies/gcp.detailedAuditLoggingMode"
  parent = "organizations/%{org_id}"

  spec {
    reset = true
  }
}


`, context)
}

func testAccOrgPolicyPolicy_OrganizationPolicyUpdate0(context map[string]interface{}) string {
	return Nprintf(`
resource "google_org_policy_policy" "primary" {
  name   = "organizations/%{org_id}/policies/gcp.detailedAuditLoggingMode"
  parent = "organizations/%{org_id}"

  spec {
    reset = false

    rules {
      enforce = true
    }
  }
}


`, context)
}

func testAccCheckOrgPolicyPolicyDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "rs.google_org_policy_policy" {
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

			obj := &orgpolicy.Policy{
				Name:   dcl.String(rs.Primary.Attributes["name"]),
				Parent: dcl.String(rs.Primary.Attributes["parent"]),
			}

			client := NewDCLOrgPolicyClient(config, config.userAgent, billingProject)
			_, err := client.GetPolicy(context.Background(), obj)
			if err == nil {
				return fmt.Errorf("google_org_policy_policy still exists %v", obj)
			}
		}
		return nil
	}
}
