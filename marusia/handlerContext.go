package marusia

import (
    "context"
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
    GetContext() context.Context
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

func (gc *GinHttpContext) GetContext() context.Context {
    return gc
}

type BaseHttpContext struct {
    Req  *http.Request
    Resp http.ResponseWriter
}

func (bhc *BaseHttpContext) GetHeader(headerName string) string {
    return bhc.Req.Header.Get(headerName)
}

func (bhc *BaseHttpContext) SendErrorResponse(code int, errorText string) {
    http.Error(bhc.Resp, errorText, code)
}

func (bhc *BaseHttpContext) ParseRequest(req interface{}) error {
    return json.NewDecoder(bhc.Req.Body).Decode(&req)
}

func (bhc *BaseHttpContext) SetHeader(headerName string, value string) {
    bhc.Resp.Header().Set(headerName, value)
}

func (bhc *BaseHttpContext) SendResponse(code int, response any) {
    bhc.Resp.WriteHeader(code)

    _ = json.NewEncoder(bhc.Resp).Encode(response)
}

func (bhc *BaseHttpContext) GetContext() context.Context {
    return bhc.Req.Context()
}
