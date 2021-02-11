// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package virtualization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tpretz/go-netbox/netbox/models"
)

// NewVirtualizationClusterGroupsCreateParams creates a new VirtualizationClusterGroupsCreateParams object
// with the default values initialized.
func NewVirtualizationClusterGroupsCreateParams() *VirtualizationClusterGroupsCreateParams {
	var ()
	return &VirtualizationClusterGroupsCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationClusterGroupsCreateParamsWithTimeout creates a new VirtualizationClusterGroupsCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVirtualizationClusterGroupsCreateParamsWithTimeout(timeout time.Duration) *VirtualizationClusterGroupsCreateParams {
	var ()
	return &VirtualizationClusterGroupsCreateParams{

		timeout: timeout,
	}
}

// NewVirtualizationClusterGroupsCreateParamsWithContext creates a new VirtualizationClusterGroupsCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewVirtualizationClusterGroupsCreateParamsWithContext(ctx context.Context) *VirtualizationClusterGroupsCreateParams {
	var ()
	return &VirtualizationClusterGroupsCreateParams{

		Context: ctx,
	}
}

// NewVirtualizationClusterGroupsCreateParamsWithHTTPClient creates a new VirtualizationClusterGroupsCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVirtualizationClusterGroupsCreateParamsWithHTTPClient(client *http.Client) *VirtualizationClusterGroupsCreateParams {
	var ()
	return &VirtualizationClusterGroupsCreateParams{
		HTTPClient: client,
	}
}

/*VirtualizationClusterGroupsCreateParams contains all the parameters to send to the API endpoint
for the virtualization cluster groups create operation typically these are written to a http.Request
*/
type VirtualizationClusterGroupsCreateParams struct {

	/*Data*/
	Data *models.ClusterGroup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the virtualization cluster groups create params
func (o *VirtualizationClusterGroupsCreateParams) WithTimeout(timeout time.Duration) *VirtualizationClusterGroupsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization cluster groups create params
func (o *VirtualizationClusterGroupsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization cluster groups create params
func (o *VirtualizationClusterGroupsCreateParams) WithContext(ctx context.Context) *VirtualizationClusterGroupsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization cluster groups create params
func (o *VirtualizationClusterGroupsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization cluster groups create params
func (o *VirtualizationClusterGroupsCreateParams) WithHTTPClient(client *http.Client) *VirtualizationClusterGroupsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization cluster groups create params
func (o *VirtualizationClusterGroupsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the virtualization cluster groups create params
func (o *VirtualizationClusterGroupsCreateParams) WithData(data *models.ClusterGroup) *VirtualizationClusterGroupsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the virtualization cluster groups create params
func (o *VirtualizationClusterGroupsCreateParams) SetData(data *models.ClusterGroup) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationClusterGroupsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
