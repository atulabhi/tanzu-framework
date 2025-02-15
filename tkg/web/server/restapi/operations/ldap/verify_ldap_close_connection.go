// Code generated by go-swagger; DO NOT EDIT.

package ldap

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// VerifyLdapCloseConnectionHandlerFunc turns a function with the right signature into a verify ldap close connection handler
type VerifyLdapCloseConnectionHandlerFunc func(VerifyLdapCloseConnectionParams) middleware.Responder

// Handle executing the request and returning a response
func (fn VerifyLdapCloseConnectionHandlerFunc) Handle(params VerifyLdapCloseConnectionParams) middleware.Responder {
	return fn(params)
}

// VerifyLdapCloseConnectionHandler interface for that can handle valid verify ldap close connection params
type VerifyLdapCloseConnectionHandler interface {
	Handle(VerifyLdapCloseConnectionParams) middleware.Responder
}

// NewVerifyLdapCloseConnection creates a new http.Handler for the verify ldap close connection operation
func NewVerifyLdapCloseConnection(ctx *middleware.Context, handler VerifyLdapCloseConnectionHandler) *VerifyLdapCloseConnection {
	return &VerifyLdapCloseConnection{Context: ctx, Handler: handler}
}

/*
VerifyLdapCloseConnection swagger:route POST /api/ldap/disconnect ldap verifyLdapCloseConnection

Validate if the LDAP connection can be closed
*/
type VerifyLdapCloseConnection struct {
	Context *middleware.Context
	Handler VerifyLdapCloseConnectionHandler
}

func (o *VerifyLdapCloseConnection) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewVerifyLdapCloseConnectionParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
