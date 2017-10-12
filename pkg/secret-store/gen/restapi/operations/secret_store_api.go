///////////////////////////////////////////////////////////////////////
// Copyright (C) 2016 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////
// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"gitlab.eng.vmware.com/serverless/serverless/pkg/secret-store/gen/restapi/operations/secret"
)

// NewSecretStoreAPI creates a new SecretStore instance
func NewSecretStoreAPI(spec *loads.Document) *SecretStoreAPI {
	return &SecretStoreAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		SecretDeleteSecretNameHandler: secret.DeleteSecretNameHandlerFunc(func(params secret.DeleteSecretNameParams) middleware.Responder {
			return middleware.NotImplemented("operation SecretDeleteSecretName has not yet been implemented")
		}),
		SecretGetHandler: secret.GetHandlerFunc(func(params secret.GetParams) middleware.Responder {
			return middleware.NotImplemented("operation SecretGet has not yet been implemented")
		}),
		SecretGetSecretNameHandler: secret.GetSecretNameHandlerFunc(func(params secret.GetSecretNameParams) middleware.Responder {
			return middleware.NotImplemented("operation SecretGetSecretName has not yet been implemented")
		}),
		SecretPostHandler: secret.PostHandlerFunc(func(params secret.PostParams) middleware.Responder {
			return middleware.NotImplemented("operation SecretPost has not yet been implemented")
		}),
		SecretPutSecretNameHandler: secret.PutSecretNameHandlerFunc(func(params secret.PutSecretNameParams) middleware.Responder {
			return middleware.NotImplemented("operation SecretPutSecretName has not yet been implemented")
		}),
	}
}

/*SecretStoreAPI An API for managing secrets for serverless functions. */
type SecretStoreAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/com.vmware.vs.secrets.v1+json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/com.vmware.vs.secrets.v1+json" mime type
	JSONProducer runtime.Producer

	// SecretDeleteSecretNameHandler sets the operation handler for the delete secret name operation
	SecretDeleteSecretNameHandler secret.DeleteSecretNameHandler
	// SecretGetHandler sets the operation handler for the get operation
	SecretGetHandler secret.GetHandler
	// SecretGetSecretNameHandler sets the operation handler for the get secret name operation
	SecretGetSecretNameHandler secret.GetSecretNameHandler
	// SecretPostHandler sets the operation handler for the post operation
	SecretPostHandler secret.PostHandler
	// SecretPutSecretNameHandler sets the operation handler for the put secret name operation
	SecretPutSecretNameHandler secret.PutSecretNameHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *SecretStoreAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *SecretStoreAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *SecretStoreAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *SecretStoreAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *SecretStoreAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *SecretStoreAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *SecretStoreAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the SecretStoreAPI
func (o *SecretStoreAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.SecretDeleteSecretNameHandler == nil {
		unregistered = append(unregistered, "secret.DeleteSecretNameHandler")
	}

	if o.SecretGetHandler == nil {
		unregistered = append(unregistered, "secret.GetHandler")
	}

	if o.SecretGetSecretNameHandler == nil {
		unregistered = append(unregistered, "secret.GetSecretNameHandler")
	}

	if o.SecretPostHandler == nil {
		unregistered = append(unregistered, "secret.PostHandler")
	}

	if o.SecretPutSecretNameHandler == nil {
		unregistered = append(unregistered, "secret.PutSecretNameHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *SecretStoreAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *SecretStoreAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// Authorizer returns the registered authorizer
func (o *SecretStoreAPI) Authorizer() runtime.Authorizer {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *SecretStoreAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/com.vmware.vs.secrets.v1+json":
			result["application/com.vmware.vs.secrets.v1+json"] = o.JSONConsumer

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *SecretStoreAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/com.vmware.vs.secrets.v1+json":
			result["application/com.vmware.vs.secrets.v1+json"] = o.JSONProducer

		case "application/json":
			result["application/json"] = o.JSONProducer

		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *SecretStoreAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the secret store API
func (o *SecretStoreAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *SecretStoreAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/{secretName}"] = secret.NewDeleteSecretName(o.context, o.SecretDeleteSecretNameHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"][""] = secret.NewGet(o.context, o.SecretGetHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/{secretName}"] = secret.NewGetSecretName(o.context, o.SecretGetSecretNameHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"][""] = secret.NewPost(o.context, o.SecretPostHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/{secretName}"] = secret.NewPutSecretName(o.context, o.SecretPutSecretNameHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *SecretStoreAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middelware as you see fit
func (o *SecretStoreAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}