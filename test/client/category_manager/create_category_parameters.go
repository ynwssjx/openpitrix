// Code generated by go-swagger; DO NOT EDIT.

package category_manager

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

// NewCreateCategoryParams creates a new CreateCategoryParams object
// with the default values initialized.
func NewCreateCategoryParams() *CreateCategoryParams {
	var ()
	return &CreateCategoryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateCategoryParamsWithTimeout creates a new CreateCategoryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateCategoryParamsWithTimeout(timeout time.Duration) *CreateCategoryParams {
	var ()
	return &CreateCategoryParams{

		timeout: timeout,
	}
}

// NewCreateCategoryParamsWithContext creates a new CreateCategoryParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateCategoryParamsWithContext(ctx context.Context) *CreateCategoryParams {
	var ()
	return &CreateCategoryParams{

		Context: ctx,
	}
}

// NewCreateCategoryParamsWithHTTPClient creates a new CreateCategoryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateCategoryParamsWithHTTPClient(client *http.Client) *CreateCategoryParams {
	var ()
	return &CreateCategoryParams{
		HTTPClient: client,
	}
}

/*CreateCategoryParams contains all the parameters to send to the API endpoint
for the create category operation typically these are written to a http.Request
*/
type CreateCategoryParams struct {

	/*Body*/
	Body *models.OpenpitrixCreateCategoryRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create category params
func (o *CreateCategoryParams) WithTimeout(timeout time.Duration) *CreateCategoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create category params
func (o *CreateCategoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create category params
func (o *CreateCategoryParams) WithContext(ctx context.Context) *CreateCategoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create category params
func (o *CreateCategoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create category params
func (o *CreateCategoryParams) WithHTTPClient(client *http.Client) *CreateCategoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create category params
func (o *CreateCategoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create category params
func (o *CreateCategoryParams) WithBody(body *models.OpenpitrixCreateCategoryRequest) *CreateCategoryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create category params
func (o *CreateCategoryParams) SetBody(body *models.OpenpitrixCreateCategoryRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateCategoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
