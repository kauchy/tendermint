package core

import (
	abci "github.com/tendermint/tendermint/abci/types"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"github.com/tendermint/tendermint/version"
	cmn "github.com/tendermint/tmlibs/common"
)

// Query the application for some information.
//
// ```shell
// curl 'localhost:26657/abci_query?path=""&data="abcd"&prove=false'
// ```
//
// ```go
// client := client.NewHTTP("tcp://0.0.0.0:26657", "/websocket")
// result, err := client.ABCIQuery("", "abcd", true)
// ```
//
// > The above command returns JSON structured like this:
//
// ```json
// {
// 	"error": "",
// 	"result": {
// 		"response": {
// 			"log": "exists",
// 			"height": 0,
// 			"value": "61626364",
// 			"key": "61626364",
// 			"index": -1,
// 			"code": 0
// 		}
// 	},
// 	"id": "",
// 	"jsonrpc": "2.0"
// }
// ```
//
// ### Query Parameters
//
// | Parameter | Type   | Default | Required | Description                                    |
// |-----------+--------+---------+----------+------------------------------------------------|
// | path      | string | false   | false    | Path to the data ("/a/b/c")                    |
// | data      | []byte | false   | true     | Data                                           |
// | height    | int64  | 0       | false    | Height (0 means latest)                        |
// | prove     | bool   | false   | false    | Includes proof if true                         |
func ABCIQuery(path string, data cmn.HexBytes, height int64, prove bool) (*ctypes.ResultABCIQuery, error) {
	resQuery, err := proxyAppQuery.QuerySync(abci.RequestQuery{
		Path:   path,
		Data:   data,
		Height: height,
		Prove:  prove,
	})
	if err != nil {
		return nil, err
	}
	logger.Info("ABCIQuery", "path", path, "data", data, "result", resQuery)
	return &ctypes.ResultABCIQuery{*resQuery}, nil
}

// Get some info about the application.
//
// ```shell
// curl 'localhost:26657/abci_info'
// ```
//
// ```go
// client := client.NewHTTP("tcp://0.0.0.0:26657", "/websocket")
// info, err := client.ABCIInfo()
// ```
//
// > The above command returns JSON structured like this:
//
// ```json
// {
// 	"error": "",
// 	"result": {
// 		"response": {
// 			"data": "{\"size\":3}"
// 		}
// 	},
// 	"id": "",
// 	"jsonrpc": "2.0"
// }
// ```
func ABCIInfo() (*ctypes.ResultABCIInfo, error) {
	resInfo, err := proxyAppQuery.InfoSync(abci.RequestInfo{version.Version})
	if err != nil {
		return nil, err
	}
	return &ctypes.ResultABCIInfo{*resInfo}, nil
}
