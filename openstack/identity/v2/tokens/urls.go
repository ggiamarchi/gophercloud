package tokens

import "github.com/ggiamarchi/gophercloud"

// CreateURL generates the URL used to create new Tokens.
func CreateURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("tokens")
}
