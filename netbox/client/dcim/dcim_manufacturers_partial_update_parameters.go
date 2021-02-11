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

package dcim

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

// NewDcimManufacturersPartialUpdateParams creates a new DcimManufacturersPartialUpdateParams object
// with the default values initialized.
func NewDcimManufacturersPartialUpdateParams() *DcimManufacturersPartialUpdateParams {
	var ()
	return &DcimManufacturersPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimManufacturersPartialUpdateParamsWithTimeout creates a new DcimManufacturersPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimManufacturersPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimManufacturersPartialUpdateParams {
	var ()
	return &DcimManufacturersPartialUpdateParams{

		timeout: timeout,
	}
}

// NewDcimManufacturersPartialUpdateParamsWithContext creates a new DcimManufacturersPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimManufacturersPartialUpdateParamsWithContext(ctx context.Context) *DcimManufacturersPartialUpdateParams {
	var ()
	return &DcimManufacturersPartialUpdateParams{

		Context: ctx,
	}
}

// NewDcimManufacturersPartialUpdateParamsWithHTTPClient creates a new DcimManufacturersPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimManufacturersPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimManufacturersPartialUpdateParams {
	var ()
	return &DcimManufacturersPartialUpdateParams{
		HTTPClient: client,
	}
}

/*DcimManufacturersPartialUpdateParams contains all the parameters to send to the API endpoint
for the dcim manufacturers partial update operation typically these are written to a http.Request
*/
type DcimManufacturersPartialUpdateParams struct {

	/*Data*/
	Data *models.Manufacturer
	/*ID
	  A unique integer value identifying this manufacturer.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim manufacturers partial update params
func (o *DcimManufacturersPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimManufacturersPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim manufacturers partial update params
func (o *DcimManufacturersPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim manufacturers partial update params
func (o *DcimManufacturersPartialUpdateParams) WithContext(ctx context.Context) *DcimManufacturersPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim manufacturers partial update params
func (o *DcimManufacturersPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim manufacturers partial update params
func (o *DcimManufacturersPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimManufacturersPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim manufacturers partial update params
func (o *DcimManufacturersPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim manufacturers partial update params
func (o *DcimManufacturersPartialUpdateParams) WithData(data *models.Manufacturer) *DcimManufacturersPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim manufacturers partial update params
func (o *DcimManufacturersPartialUpdateParams) SetData(data *models.Manufacturer) {
	o.Data = data
}

// WithID adds the id to the dcim manufacturers partial update params
func (o *DcimManufacturersPartialUpdateParams) WithID(id int64) *DcimManufacturersPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim manufacturers partial update params
func (o *DcimManufacturersPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimManufacturersPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
