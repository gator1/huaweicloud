package huaweistack

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gator1/huaweicloud"
	"github.com/gophercloud/gophercloud/openstack"
	tokens2 "github.com/gophercloud/gophercloud/openstack/identity/v2/tokens"
	tokens3 "github.com/gophercloud/gophercloud/openstack/identity/v3/tokens"
	"github.com/gophercloud/gophercloud/openstack/utils"
	"strings"
)

const (
        v20 = "v2.0"
        v30 = "v3.0"
)

func GetProjectId(client *gophercloud.ProviderClient) (string, error) {
	versions := []*utils.Version{
		{ID: v20, Priority: 20, Suffix: "/v2.0/"},
		{ID: v30, Priority: 30, Suffix: "/v3/"},
	}

	chosen, endpoint, err := utils.ChooseVersion(client, versions)
	if err != nil {
		return "", err
	}

	switch chosen.ID {
	case v20:
		return getV2ProjectId(client, endpoint)
	case v30:
		return getV3ProjectId(client, endpoint)
	default:
		return "", fmt.Errorf("Unrecognized identity version: %s", chosen.ID)
	}
}

func getV2ProjectId(client *gophercloud.ProviderClient, endpoint string) (string, error) {
	v2Client, err := openstack.NewIdentityV2(client, gophercloud.EndpointOpts{})
	if err != nil {
		return "", err
	}

	if endpoint != "" {
		v2Client.Endpoint = endpoint
	}

	result := tokens2.Get(v2Client, client.TokenID)
	token, err := result.ExtractToken()
	if err != nil {
		return "", err
	}

	return token.Tenant.ID, nil
}

func getV3ProjectId(client *gophercloud.ProviderClient, endpoint string) (string, error) {
	v3Client, err := openstack.NewIdentityV3(client, gophercloud.EndpointOpts{})
	if err != nil {
		return "", err
	}

	if endpoint != "" {
		v3Client.Endpoint = endpoint
	}

	result := tokens3.Get(v3Client, client.TokenID)
	project, err := result.ExtractProject()
	if err != nil {
		return "", err
	}

	return project.ID, nil
}

func initClientOpts(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts, clientType string) (*gophercloud.ServiceClient, error) {
	sc := new(gophercloud.ServiceClient)
	eo.ApplyDefaults(clientType)
	url, err := client.EndpointLocator(eo)
	if err != nil {
		return sc, err
	}
	sc.ProviderClient = client
	sc.Endpoint = url
	sc.Type = clientType
	return sc, nil
}

func initClientOpts1(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts, clientType string) (*huaweicloud.ServiceClient1, error) {
	pid, e := GetProjectId(client)
	if e != nil {
		return nil, e
	}

	c, e := initClientOpts(client, eo, clientType)
	if e != nil {
		return nil, e
	}

	sc := new(huaweicloud.ServiceClient1)
	sc.ServiceClient = c
	sc.ProjectID = pid
	return sc, nil
}

// NewOtcV1 creates a ServiceClient that may be used with the v1 network package.
func NewOtcV1(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts, otctype string) (*huaweicloud.ServiceClient1, error) {
	sc, err := initClientOpts1(client, eo, "compute")
	//fmt.Printf("client=%+v.\n", sc)
	sc.Endpoint = strings.Replace(strings.Replace(sc.Endpoint, "ecs", otctype, 1), "/v2/", "/v1.0/", 1)
	//fmt.Printf("url=%s.\n", sc.Endpoint)
	sc.ResourceBase = sc.Endpoint
	sc.Type = otctype
	return sc, err
}

func NewCESClient(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*huaweicloud.ServiceClient1, error) {
	sc, err := initClientOpts1(client, eo, "ces")
	if err != nil {
		return nil, err
	}
	sc.ResourceBase = sc.Endpoint
	return sc, err
}

// NewSmnServiceV2 creates a ServiceClient that may be used to access the v2 Simple Message Notification service.
func NewSmnServiceV2(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {

	sc, err := initClientOpts(client, eo, "compute")
	sc.Endpoint = strings.Replace(sc.Endpoint, "ecs", "smn", 1)
	sc.ResourceBase = sc.Endpoint + "notifications/"
	sc.Type = "smn"
	return sc, err
}
