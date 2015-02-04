package flavors

import (
	"github.com/racker/perigee"
	"github.com/ggiamarchi/gophercloud"
	"github.com/ggiamarchi/gophercloud/pagination"
)

// List returns a single page of CDN flavors.
func List(c *gophercloud.ServiceClient) pagination.Pager {
	url := listURL(c)
	createPage := func(r pagination.PageResult) pagination.Page {
		return FlavorPage{pagination.SinglePageBase(r)}
	}
	return pagination.NewPager(c, url, createPage)
}

// Get retrieves a specific flavor based on its unique ID.
func Get(c *gophercloud.ServiceClient, id string) GetResult {
	var res GetResult
	_, res.Err = perigee.Request("GET", getURL(c, id), perigee.Options{
		MoreHeaders: c.AuthenticatedHeaders(),
		Results:     &res.Body,
		OkCodes:     []int{200},
	})
	return res
}
