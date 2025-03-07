// Code generated by go-swagger; DO NOT EDIT.

package read

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

// NewGetRelationTuplesParams creates a new GetRelationTuplesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRelationTuplesParams() *GetRelationTuplesParams {
	return &GetRelationTuplesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRelationTuplesParamsWithTimeout creates a new GetRelationTuplesParams object
// with the ability to set a timeout on a request.
func NewGetRelationTuplesParamsWithTimeout(timeout time.Duration) *GetRelationTuplesParams {
	return &GetRelationTuplesParams{
		timeout: timeout,
	}
}

// NewGetRelationTuplesParamsWithContext creates a new GetRelationTuplesParams object
// with the ability to set a context for a request.
func NewGetRelationTuplesParamsWithContext(ctx context.Context) *GetRelationTuplesParams {
	return &GetRelationTuplesParams{
		Context: ctx,
	}
}

// NewGetRelationTuplesParamsWithHTTPClient creates a new GetRelationTuplesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRelationTuplesParamsWithHTTPClient(client *http.Client) *GetRelationTuplesParams {
	return &GetRelationTuplesParams{
		HTTPClient: client,
	}
}

/* GetRelationTuplesParams contains all the parameters to send to the API endpoint
   for the get relation tuples operation.

   Typically these are written to a http.Request.
*/
type GetRelationTuplesParams struct {

	/* Namespace.

	   Namespace of the Relation Tuple
	*/
	Namespace string

	/* Object.

	   Object of the Relation Tuple
	*/
	Object *string

	// PageSize.
	//
	// Format: int64
	PageSize *int64

	// PageToken.
	PageToken *string

	/* Relation.

	   Relation of the Relation Tuple
	*/
	Relation *string

	/* SubjectID.

	   SubjectID of the Relation Tuple
	*/
	SubjectID *string

	/* SubjectSetNamespace.

	   Namespace of the Subject Set
	*/
	SubjectSetNamespace *string

	/* SubjectSetObject.

	   Object of the Subject Set
	*/
	SubjectSetObject *string

	/* SubjectSetRelation.

	   Relation of the Subject Set
	*/
	SubjectSetRelation *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get relation tuples params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRelationTuplesParams) WithDefaults() *GetRelationTuplesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get relation tuples params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRelationTuplesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get relation tuples params
