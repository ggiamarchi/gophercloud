package flavors

import (
	"github.com/ggiamarchi/gophercloud"

	os "github.com/ggiamarchi/gophercloud/openstack/cdn/v1/flavors"
	"github.com/ggiamarchi/gophercloud/pagination"
)

// List returns a single page of CDN flavors.
func List(c *gophercloud.ServiceClient) pagination.Pager {
	return os.List(c)
}

// Get retrieves a specific flavor based on its unique ID.
func Get(c *gophercloud.ServiceClient, id string) os.GetResult {
	return os.Get(c, id)
}
