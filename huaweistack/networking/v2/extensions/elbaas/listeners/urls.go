package listeners

import (
	"github.com/gator1/huaweicloud"
         "github.com/gophercloud/gophercloud"
)

const (
	rootPath     = "elbaas"
	resourcePath = "listeners"
)

func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(rootPath, resourcePath)
}

func rootURLh(c *huaweicloud.ServiceClient1) string {
	return c.ServiceURL(rootPath, resourcePath)
}

func resourceURL(c *huaweicloud.ServiceClient1, id string) string {
	return c.ServiceURL(rootPath, resourcePath, id)
}
