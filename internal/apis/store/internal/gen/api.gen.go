// Package gen provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.2 DO NOT EDIT.
package gen

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/labstack/echo/v4"
)

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// Defines values for OrderItem.
const (
	TeaTableGreen OrderItem = "Tea Table Green"
	TeaTableRed   OrderItem = "Tea Table Red"
)

// Order defines model for Order.
type Order struct {
	Id    *string    `json:"id,omitempty"`
	Item  *OrderItem `json:"item,omitempty"`
	Price *int       `json:"price,omitempty"`
}

// OrderItem defines model for Order.Item.
type OrderItem string

// PutOrderIdJSONRequestBody defines body for PutOrderId for application/json ContentType.
type PutOrderIdJSONRequestBody = Order

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// TEst
	// (GET /order/{id})
	GetOrderId(ctx echo.Context, id string) error
	// Create an order
	// (PUT /order/{id})
	PutOrderId(ctx echo.Context, id interface{}) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetOrderId converts echo context to params.
func (w *ServerInterfaceWrapper) GetOrderId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{"read"})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetOrderId(ctx, id)
	return err
}

// PutOrderId converts echo context to params.
func (w *ServerInterfaceWrapper) PutOrderId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id interface{}

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{"read"})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PutOrderId(ctx, id)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/order/:id", wrapper.GetOrderId)
	router.PUT(baseURL+"/order/:id", wrapper.PutOrderId)

}

type GetOrderIdRequestObject struct {
	Id string `json:"id,omitempty"`
}

type GetOrderIdResponseObject interface {
	VisitGetOrderIdResponse(w http.ResponseWriter) error
}

type GetOrderId200JSONResponse Order

func (response GetOrderId200JSONResponse) VisitGetOrderIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PutOrderIdRequestObject struct {
	Id   interface{} `json:"id,omitempty"`
	Body *PutOrderIdJSONRequestBody
}

type PutOrderIdResponseObject interface {
	VisitPutOrderIdResponse(w http.ResponseWriter) error
}

type PutOrderId201Response struct {
}

func (response PutOrderId201Response) VisitPutOrderIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(201)
	return nil
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// TEst
	// (GET /order/{id})
	GetOrderId(ctx context.Context, request GetOrderIdRequestObject) (GetOrderIdResponseObject, error)
	// Create an order
	// (PUT /order/{id})
	PutOrderId(ctx context.Context, request PutOrderIdRequestObject) (PutOrderIdResponseObject, error)
}

type StrictHandlerFunc = runtime.StrictEchoHandlerFunc
type StrictMiddlewareFunc = runtime.StrictEchoMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// GetOrderId operation middleware
func (sh *strictHandler) GetOrderId(ctx echo.Context, id string) error {
	var request GetOrderIdRequestObject

	request.Id = id

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetOrderId(ctx.Request().Context(), request.(GetOrderIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetOrderId")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetOrderIdResponseObject); ok {
		return validResponse.VisitGetOrderIdResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("Unexpected response type: %T", response)
	}
	return nil
}

// PutOrderId operation middleware
func (sh *strictHandler) PutOrderId(ctx echo.Context, id interface{}) error {
	var request PutOrderIdRequestObject

	request.Id = id

	var body PutOrderIdJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PutOrderId(ctx.Request().Context(), request.(PutOrderIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PutOrderId")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PutOrderIdResponseObject); ok {
		return validResponse.VisitPutOrderIdResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("Unexpected response type: %T", response)
	}
	return nil
}