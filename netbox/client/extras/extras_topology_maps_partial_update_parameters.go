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

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tpretz/go-netbox/netbox/models"
)

// NewExtrasTopologyMapsPartialUpdateParams creates a new ExtrasTopologyMapsPartialUpdateParams object
// with the default values initialized.
func NewExtrasTopologyMapsPartialUpdateParams() *ExtrasTopologyMapsPartialUpdateParams {
	var ()
	return &ExtrasTopologyMapsPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasTopologyMapsPartialUpdateParamsWithTimeout creates a new ExtrasTopologyMapsPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExtrasTopologyMapsPartialUpdateParamsWithTimeout(timeout time.Duration) *ExtrasTopologyMapsPartialUpdateParams {
	var ()
	return &ExtrasTopologyMapsPartialUpdateParams{

		timeout: timeout,
	}
}

// NewExtrasTopologyMapsPartialUpdateParamsWithContext creates a new ExtrasTopologyMapsPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewExtrasTopologyMapsPartialUpdateParamsWithContext(ctx context.Context) *ExtrasTopologyMapsPartialUpdateParams {
	var ()
	return &ExtrasTopologyMapsPartialUpdateParams{

		Context: ctx,
	}
}

// NewExtrasTopologyMapsPartialUpdateParamsWithHTTPClient creates a new ExtrasTopologyMapsPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExtrasTopologyMapsPartialUpdateParamsWithHTTPClient(client *http.Client) *ExtrasTopologyMapsPartialUpdateParams {
	var ()
	return &ExtrasTopologyMapsPartialUpdateParams{
		HTTPClient: client,
	}
}

/*ExtrasTopologyMapsPartialUpdateParams contains all the parameters to send to the API endpoint
for the extras topology maps partial update operation typically these are written to a http.Request
*/
type ExtrasTopologyMapsPartialUpdateParams struct {

	/*Data*/
	Data *models.TopologyMap
	/*ID
	  A unique integer value identifying this topology map.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the extras topology maps partial update params
func (o *ExtrasTopologyMapsPartialUpdateParams) WithTimeout(timeout time.Duration) *ExtrasTopologyMapsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras topology maps partial update params
func (o *ExtrasTopologyMapsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras topology maps partial update params
func (o *ExtrasTopologyMapsPartialUpdateParams) WithContext(ctx context.Context) *ExtrasTopologyMapsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras topology maps partial update params
func (o *ExtrasTopologyMapsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras topology maps partial update params
func (o *ExtrasTopologyMapsPartialUpdateParams) WithHTTPClient(client *http.Client) *ExtrasTopologyMapsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras topology maps partial update params
func (o *ExtrasTopologyMapsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras topology maps partial update params
func (o *ExtrasTopologyMapsPartialUpdateParams) WithData(data *models.TopologyMap) *ExtrasTopologyMapsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras topology maps partial update params
func (o *ExtrasTopologyMapsPartialUpdateParams) SetData(data *models.TopologyMap) {
	o.Data = data
}

// WithID adds the id to the extras topology maps partial update params
func (o *ExtrasTopologyMapsPartialUpdateParams) WithID(id int64) *ExtrasTopologyMapsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras topology maps partial update params
func (o *ExtrasTopologyMapsPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasTopologyMapsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
