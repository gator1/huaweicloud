package v2

import (
	"testing"

	"github.com/gophercloud/gophercloud/acceptance/clients"
	"github.com/gophercloud/gophercloud/acceptance/tools"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v2/volumes"
	"github.com/gophercloud/gophercloud/acceptance/openstack/blockstorage/v2"
	"reflect"
)

func TestVolumesTags(t *testing.T) {
	client, err := clients.NewBlockStorageV2Client()
	if err != nil {
		t.Fatalf("Unable to create blockstorage client: %v", err)
	}

	volume, err := v2.CreateVolume(t, client)
	if err != nil {
		t.Fatalf("Unable to create volume: %v", err)
	}
	defer v2.DeleteVolume(t, client, volume)

	tagmap0, err := v2.GetVolumeTags(t, client, "volumes", volume.ID)
	if err != nil {
		t.Errorf("Unable to get initial tags from volume: %v", err)
	}
	tools.PrintResource(t, tagmap0)

	tagmap := map[string]string{"foo" : "bar", "name" : "value"}
	tagmap2, err := v2.CreateVolumeTags(t, client, "volumes", volume.ID, tagmap)
	if err != nil {
		t.Errorf("Unable to create tags for volume: %v", err)
	}
	tagmap3, err := v2.GetVolumeTags(t, client, "volumes", volume.ID)
	if err != nil {
		t.Errorf("Unable to get tags from volume: %v", err)
	}
	if !reflect.DeepEqual(tagmap3.Tags, tagmap) {
		t.Errorf("Tags aren't equal after set/get: %v != %v", tagmap3.Tags, tagmap)
	}
	tools.PrintResource(t, tagmap2)

	tagmap4 := map[string]string{"foo2" : "bar2", "name2" : "value2"}
	tagmap5, err := v2.CreateVolumeTags(t, client, "volumes", volume.ID, tagmap4)
	if err != nil {
		t.Errorf("Unable to create tags for volume: %v", err)
	}
	tagmap6, err := v2.GetVolumeTags(t, client, "volumes", volume.ID)
	if err != nil {
		t.Errorf("Unable to get tags from volume: %v", err)
	}
	if !reflect.DeepEqual(tagmap6.Tags, tagmap4) {
		t.Errorf("Tags aren't equal after set/get: %v != %v", tagmap6.Tags, tagmap4)
	}
	tools.PrintResource(t, tagmap5)

	tagmap0a := map[string]string{}
	err = v2.DeleteVolumeTags(t, client, "volumes", volume.ID)
	if err != nil {
		t.Errorf("Unable to delete tags for volume: %v", err)
	}
	tagmap0c, err := v2.GetVolumeTags(t, client, "volumes", volume.ID)
	if err != nil {
		t.Errorf("Unable to get empty tags from volume: %v", err)
	}
	if !reflect.DeepEqual(tagmap0c.Tags, tagmap0a) {
		t.Errorf("Tags aren't equal after set/get: %v != %v", tagmap0c.Tags, tagmap0a)
	}
	tools.PrintResource(t, tagmap0a)

	//err = DeleteVolumeTags(t, client, "volumes", volume.ID)
	newVolume, err := volumes.Get(client, volume.ID).Extract()
	if err != nil {
		t.Errorf("Unable to retrieve volume: %v", err)
	}

	tools.PrintResource(t, newVolume)
}


