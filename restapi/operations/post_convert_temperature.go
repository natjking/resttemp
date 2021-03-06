package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostConvertTemperatureHandlerFunc turns a function with the right signature into a post convert temperature handler
type PostConvertTemperatureHandlerFunc func(PostConvertTemperatureParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostConvertTemperatureHandlerFunc) Handle(params PostConvertTemperatureParams) middleware.Responder {
	return fn(params)
}

// PostConvertTemperatureHandler interface for that can handle valid post convert temperature params
type PostConvertTemperatureHandler interface {
	Handle(PostConvertTemperatureParams) middleware.Responder
}

// NewPostConvertTemperature creates a new http.Handler for the post convert temperature operation
func NewPostConvertTemperature(ctx *middleware.Context, handler PostConvertTemperatureHandler) *PostConvertTemperature {
	return &PostConvertTemperature{Context: ctx, Handler: handler}
}

/*PostConvertTemperature swagger:route POST /convertTemperature postConvertTemperature

PostConvertTemperature post convert temperature API

*/
type PostConvertTemperature struct {
	Context *middleware.Context
	Handler PostConvertTemperatureHandler
}

func (o *PostConvertTemperature) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewPostConvertTemperatureParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
