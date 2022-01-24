// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package adder

const headerRequestID = "X-Request-Id"

type Option func(cli *ClientJsonRPC)

func DecodeError(decoder ErrorDecoder) Option {
	return func(cli *ClientJsonRPC) {
		cli.errorDecoder = decoder
	}
}

func Headers(headers ...string) Option {
	return func(cli *ClientJsonRPC) {
		cli.headers = headers
	}
}
