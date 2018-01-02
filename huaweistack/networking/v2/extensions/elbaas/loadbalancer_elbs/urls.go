package loadbalancer_elbs

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gator1/huaweicloud"
)
const (
	rootPath     = "elbaas"
	resourcePath = "loadbalancers"
)

func rootURLh(c *huaweicloud.ServiceClient1) string {
	return c.ServiceURL(rootPath, resourcePath)
}

func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(rootPath, resourcePath)
}

func resourceURL(c *huaweicloud.ServiceClient1, id string) string {
	return c.ServiceURL(rootPath, resourcePath, id)
}
