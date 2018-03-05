// Code generated by go-swagger; DO NOT EDIT.

package task_manager

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
)

// NewDescribeTasksParams creates a new DescribeTasksParams object
// with the default values initialized.
func NewDescribeTasksParams() *DescribeTasksParams {
	var ()
	return &DescribeTasksParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDescribeTasksParamsWithTimeout creates a new DescribeTasksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDescribeTasksParamsWithTimeout(timeout time.Duration) *DescribeTasksParams {
	var ()
	return &DescribeTasksParams{

		timeout: timeout,
	}
}

// NewDescribeTasksParamsWithContext creates a new DescribeTasksParams object
// with the default values initialized, and the ability to set a context for a request
func NewDescribeTasksParamsWithContext(ctx context.Context) *DescribeTasksParams {
	var ()
	return &DescribeTasksParams{

		Context: ctx,
	}
}

// NewDescribeTasksParamsWithHTTPClient creates a new DescribeTasksParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDescribeTasksParamsWithHTTPClient(client *http.Client) *DescribeTasksParams {
	var ()
	return &DescribeTasksParams{
		HTTPClient: client,
	}
}

/*DescribeTasksParams contains all the parameters to send to the API endpoint
for the describe tasks operation typically these are written to a http.Request
*/
type DescribeTasksParams struct {

	/*JobID*/
	JobID []string
	/*Limit
	  default is 20, max value is 200.

	*/
	Limit *int64
	/*Offset
	  default is 0.

	*/
	Offset *int64
	/*Status*/
	Status []string
	/*TaskID*/
	TaskID []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the describe tasks params
func (o *DescribeTasksParams) WithTimeout(timeout time.Duration) *DescribeTasksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe tasks params
func (o *DescribeTasksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe tasks params
func (o *DescribeTasksParams) WithContext(ctx context.Context) *DescribeTasksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe tasks params
func (o *DescribeTasksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe tasks params
func (o *DescribeTasksParams) WithHTTPClient(client *http.Client) *DescribeTasksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe tasks params
func (o *DescribeTasksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJobID adds the jobID to the describe tasks params
func (o *DescribeTasksParams) WithJobID(jobID []string) *DescribeTasksParams {
	o.SetJobID(jobID)
	return o
}

// SetJobID adds the jobId to the describe tasks params
func (o *DescribeTasksParams) SetJobID(jobID []string) {
	o.JobID = jobID
}

// WithLimit adds the limit to the describe tasks params
func (o *DescribeTasksParams) WithLimit(limit *int64) *DescribeTasksParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the describe tasks params
func (o *DescribeTasksParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the describe tasks params
func (o *DescribeTasksParams) WithOffset(offset *int64) *DescribeTasksParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the describe tasks params
func (o *DescribeTasksParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithStatus adds the status to the describe tasks params
func (o *DescribeTasksParams) WithStatus(status []string) *DescribeTasksParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the describe tasks params
func (o *DescribeTasksParams) SetStatus(status []string) {
	o.Status = status
}

// WithTaskID adds the taskID to the describe tasks params
func (o *DescribeTasksParams) WithTaskID(taskID []string) *DescribeTasksParams {
	o.SetTaskID(taskID)
	return o
}

// SetTaskID adds the taskId to the describe tasks params
func (o *DescribeTasksParams) SetTaskID(taskID []string) {
	o.TaskID = taskID
}

// WriteToRequest writes these params to a swagger request
func (o *DescribeTasksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesJobID := o.JobID

	joinedJobID := swag.JoinByFormat(valuesJobID, "")
	// query array param job_id
	if err := r.SetQueryParam("job_id", joinedJobID...); err != nil {
		return err
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	valuesStatus := o.Status

	joinedStatus := swag.JoinByFormat(valuesStatus, "")
	// query array param status
	if err := r.SetQueryParam("status", joinedStatus...); err != nil {
		return err
	}

	valuesTaskID := o.TaskID

	joinedTaskID := swag.JoinByFormat(valuesTaskID, "")
	// query array param task_id
	if err := r.SetQueryParam("task_id", joinedTaskID...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
