package marusia

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ThCompiler/go_game_constractor/pkg/ginutilits"
)

type HTTPContext interface {
	GetHeader(headerName string) string
	SendErrorResponse(code int, errorText string)
	SetHeader(headerName, value string)
	SendResponse(code int, response any) error
	ParseRequest(req interface{}) error
	GetContext() context.Context
}

type GinHTTPContext struct {
	*gin.Context
}

func (gc *GinHTTPContext) GetHeader(headerName string) string {
	return gc.Request.Header.Get(headerName)
}

func (gc *GinHTTPContext) SendErrorResponse(code int, errorText string) {
	ginutilits.ErrorResponse(gc.Context, code, errorText)
}

func (gc *GinHTTPContext) ParseRequest(req interface{}) error {
	return gc.ShouldBindJSON(req)
}

func (gc *GinHTTPContext) SetHeader(headerName, value string) {
	gc.Header(headerName, value)
}

func (gc *GinHTTPContext) SendResponse(code int, response any) error {
	gc.JSON(code, response)

	return nil
}

func (gc *GinHTTPContext) GetContext() context.Context {
	return gc
}

type BaseHTTPContext struct {
	Req  *http.Request
	Resp http.ResponseWriter
}

func (bhc *BaseHTTPContext) GetHeader(headerName string) string {
	return bhc.Req.Header.Get(headerName)
}

func (bhc *BaseHTTPContext) SendErrorResponse(code int, errorText string) {
	http.Error(bhc.Resp, errorText, code)
}

func (bhc *BaseHTTPContext) ParseRequest(req interface{}) error {
	return json.NewDecoder(bhc.Req.Body).Decode(&req)
}

func (bhc *BaseHTTPContext) SetHeader(headerName, value string) {
	bhc.Resp.Header().Set(headerName, value)
}

func (bhc *BaseHTTPContext) SendResponse(code int, response any) error {
	bhc.Resp.WriteHeader(code)

	return json.NewEncoder(bhc.Resp).Encode(response)
}

func (bhc *BaseHTTPContext) GetContext() context.Context {
	return bhc.Req.Context()
}
