package common

import (
	"github.com/ggiamarchi/gophercloud"
	"github.com/ggiamarchi/gophercloud/testhelper/client"
)

const TokenID = client.TokenID

func ServiceClient() *gophercloud.ServiceClient {
	return client.ServiceClient()
}
