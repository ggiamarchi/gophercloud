package tenants

import (
	"github.com/ggiamarchi/gophercloud"
	os "github.com/ggiamarchi/gophercloud/openstack/identity/v2/tenants"
	"github.com/ggiamarchi/gophercloud/pagination"
)

// ExtractTenants interprets a page of List results as a more usable slice of Tenant structs.
func ExtractTenants(page pagination.Page) ([]os.Tenant, error) {
	return os.ExtractTenants(page)
}

// List enumerates the tenants to which the current token grants access.
func List(client *gophercloud.ServiceClient, opts *os.ListOpts) pagination.Pager {
	return os.List(client, opts)
}
