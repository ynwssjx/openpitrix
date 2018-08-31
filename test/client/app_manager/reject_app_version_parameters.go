// Code generated by go-swagger; DO NOT EDIT.

package app_manager

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

	"openpitrix.io/openpitrix/test/models"
)

// NewRejectAppVersionParams creates a new RejectAppVersionParams object
// with the default values initialized.
func NewRejectAppVersionParams() *RejectAppVersionParams {
	var ()
	return &RejectAppVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRejectAppVersionParamsWithTimeout creates a new RejectAppVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRejectAppVersionParamsWithTimeout(timeout time.Duration) *RejectAppVersionParams {
	var ()
	return &RejectAppVersionParams{

		timeout: timeout,
	}
}

// NewRejectAppVersionParamsWithContext creates a new RejectAppVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewRejectAppVersionParamsWithContext(ctx context.Context) *RejectAppVersionParams {
	var ()
	return &RejectAppVersionParams{

		Context: ctx,
	}
}

// NewRejectAppVersionParamsWithHTTPClient creates a new RejectAppVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRejectAppVersionParamsWithHTTPClient(client *http.Client) *RejectAppVersionParams {
	var ()
	return &RejectAppVersionParams{
		HTTPClient: client,
	}
}

/*RejectAppVersionParams contains all the parameters to send to the API endpoint
for the reject app version operation typically these are written to a http.Request
*/
type RejectAppVersionParams struct {

	/*Body*/
	Body *models.OpenpitrixRejectAppVersionRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the reject app version params
func (o *RejectAppVersionParams) WithTimeout(timeout time.Duration) *RejectAppVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the reject app version params
func (o *RejectAppVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the reject app version params
func (o *RejectAppVersionParams) WithContext(ctx context.Context) *RejectAppVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the reject app version params
func (o *RejectAppVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the reject app version params
func (o *RejectAppVersionParams) WithHTTPClient(client *http.Client) *RejectAppVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the reject app version params
func (o *RejectAppVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the reject app version params
func (o *RejectAppVersionParams) WithBody(body *models.OpenpitrixRejectAppVersionRequest) *RejectAppVersionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the reject app version params
func (o *RejectAppVersionParams) SetBody(body *models.OpenpitrixRejectAppVersionRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *RejectAppVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}