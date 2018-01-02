package healthcheck

import (
	"github.com/gator1/huaweicloud"
//        "github.com/gophercloud/gophercloud"
)

const (
	rootPath     = "elbaas"
	resourcePath = "healthcheck"
)

func rootURL(c *huaweicloud.ServiceClient1) string {
	return c.ServiceURL(rootPath, resourcePath)
}

func resourceURL(c *huaweicloud.ServiceClient1, id string) string {
	return c.ServiceURL(rootPath, resourcePath, id)
}
