package images

import (
	"github.com/ggiamarchi/gophercloud"
	os "github.com/ggiamarchi/gophercloud/openstack/compute/v2/images"
	"github.com/ggiamarchi/gophercloud/pagination"
)

// ListDetail enumerates the available server images.
func ListDetail(client *gophercloud.ServiceClient, opts os.ListOptsBuilder) pagination.Pager {
	return os.ListDetail(client, opts)
}

// Get acquires additional detail about a specific image by ID.
func Get(client *gophercloud.ServiceClient, id string) os.GetResult {
	return os.Get(client, id)
}

// ExtractImages interprets a page as a collection of server images.
func ExtractImages(page pagination.Page) ([]os.Image, error) {
	return os.ExtractImages(page)
}
