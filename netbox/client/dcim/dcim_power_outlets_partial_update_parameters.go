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

// NewDcimPowerOutletsPartialUpdateParams creates a new DcimPowerOutletsPartialUpdateParams object
// with the default values initialized.
func NewDcimPowerOutletsPartialUpdateParams() *DcimPowerOutletsPartialUpdateParams {
	var ()
	return &DcimPowerOutletsPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerOutletsPartialUpdateParamsWithTimeout creates a new DcimPowerOutletsPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimPowerOutletsPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimPowerOutletsPartialUpdateParams {
	var ()
	return &DcimPowerOutletsPartialUpdateParams{

		timeout: timeout,
	}
}

// NewDcimPowerOutletsPartialUpdateParamsWithContext creates a new DcimPowerOutletsPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimPowerOutletsPartialUpdateParamsWithContext(ctx context.Context) *DcimPowerOutletsPartialUpdateParams {
	var ()
	return &DcimPowerOutletsPartialUpdateParams{

		Context: ctx,
	}
}

// NewDcimPowerOutletsPartialUpdateParamsWithHTTPClient creates a new DcimPowerOutletsPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimPowerOutletsPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimPowerOutletsPartialUpdateParams {
	var ()
	return &DcimPowerOutletsPartialUpdateParams{
		HTTPClient: client,
	}
}

/*DcimPowerOutletsPartialUpdateParams contains all the parameters to send to the API endpoint
for the dcim power outlets partial update operation typically these are written to a http.Request
*/
type DcimPowerOutletsPartialUpdateParams struct {

	/*Data*/
	Data *models.PowerOutlet
	/*ID
	  A unique integer value identifying this power outlet.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim power outlets partial update params
func (o *DcimPowerOutletsPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimPowerOutletsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power outlets partial update params
func (o *DcimPowerOutletsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power outlets partial update params
func (o *DcimPowerOutletsPartialUpdateParams) WithContext(ctx context.Context) *DcimPowerOutletsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power outlets partial update params
func (o *DcimPowerOutletsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power outlets partial update params
func (o *DcimPowerOutletsPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimPowerOutletsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power outlets partial update params
func (o *DcimPowerOutletsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim power outlets partial update params
func (o *DcimPowerOutletsPartialUpdateParams) WithData(data *models.PowerOutlet) *DcimPowerOutletsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim power outlets partial update params
func (o *DcimPowerOutletsPartialUpdateParams) SetData(data *models.PowerOutlet) {
	o.Data = data
}

// WithID adds the id to the dcim power outlets partial update params
func (o *DcimPowerOutletsPartialUpdateParams) WithID(id int64) *DcimPowerOutletsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim power outlets partial update params
func (o *DcimPowerOutletsPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerOutletsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
