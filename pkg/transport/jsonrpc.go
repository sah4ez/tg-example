// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package transport

import (
	"strings"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/seniorGolang/json"
)

const (
	maxParallelBatch = 100
	// Version defines the version of the JSON RPC implementation
	Version = "2.0"
	// contentTypeJson defines the content type to be served
	contentTypeJson = "application/json"
	// ParseError defines invalid JSON was received by the server
	// An error occurred on the server while parsing the JSON text
	parseError = -32700
	// InvalidRequestError defines the JSON sent is not a valid Request object
	invalidRequestError = -32600
	// MethodNotFoundError defines the method does not exist / is not available
	methodNotFoundError = -32601
	// InvalidParamsError defines invalid method parameter(s)
	invalidParamsError = -32602
	// InternalError defines a server error
	internalError = -32603
)

type idJsonRPC = json.RawMessage

type baseJsonRPC struct {
	ID      idJsonRPC       `json:"id"`
	Version string          `json:"jsonrpc"`
	Method  string          `json:"method,omitempty"`
	Error   *errorJsonRPC   `json:"error,omitempty"`
	Params  json.RawMessage `json:"params,omitempty"`
	Result  json.RawMessage `json:"result,omitempty"`
}

type errorJsonRPC struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (err errorJsonRPC) Error() string {
	return err.Message
}

type jsonrpcResponses []baseJsonRPC

func (responses *jsonrpcResponses) append(response *baseJsonRPC) {
	if response == nil {
		return
	}
	if response.ID != nil {
		*responses = append(*responses, *response)
	}
}

type methodJsonRPC func(ctx *fiber.Ctx, requestBase baseJsonRPC) (responseBase *baseJsonRPC)

func (srv *Server) serveBatch(ctx *fiber.Ctx) (err error) {
	methodHTTP := ctx.Method()
	if methodHTTP != fiber.MethodPost {
		ctx.Response().SetStatusCode(fiber.StatusMethodNotAllowed)
		if _, err = ctx.WriteString("only POST method supported"); err != nil {
			return
		}
		return
	}
	if value := ctx.Context().Value(CtxCancelRequest); value != nil {
		return
	}
	var single bool
	var requests []baseJsonRPC
	if err = json.Unmarshal(ctx.Body(), &requests); err != nil {
		var request baseJsonRPC
		if err = json.Unmarshal(ctx.Body(), &request); err != nil {
			return sendResponse(srv.log, ctx, makeErrorResponseJsonRPC([]byte("\"0\""), parseError, "request body could not be decoded: "+err.Error(), nil))
		}
		single = true
		requests = append(requests, request)
	}
	responses := make(jsonrpcResponses, 0, len(requests))
	var n int
	var wg sync.WaitGroup
	for _, request := range requests {
		methodNameOrigin := request.Method
		method := strings.ToLower(request.Method)
		switch method {

		case "adder.add":
			wg.Add(1)
			go func(request baseJsonRPC) {
				if request.ID != nil {
					responses.append(srv.httpAdder.add(ctx, request))
					wg.Done()
					return
				}
				srv.httpAdder.add(ctx, request)
				wg.Done()
			}(request)
		default:
			responses.append(makeErrorResponseJsonRPC(request.ID, methodNotFoundError, "invalid method '"+methodNameOrigin+"'", nil))
		}
		if n > maxParallelBatch {
			n = 0
			wg.Wait()
		}
		n++
	}
	wg.Wait()
	if single {
		return sendResponse(srv.log, ctx, responses[0])
	}
	return sendResponse(srv.log, ctx, responses)
}

func makeErrorResponseJsonRPC(id idJsonRPC, code int, msg string, data interface{}) *baseJsonRPC {

	if id == nil {
		return nil
	}

	return &baseJsonRPC{
		Error: &errorJsonRPC{
			Code:    code,
			Data:    data,
			Message: msg,
		},
		ID:      id,
		Version: Version,
	}
}