func (o *GetRelationTuplesParams) WithTimeout(timeout time.Duration) *GetRelationTuplesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get relation tuples params
func (o *GetRelationTuplesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get relation tuples params
func (o *GetRelationTuplesParams) WithContext(ctx context.Context) *GetRelationTuplesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get relation tuples params
func (o *GetRelationTuplesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get relation tuples params
func (o *GetRelationTuplesParams) WithHTTPClient(client *http.Client) *GetRelationTuplesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get relation tuples params
func (o *GetRelationTuplesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the get relation tuples params
func (o *GetRelationTuplesParams) WithNamespace(namespace string) *GetRelationTuplesParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get relation tuples params
func (o *GetRelationTuplesParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithObject adds the object to the get relation tuples params
func (o *GetRelationTuplesParams) WithObject(object *string) *GetRelationTuplesParams {
	o.SetObject(object)
	return o
}

// SetObject adds the object to the get relation tuples params
func (o *GetRelationTuplesParams) SetObject(object *string) {
	o.Object = object
}

// WithPageSize adds the pageSize to the get relation tuples params
func (o *GetRelationTuplesParams) WithPageSize(pageSize *int64) *GetRelationTuplesParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get relation tuples params
func (o *GetRelationTuplesParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithPageToken adds the pageToken to the get relation tuples params
func (o *GetRelationTuplesParams) WithPageToken(pageToken *string) *GetRelationTuplesParams {
	o.SetPageToken(pageToken)
	return o
}

// SetPageToken adds the pageToken to the get relation tuples params
func (o *GetRelationTuplesParams) SetPageToken(pageToken *string) {
	o.PageToken = pageToken
}

// WithRelation adds the relation to the get relation tuples params
func (o *GetRelationTuplesParams) WithRelation(relation *string) *GetRelationTuplesParams {
	o.SetRelation(relation)
	return o
}

// SetRelation adds the relation to the get relation tuples params
func (o *GetRelationTuplesParams) SetRelation(relation *string) {
	o.Relation = relation
}

// WithSubjectID adds the subjectID to the get relation tuples params
func (o *GetRelationTuplesParams) WithSubjectID(subjectID *string) *GetRelationTuplesParams {
	o.SetSubjectID(subjectID)
	return o
}

// SetSubjectID adds the subjectId to the get relation tuples params
func (o *GetRelationTuplesParams) SetSubjectID(subjectID *string) {
	o.SubjectID = subjectID
}

// WithSubjectSetNamespace adds the subjectSetNamespace to the get relation tuples params
func (o *GetRelationTuplesParams) WithSubjectSetNamespace(subjectSetNamespace *string) *GetRelationTuplesParams {
	o.SetSubjectSetNamespace(subjectSetNamespace)
	return o
}

// SetSubjectSetNamespace adds the subjectSetNamespace to the get relation tuples params
func (o *GetRelationTuplesParams) SetSubjectSetNamespace(subjectSetNamespace *string) {
	o.SubjectSetNamespace = subjectSetNamespace
}

// WithSubjectSetObject adds the subjectSetObject to the get relation tuples params
func (o *GetRelationTuplesParams) WithSubjectSetObject(subjectSetObject *string) *GetRelationTuplesParams {
	o.SetSubjectSetObject(subjectSetObject)
	return o
}

// SetSubjectSetObject adds the subjectSetObject to the get relation tuples params
func (o *GetRelationTuplesParams) SetSubjectSetObject(subjectSetObject *string) {
	o.SubjectSetObject = subjectSetObject
}

// WithSubjectSetRelation adds the subjectSetRelation to the get relation tuples params
func (o *GetRelationTuplesParams) WithSubjectSetRelation(subjectSetRelation *string) *GetRelationTuplesParams {
	o.SetSubjectSetRelation(subjectSetRelation)
	return o
}

// SetSubjectSetRelation adds the subjectSetRelation to the get relation tuples params
func (o *GetRelationTuplesParams) SetSubjectSetRelation(subjectSetRelation *string) {
	o.SubjectSetRelation = subjectSetRelation
}

// WriteToRequest writes these params to a swagger request
func (o *GetRelationTuplesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param namespace
	qrNamespace := o.Namespace
	qNamespace := qrNamespace
	if qNamespace != "" {

		if err := r.SetQueryParam("namespace", qNamespace); err != nil {
			return err
		}
	}

	if o.Object != nil {

		// query param object
		var qrObject string

		if o.Object != nil {
			qrObject = *o.Object
		}
		qObject := qrObject
		if qObject != "" {

			if err := r.SetQueryParam("object", qObject); err != nil {
				return err
			}
		}
	}

	if o.PageSize != nil {

		// query param page_size
		var qrPageSize int64

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt64(qrPageSize)
		if qPageSize != "" {

			if err := r.SetQueryParam("page_size", qPageSize); err != nil {
				return err
			}
		}
	}

	if o.PageToken != nil {

		// query param page_token
		var qrPageToken string

		if o.PageToken != nil {
			qrPageToken = *o.PageToken
		}
		qPageToken := qrPageToken
		if qPageToken != "" {

			if err := r.SetQueryParam("page_token", qPageToken); err != nil {
				return err
			}
		}
	}

	if o.Relation != nil {

		// query param relation
		var qrRelation string

		if o.Relation != nil {
			qrRelation = *o.Relation
		}
		qRelation := qrRelation
		if qRelation != "" {

			if err := r.SetQueryParam("relation", qRelation); err != nil {
				return err
			}
		}
	}

	if o.SubjectID != nil {

		// query param subject_id
		var qrSubjectID string

		if o.SubjectID != nil {
			qrSubjectID = *o.SubjectID
		}
		qSubjectID := qrSubjectID
		if qSubjectID != "" {

			if err := r.SetQueryParam("subject_id", qSubjectID); err != nil {
				return err
			}
		}
	}

	if o.SubjectSetNamespace != nil {

		// query param subject_set.namespace
		var qrSubjectSetNamespace string

		if o.SubjectSetNamespace != nil {
			qrSubjectSetNamespace = *o.SubjectSetNamespace
		}
		qSubjectSetNamespace := qrSubjectSetNamespace
		if qSubjectSetNamespace != "" {

			if err := r.SetQueryParam("subject_set.namespace", qSubjectSetNamespace); err != nil {
				return err
			}
		}
	}

	if o.SubjectSetObject != nil {

		// query param subject_set.object
		var qrSubjectSetObject string

		if o.SubjectSetObject != nil {
			qrSubjectSetObject = *o.SubjectSetObject
		}
		qSubjectSetObject := qrSubjectSetObject
		if qSubjectSetObject != "" {

			if err := r.SetQueryParam("subject_set.object", qSubjectSetObject); err != nil {
				return err
			}
		}
	}

	if o.SubjectSetRelation != nil {

		// query param subject_set.relation
		var qrSubjectSetRelation string

		if o.SubjectSetRelation != nil {
			qrSubjectSetRelation = *o.SubjectSetRelation
		}
		qSubjectSetRelation := qrSubjectSetRelation
		if qSubjectSetRelation != "" {

			if err := r.SetQueryParam("subject_set.relation", qSubjectSetRelation); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
