package agentregistry_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
)

func TestAccAgentRegistryMcpServer_update(t *testing.T) {
	t.Parallel()

	randomSuffix := acctest.RandString(t, 10)
	mcpServerId := "tf-test-mcp-server-" + randomSuffix

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckAgentRegistryMcpServerDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccAgentRegistryMcpServer_basic(mcpServerId, randomSuffix, "My MCP Server", "An MCP Server registered in Agent Registry"),
			},
			{
				ResourceName:            "google_agent_registry_mcp_server.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "mcp_server_id"},
			},
			{
				Config: testAccAgentRegistryMcpServer_basic(mcpServerId, randomSuffix, "Updated MCP Server", "An updated description"),
			},
			{
				ResourceName:            "google_agent_registry_mcp_server.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "mcp_server_id"},
			},
		},
	})
}

func testAccAgentRegistryMcpServer_basic(mcpServerId, suffix, displayName, description string) string {
	return fmt.Sprintf(`
resource "google_agent_registry_mcp_server" "example" {
  provider      = google-beta
  mcp_server_id = "%s"
  location      = "global"
  display_name  = "%s"
  description   = "%s"

  interfaces {
    url              = "https://mymcp-%s.com"
    protocol_binding = "JSONRPC"
  }

  mcp_server_spec {
    type    = "NO_SPEC"
  }
}
`, mcpServerId, displayName, description, suffix)
}
