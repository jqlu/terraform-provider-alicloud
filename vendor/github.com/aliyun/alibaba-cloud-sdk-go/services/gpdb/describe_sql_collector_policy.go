package gpdb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeSQLCollectorPolicy invokes the gpdb.DescribeSQLCollectorPolicy API synchronously
// api document: https://help.aliyun.com/api/gpdb/describesqlcollectorpolicy.html
func (client *Client) DescribeSQLCollectorPolicy(request *DescribeSQLCollectorPolicyRequest) (response *DescribeSQLCollectorPolicyResponse, err error) {
	response = CreateDescribeSQLCollectorPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSQLCollectorPolicyWithChan invokes the gpdb.DescribeSQLCollectorPolicy API asynchronously
// api document: https://help.aliyun.com/api/gpdb/describesqlcollectorpolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSQLCollectorPolicyWithChan(request *DescribeSQLCollectorPolicyRequest) (<-chan *DescribeSQLCollectorPolicyResponse, <-chan error) {
	responseChan := make(chan *DescribeSQLCollectorPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSQLCollectorPolicy(request)
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

// DescribeSQLCollectorPolicyWithCallback invokes the gpdb.DescribeSQLCollectorPolicy API asynchronously
// api document: https://help.aliyun.com/api/gpdb/describesqlcollectorpolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSQLCollectorPolicyWithCallback(request *DescribeSQLCollectorPolicyRequest, callback func(response *DescribeSQLCollectorPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSQLCollectorPolicyResponse
		var err error
		defer close(result)
		response, err = client.DescribeSQLCollectorPolicy(request)
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

// DescribeSQLCollectorPolicyRequest is the request struct for api DescribeSQLCollectorPolicy
type DescribeSQLCollectorPolicyRequest struct {
	*requests.RpcRequest
	DBInstanceId string `position:"Query" name:"DBInstanceId"`
}

// DescribeSQLCollectorPolicyResponse is the response struct for api DescribeSQLCollectorPolicy
type DescribeSQLCollectorPolicyResponse struct {
	*responses.BaseResponse
	RequestId          string `json:"RequestId" xml:"RequestId"`
	SQLCollectorStatus string `json:"SQLCollectorStatus" xml:"SQLCollectorStatus"`
}

// CreateDescribeSQLCollectorPolicyRequest creates a request to invoke DescribeSQLCollectorPolicy API
func CreateDescribeSQLCollectorPolicyRequest() (request *DescribeSQLCollectorPolicyRequest) {
	request = &DescribeSQLCollectorPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("gpdb", "2016-05-03", "DescribeSQLCollectorPolicy", "gpdb", "openAPI")
	return
}

// CreateDescribeSQLCollectorPolicyResponse creates a response to parse from DescribeSQLCollectorPolicy response
func CreateDescribeSQLCollectorPolicyResponse() (response *DescribeSQLCollectorPolicyResponse) {
	response = &DescribeSQLCollectorPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
