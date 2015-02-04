package users

import "github.com/ggiamarchi/gophercloud"

func resetAPIKeyURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("users", id, "OS-KSADM", "credentials", "RAX-KSKEY:apiKeyCredentials", "RAX-AUTH", "reset")
}
