package marusia

import (
	"encoding/json"
	"github.com/ThCompiler/go_game_constractor/pkg/ginutilits"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HttpContext interface {
	GetHeader(headerName string) string
	SendErrorResponse(code int, errorText string)
	SetHeader(headerName string, value string)
	SendResponse(code int, response any)
	ParseRequest(req interface{}) error
}

type GinHttpContext struct {
	*gin.Context
}

func (gc *GinHttpContext) GetHeader(headerName string) string {
	return gc.Request.Header.Get(headerName)
}

func (gc *GinHttpContext) SendErrorResponse(code int, errorText string) {
	ginutilits.ErrorResponse(gc.Context, code, errorText)
}

func (gc *GinHttpContext) ParseRequest(req interface{}) error {
	return gc.ShouldBindJSON(req)
}

func (gc *GinHttpContext) SetHeader(headerName string, value string) {
	gc.Header(headerName, value)
}

func (gc *GinHttpContext) SendResponse(code int, response any) {
	gc.JSON(code, response)
}

type BaseHttpContext struct {
	Req  *http.Request
	Resp http.ResponseWriter
}

func (gc *BaseHttpContext) GetHeader(headerName string) string {
	return gc.Req.Header.Get(headerName)
}

func (gc *BaseHttpContext) SendErrorResponse(code int, errorText string) {
	http.Error(gc.Resp, errorText, code)
}

func (gc *BaseHttpContext) ParseRequest(req interface{}) error {
	return json.NewDecoder(gc.Req.Body).Decode(&req)
}

func (gc *BaseHttpContext) SetHeader(headerName string, value string) {
	gc.Resp.Header().Set(headerName, value)
}

func (gc *BaseHttpContext) SendResponse(code int, response any) {
	gc.Resp.WriteHeader(code)

	_ = json.NewEncoder(gc.Resp).Encode(response)
}
