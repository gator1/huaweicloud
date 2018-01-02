package backendmember


import (
	"github.com/gator1/huaweicloud"
        // "github.com/gophercloud/gophercloud"
)
const (
	rootPath     = "elbaas"
	resourcePath = "listeners"
)

func addURL(c *huaweicloud.ServiceClient1, listener_id string) string {
	return c.ServiceURL(rootPath, resourcePath, listener_id, "members")
}

func removeURL(c *huaweicloud.ServiceClient1, listener_id string) string {
	return c.ServiceURL(rootPath, resourcePath, listener_id, "members", "action")
}

func resourceURL(c *huaweicloud.ServiceClient1, listener_id string, id string) string {
	return c.ServiceURL(rootPath, resourcePath, listener_id, "members?id=" + id)
}
