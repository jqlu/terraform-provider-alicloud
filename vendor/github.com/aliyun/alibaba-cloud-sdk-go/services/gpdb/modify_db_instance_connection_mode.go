package gpdb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ModifyDBInstanceConnectionMode invokes the gpdb.ModifyDBInstanceConnectionMode API synchronously
// api document: https://help.aliyun.com/api/gpdb/modifydbinstanceconnectionmode.html
func (client *Client) ModifyDBInstanceConnectionMode(request *ModifyDBInstanceConnectionModeRequest) (response *ModifyDBInstanceConnectionModeResponse, err error) {
	response = CreateModifyDBInstanceConnectionModeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDBInstanceConnectionModeWithChan invokes the gpdb.ModifyDBInstanceConnectionMode API asynchronously
// api document: https://help.aliyun.com/api/gpdb/modifydbinstanceconnectionmode.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDBInstanceConnectionModeWithChan(request *ModifyDBInstanceConnectionModeRequest) (<-chan *ModifyDBInstanceConnectionModeResponse, <-chan error) {
	responseChan := make(chan *ModifyDBInstanceConnectionModeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDBInstanceConnectionMode(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ModifyDBInstanceConnectionModeWithCallback invokes the gpdb.ModifyDBInstanceConnectionMode API asynchronously
// api document: https://help.aliyun.com/api/gpdb/modifydbinstanceconnectionmode.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDBInstanceConnectionModeWithCallback(request *ModifyDBInstanceConnectionModeRequest, callback func(response *ModifyDBInstanceConnectionModeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDBInstanceConnectionModeResponse
		var err error
		defer close(result)
		response, err = client.ModifyDBInstanceConnectionMode(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ModifyDBInstanceConnectionModeRequest is the request struct for api ModifyDBInstanceConnectionMode
type ModifyDBInstanceConnectionModeRequest struct {
	*requests.RpcRequest
	ConnectionMode string `position:"Query" name:"ConnectionMode"`
	DBInstanceId   string `position:"Query" name:"DBInstanceId"`
}

// ModifyDBInstanceConnectionModeResponse is the response struct for api ModifyDBInstanceConnectionMode
type ModifyDBInstanceConnectionModeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDBInstanceConnectionModeRequest creates a request to invoke ModifyDBInstanceConnectionMode API
func CreateModifyDBInstanceConnectionModeRequest() (request *ModifyDBInstanceConnectionModeRequest) {
	request = &ModifyDBInstanceConnectionModeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("gpdb", "2016-05-03", "ModifyDBInstanceConnectionMode", "gpdb", "openAPI")
	return
}

// CreateModifyDBInstanceConnectionModeResponse creates a response to parse from ModifyDBInstanceConnectionMode response
func CreateModifyDBInstanceConnectionModeResponse() (response *ModifyDBInstanceConnectionModeResponse) {
	response = &ModifyDBInstanceConnectionModeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
