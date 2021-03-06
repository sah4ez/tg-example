// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package transport

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (http *httpAdder) sum(ctx context.Context, request requestAdderSum) (response responseAdderSum, err error) {

	response.C, err = http.svc.Sum(ctx, request.AInt, request.BInt)
	if err != nil {
		if http.errorHandler != nil {
			err = http.errorHandler(err)
		}
	}
	return
}
func (http *httpAdder) serveSum(ctx *fiber.Ctx) (err error) {

	var request requestAdderSum
	ctx.Response().SetStatusCode(200)

	if _aInt := ctx.Query("a"); _aInt != "" {
		var aInt int
		aInt, err = strconv.Atoi(_aInt)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest)
			return sendResponse(ctx, "url arguments could not be decoded: "+err.Error())
		}
		request.AInt = aInt
	}
	if _bInt := ctx.Query("b"); _bInt != "" {
		var bInt int
		bInt, err = strconv.Atoi(_bInt)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest)
			return sendResponse(ctx, "url arguments could not be decoded: "+err.Error())
		}
		request.BInt = bInt
	}

	var response responseAdderSum
	if response, err = http.sum(ctx.UserContext(), request); err == nil {
		return sendResponse(ctx, response)
	}
	if errCoder, ok := err.(withErrorCode); ok {
		ctx.Status(errCoder.Code())
	} else {
		ctx.Status(fiber.StatusInternalServerError)
	}
	return sendResponse(ctx, err)
}
