package extensions

import (
	"testing"

	"github.com/ggiamarchi/gophercloud"
	th "github.com/ggiamarchi/gophercloud/testhelper"
)

const endpoint = "http://localhost:57909/"

func endpointClient() *gophercloud.ServiceClient {
	return &gophercloud.ServiceClient{Endpoint: endpoint}
}

func TestExtensionURL(t *testing.T) {
	actual := ExtensionURL(endpointClient(), "agent")
	expected := endpoint + "extensions/agent"
	th.AssertEquals(t, expected, actual)
}

func TestListExtensionURL(t *testing.T) {
	actual := ListExtensionURL(endpointClient())
	expected := endpoint + "extensions"
	th.AssertEquals(t, expected, actual)
}
