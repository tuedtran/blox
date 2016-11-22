// Copyright 2016 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateDeployment Create a deployment under provided environment
*/
func (a *Client) CreateDeployment(params *CreateDeploymentParams) (*CreateDeploymentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDeploymentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createDeployment",
		Method:             "POST",
		PathPattern:        "/environments/{name}/deployments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateDeploymentReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateDeploymentOK), nil

}

/*
CreateEnvironment Create a new environment
*/
func (a *Client) CreateEnvironment(params *CreateEnvironmentParams) (*CreateEnvironmentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateEnvironmentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createEnvironment",
		Method:             "POST",
		PathPattern:        "/environments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateEnvironmentReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateEnvironmentOK), nil

}

/*
DeleteEnvironment Delete an environment by name
*/
func (a *Client) DeleteEnvironment(params *DeleteEnvironmentParams) (*DeleteEnvironmentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteEnvironmentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteEnvironment",
		Method:             "DELETE",
		PathPattern:        "/environments/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteEnvironmentReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteEnvironmentOK), nil

}

/*
GetDeployment Get deployment
*/
func (a *Client) GetDeployment(params *GetDeploymentParams) (*GetDeploymentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeploymentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDeployment",
		Method:             "GET",
		PathPattern:        "/environments/{name}/deployments/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDeploymentReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDeploymentOK), nil

}

/*
GetEnvironment Get an environment by name
*/
func (a *Client) GetEnvironment(params *GetEnvironmentParams) (*GetEnvironmentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEnvironmentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getEnvironment",
		Method:             "GET",
		PathPattern:        "/environments/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetEnvironmentReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetEnvironmentOK), nil

}

/*
ListDeployments Get deployments under provided environment
*/
func (a *Client) ListDeployments(params *ListDeploymentsParams) (*ListDeploymentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListDeploymentsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listDeployments",
		Method:             "GET",
		PathPattern:        "/environments/{name}/deployments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListDeploymentsReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListDeploymentsOK), nil

}

/*
ListEnvironments Gets all the environments
*/
func (a *Client) ListEnvironments(params *ListEnvironmentsParams) (*ListEnvironmentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListEnvironmentsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listEnvironments",
		Method:             "GET",
		PathPattern:        "/environments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListEnvironmentsReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListEnvironmentsOK), nil

}

/*
Ping Ping server for health status
*/
func (a *Client) Ping(params *PingParams) (*PingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPingParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ping",
		Method:             "GET",
		PathPattern:        "/ping",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PingReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PingOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
