// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package user

import (
	"context"
	"encoding/json"

	goUUID "github.com/google/uuid"
)

type ClientUser struct {
	*ClientJsonRPC
}

type retUserGetUserNameByID func(name string, err error)

func (cli *ClientUser) ReqGetUserNameByID(ret retUserGetUserNameByID, id int) (request baseJsonRPC) {

	request = baseJsonRPC{
		Method:  "user.getusernamebyid",
		Params:  requestUserGetUserNameByID{Id: id},
		Version: Version,
	}
	var err error
	var response responseUserGetUserNameByID

	if ret != nil {
		request.retHandler = func(jsonrpcResponse baseJsonRPC) {
			if jsonrpcResponse.Error != nil {
				err = cli.errorDecoder(jsonrpcResponse.Error)
				ret(response.Name, err)
				return
			}
			err = json.Unmarshal(jsonrpcResponse.Result, &response)
			ret(response.Name, err)
		}
		request.ID = []byte("\"" + goUUID.New().String() + "\"")
	}
	return
}

func (cli *ClientUser) GetUserNameByID(ctx context.Context, id int) (name string, err error) {

	retHandler := func(_name string, _err error) {
		name = _name
		err = _err
	}
	if blockErr := cli.Batch(ctx, cli.ReqGetUserNameByID(retHandler, id)); blockErr != nil {
		err = blockErr
		return
	}
	return
}
