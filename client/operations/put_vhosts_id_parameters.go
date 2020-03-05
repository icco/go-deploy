// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/aptible/go-deploy/models"
)

// NewPutVhostsIDParams creates a new PutVhostsIDParams object
// with the default values initialized.
func NewPutVhostsIDParams() *PutVhostsIDParams {
	var ()
	return &PutVhostsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutVhostsIDParamsWithTimeout creates a new PutVhostsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutVhostsIDParamsWithTimeout(timeout time.Duration) *PutVhostsIDParams {
	var ()
	return &PutVhostsIDParams{

		timeout: timeout,
	}
}

// NewPutVhostsIDParamsWithContext creates a new PutVhostsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutVhostsIDParamsWithContext(ctx context.Context) *PutVhostsIDParams {
	var ()
	return &PutVhostsIDParams{

		Context: ctx,
	}
}

// NewPutVhostsIDParamsWithHTTPClient creates a new PutVhostsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutVhostsIDParamsWithHTTPClient(client *http.Client) *PutVhostsIDParams {
	var ()
	return &PutVhostsIDParams{
		HTTPClient: client,
	}
}

/*PutVhostsIDParams contains all the parameters to send to the API endpoint
for the put vhosts ID operation typically these are written to a http.Request
*/
type PutVhostsIDParams struct {

	/*AppRequest*/
	AppRequest *models.AppRequest34
	/*ID
	  id

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put vhosts ID params
func (o *PutVhostsIDParams) WithTimeout(timeout time.Duration) *PutVhostsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put vhosts ID params
func (o *PutVhostsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put vhosts ID params
func (o *PutVhostsIDParams) WithContext(ctx context.Context) *PutVhostsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put vhosts ID params
func (o *PutVhostsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put vhosts ID params
func (o *PutVhostsIDParams) WithHTTPClient(client *http.Client) *PutVhostsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put vhosts ID params
func (o *PutVhostsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppRequest adds the appRequest to the put vhosts ID params
func (o *PutVhostsIDParams) WithAppRequest(appRequest *models.AppRequest34) *PutVhostsIDParams {
	o.SetAppRequest(appRequest)
	return o
}

// SetAppRequest adds the appRequest to the put vhosts ID params
func (o *PutVhostsIDParams) SetAppRequest(appRequest *models.AppRequest34) {
	o.AppRequest = appRequest
}

// WithID adds the id to the put vhosts ID params
func (o *PutVhostsIDParams) WithID(id int64) *PutVhostsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put vhosts ID params
func (o *PutVhostsIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutVhostsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AppRequest != nil {
		if err := r.SetBodyParam(o.AppRequest); err != nil {
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
