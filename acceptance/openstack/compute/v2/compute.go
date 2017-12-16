package v2

import (
	//"fmt"
	"testing"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/tags"
)

func CreateServerTags(t *testing.T, client *gophercloud.ServiceClient, server_id string, taglist []string) (*tags.Tags, error) {
	createOpts := tags.CreateOpts{
		Tags: taglist,
	}
	return tags.Create(client, server_id, createOpts).Extract()
}

func DeleteServerTags(t *testing.T, client *gophercloud.ServiceClient, server_id string) error {
	return tags.Delete(client, server_id).ExtractErr()
}

func GetServerTags(t *testing.T, client *gophercloud.ServiceClient, server_id string) (*tags.Tags, error) {
	return tags.Get(client, server_id).Extract()
}


