package gpdb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ReleaseInstancePublicConnection invokes the gpdb.ReleaseInstancePublicConnection API synchronously
// api document: https://help.aliyun.com/api/gpdb/releaseinstancepublicconnection.html
func (client *Client) ReleaseInstancePublicConnection(request *ReleaseInstancePublicConnectionRequest) (response *ReleaseInstancePublicConnectionResponse, err error) {
	response = CreateReleaseInstancePublicConnectionResponse()
	err = client.DoAction(request, response)
	return
}

// ReleaseInstancePublicConnectionWithChan invokes the gpdb.ReleaseInstancePublicConnection API asynchronously
// api document: https://help.aliyun.com/api/gpdb/releaseinstancepublicconnection.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ReleaseInstancePublicConnectionWithChan(request *ReleaseInstancePublicConnectionRequest) (<-chan *ReleaseInstancePublicConnectionResponse, <-chan error) {
	responseChan := make(chan *ReleaseInstancePublicConnectionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReleaseInstancePublicConnection(request)
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

// ReleaseInstancePublicConnectionWithCallback invokes the gpdb.ReleaseInstancePublicConnection API asynchronously
// api document: https://help.aliyun.com/api/gpdb/releaseinstancepublicconnection.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ReleaseInstancePublicConnectionWithCallback(request *ReleaseInstancePublicConnectionRequest, callback func(response *ReleaseInstancePublicConnectionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReleaseInstancePublicConnectionResponse
		var err error
		defer close(result)
		response, err = client.ReleaseInstancePublicConnection(request)
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

// ReleaseInstancePublicConnectionRequest is the request struct for api ReleaseInstancePublicConnection
type ReleaseInstancePublicConnectionRequest struct {
	*requests.RpcRequest
	DBInstanceId            string `position:"Query" name:"DBInstanceId"`
	CurrentConnectionString string `position:"Query" name:"CurrentConnectionString"`
}

// ReleaseInstancePublicConnectionResponse is the response struct for api ReleaseInstancePublicConnection
type ReleaseInstancePublicConnectionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateReleaseInstancePublicConnectionRequest creates a request to invoke ReleaseInstancePublicConnection API
func CreateReleaseInstancePublicConnectionRequest() (request *ReleaseInstancePublicConnectionRequest) {
	request = &ReleaseInstancePublicConnectionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("gpdb", "2016-05-03", "ReleaseInstancePublicConnection", "gpdb", "openAPI")
	return
}

// CreateReleaseInstancePublicConnectionResponse creates a response to parse from ReleaseInstancePublicConnection response
func CreateReleaseInstancePublicConnectionResponse() (response *ReleaseInstancePublicConnectionResponse) {
	response = &ReleaseInstancePublicConnectionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
