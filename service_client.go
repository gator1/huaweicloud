package huaweicloud

import (
        "io"
        "net/http"
	"github.com/gophercloud/gophercloud"
)
 

type ServiceClient1 struct {
	// ServiceClient is a reference to the ServiceClient.
	*gophercloud.ServiceClient

	// ProjectID is the ID of project to which User is authorized.
	ProjectID string
}

func (client *ServiceClient1) initReqOpts(url string, JSONBody interface{}, JSONResponse interface{}, opts *gophercloud.RequestOpts) {
        if v, ok := (JSONBody).(io.Reader); ok {
                opts.RawBody = v
        } else if JSONBody != nil {
                opts.JSONBody = JSONBody
        }

        if JSONResponse != nil {
                opts.JSONResponse = JSONResponse
        }

        if opts.MoreHeaders == nil {
                opts.MoreHeaders = make(map[string]string)
        }

        if client.Microversion != "" {
                client.setMicroversionHeader(opts)
        }
}


// Delete calls `Request` with the "DELETE" HTTP verb.
func (client *ServiceClient1) Delete0(url string, opts *gophercloud.RequestOpts) (*http.Response, error) {
	if opts == nil {
		opts = new(gophercloud.RequestOpts)
	}
	JSONResponse := new(interface{})
	client.initReqOpts(url, nil, JSONResponse, opts)
	return client.Request("DELETE", url, opts)
}

// Delete calls `Request` with the "DELETE" HTTP verb.
func (client *ServiceClient1) Delete2(url string, JSONResponse interface{}, opts *gophercloud.RequestOpts) (*http.Response, error) {
	if opts == nil {
		opts = new(gophercloud.RequestOpts)
	}
	client.initReqOpts(url, nil, JSONResponse, opts)
	return client.Request("DELETE", url, opts)
}

func (client *ServiceClient1) setMicroversionHeader(opts *gophercloud.RequestOpts) {
	switch client.Type {
	case "compute":
		opts.MoreHeaders["X-OpenStack-Nova-API-Version"] = client.Microversion
	case "sharev2":
		opts.MoreHeaders["X-OpenStack-Manila-API-Version"] = client.Microversion
	}

	if client.Type != "" {
		opts.MoreHeaders["OpenStack-API-Version"] = client.Type + " " + client.Microversion
	}
}
