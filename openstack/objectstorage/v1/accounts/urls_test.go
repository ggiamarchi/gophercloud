package accounts

import (
	"testing"

	"github.com/ggiamarchi/gophercloud"
	th "github.com/ggiamarchi/gophercloud/testhelper"
)

const endpoint = "http://localhost:57909/"

func endpointClient() *gophercloud.ServiceClient {
	return &gophercloud.ServiceClient{Endpoint: endpoint}
}

func TestGetURL(t *testing.T) {
	actual := getURL(endpointClient())
	expected := endpoint
	th.CheckEquals(t, expected, actual)
}

func TestUpdateURL(t *testing.T) {
	actual := updateURL(endpointClient())
	expected := endpoint
	th.CheckEquals(t, expected, actual)
}
