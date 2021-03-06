package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/httpkit/middleware"
	"github.com/go-swagger/go-swagger/httpkit/security"
	"github.com/go-swagger/go-swagger/spec"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/go-swagger/go-swagger/examples/todo-list/restapi/operations/todos"
)

// NewSimpleToDoListAPIAPI creates a new SimpleToDoListAPI instance
func NewSimpleToDoListAPIAPI(spec *spec.Document) *SimpleToDoListAPIAPI {
	o := &SimpleToDoListAPIAPI{
		spec:            spec,
		handlers:        make(map[string]http.Handler),
		formats:         strfmt.Default,
		defaultConsumes: "application/io.swagger.examples.todo-list.v1+json",
		defaultProduces: "application/io.swagger.examples.todo-list.v1+json",
	}

	return o
}

/*SimpleToDoListAPIAPI the simple to do list API API */
type SimpleToDoListAPIAPI struct {
	spec            *spec.Document
	context         *middleware.Context
	handlers        map[string]http.Handler
	formats         strfmt.Registry
	defaultConsumes string
	defaultProduces string
	// JSONConsumer registers a consumer for a "application/io.swagger.examples.todo-list.v1+json" mime type
	JSONConsumer httpkit.Consumer

	// JSONProducer registers a producer for a "application/io.swagger.examples.todo-list.v1+json" mime type
	JSONProducer httpkit.Producer

	// KeyAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key x-petstore-token provided in the header
	KeyAuth func(string) (interface{}, error)

	// AddOneHandler sets the operation handler for the add one operation
	AddOneHandler todos.AddOneHandler
	// DestroyOneHandler sets the operation handler for the destroy one operation
	DestroyOneHandler todos.DestroyOneHandler
	// FindHandler sets the operation handler for the find operation
	FindHandler todos.FindHandler
	// UpdateOneHandler sets the operation handler for the update one operation
	UpdateOneHandler todos.UpdateOneHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)
}

// SetDefaultProduces sets the default produces media type
func (o *SimpleToDoListAPIAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *SimpleToDoListAPIAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// DefaultProduces returns the default produces media type
func (o *SimpleToDoListAPIAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *SimpleToDoListAPIAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *SimpleToDoListAPIAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *SimpleToDoListAPIAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the SimpleToDoListAPIAPI
func (o *SimpleToDoListAPIAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.KeyAuth == nil {
		unregistered = append(unregistered, "XPetstoreTokenAuth")
	}

	if o.AddOneHandler == nil {
		unregistered = append(unregistered, "AddOneHandler")
	}

	if o.DestroyOneHandler == nil {
		unregistered = append(unregistered, "DestroyOneHandler")
	}

	if o.FindHandler == nil {
		unregistered = append(unregistered, "FindHandler")
	}

	if o.UpdateOneHandler == nil {
		unregistered = append(unregistered, "UpdateOneHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *SimpleToDoListAPIAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *SimpleToDoListAPIAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]httpkit.Authenticator {

	result := make(map[string]httpkit.Authenticator)
	for name, scheme := range schemes {
		switch name {

		case "key":

			result[name] = security.APIKeyAuth(scheme.Name, scheme.In, func(tok string) (interface{}, error) { return o.KeyAuth(tok) })

		}
	}
	return result

}

// ConsumersFor gets the consumers for the specified media types
func (o *SimpleToDoListAPIAPI) ConsumersFor(mediaTypes []string) map[string]httpkit.Consumer {

	result := make(map[string]httpkit.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/io.swagger.examples.todo-list.v1+json":
			result["application/io.swagger.examples.todo-list.v1+json"] = o.JSONConsumer

		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *SimpleToDoListAPIAPI) ProducersFor(mediaTypes []string) map[string]httpkit.Producer {

	result := make(map[string]httpkit.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/io.swagger.examples.todo-list.v1+json":
			result["application/io.swagger.examples.todo-list.v1+json"] = o.JSONProducer

		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation id
func (o *SimpleToDoListAPIAPI) HandlerFor(operationID string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	h, ok := o.handlers[operationID]
	return h, ok
}

func (o *SimpleToDoListAPIAPI) initHandlerCache() {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	o.handlers = make(map[string]http.Handler)

	o.handlers["addOne"] = todos.NewAddOne(o.context, o.AddOneHandler)

	o.handlers["destroyOne"] = todos.NewDestroyOne(o.context, o.DestroyOneHandler)

	o.handlers["find"] = todos.NewFind(o.context, o.FindHandler)

	o.handlers["updateOne"] = todos.NewUpdateOne(o.context, o.UpdateOneHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve())
func (o *SimpleToDoListAPIAPI) Serve() http.Handler {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}

	return o.context.APIHandler()
}
