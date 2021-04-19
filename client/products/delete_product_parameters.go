// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewDeleteProductParams creates a new DeleteProductParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteProductParams() *DeleteProductParams {
	return &DeleteProductParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteProductParamsWithTimeout creates a new DeleteProductParams object
// with the ability to set a timeout on a request.
func NewDeleteProductParamsWithTimeout(timeout time.Duration) *DeleteProductParams {
	return &DeleteProductParams{
		timeout: timeout,
	}
}

// NewDeleteProductParamsWithContext creates a new DeleteProductParams object
// with the ability to set a context for a request.
func NewDeleteProductParamsWithContext(ctx context.Context) *DeleteProductParams {
	return &DeleteProductParams{
		Context: ctx,
	}
}

// NewDeleteProductParamsWithHTTPClient creates a new DeleteProductParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteProductParamsWithHTTPClient(client *http.Client) *DeleteProductParams {
	return &DeleteProductParams{
		HTTPClient: client,
	}
}

/* DeleteProductParams contains all the parameters to send to the API endpoint
   for the delete product operation.

   Typically these are written to a http.Request.
*/
type DeleteProductParams struct {

	/* ID.

	   ID of the product to be deleted.

	   Format: int64
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete product params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteProductParams) WithDefaults() *DeleteProductParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete product params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteProductParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete product params
func (o *DeleteProductParams) WithTimeout(timeout time.Duration) *DeleteProductParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete product params
func (o *DeleteProductParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete product params
func (o *DeleteProductParams) WithContext(ctx context.Context) *DeleteProductParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete product params
func (o *DeleteProductParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete product params
func (o *DeleteProductParams) WithHTTPClient(client *http.Client) *DeleteProductParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete product params
func (o *DeleteProductParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete product params
func (o *DeleteProductParams) WithID(id int64) *DeleteProductParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete product params
func (o *DeleteProductParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteProductParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
