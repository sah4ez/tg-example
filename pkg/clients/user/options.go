// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package user

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
