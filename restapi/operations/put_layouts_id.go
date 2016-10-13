package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	context "golang.org/x/net/context"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PutLayoutsIDHandlerFunc turns a function with the right signature into a put layouts ID handler
type PutLayoutsIDHandlerFunc func(context.Context, PutLayoutsIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutLayoutsIDHandlerFunc) Handle(ctx context.Context, params PutLayoutsIDParams) middleware.Responder {
	return fn(ctx, params)
}

// PutLayoutsIDHandler interface for that can handle valid put layouts ID params
type PutLayoutsIDHandler interface {
	Handle(context.Context, PutLayoutsIDParams) middleware.Responder
}

// NewPutLayoutsID creates a new http.Handler for the put layouts ID operation
func NewPutLayoutsID(ctx *middleware.Context, handler PutLayoutsIDHandler) *PutLayoutsID {
	return &PutLayoutsID{Context: ctx, Handler: handler}
}

/*PutLayoutsID swagger:route PUT /layouts/{id} putLayoutsId

Replace layout configuration.

*/
type PutLayoutsID struct {
	Context *middleware.Context
	Handler PutLayoutsIDHandler
}

func (o *PutLayoutsID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewPutLayoutsIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(context.Background(), Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}