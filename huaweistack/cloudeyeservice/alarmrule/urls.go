package alarmrule

import (
	//"github.com/gophercloud/gophercloud"
	"github.com/gator1/huaweicloud"
)


const (
	rootPath = "alarms"
)

func rootURL(c *huaweicloud.ServiceClient1) string {
	return c.ServiceURL(c.ProjectID, rootPath)
}

func resourceURL(c *huaweicloud.ServiceClient1, id string) string {
	return c.ServiceURL(c.ProjectID, rootPath, id)
}

func actionURL(c *huaweicloud.ServiceClient1, id string) string {
	return c.ServiceURL(c.ProjectID, rootPath, id, "action")
}
