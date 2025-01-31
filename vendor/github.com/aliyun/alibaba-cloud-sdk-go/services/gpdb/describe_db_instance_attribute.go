package gpdb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeDBInstanceAttribute invokes the gpdb.DescribeDBInstanceAttribute API synchronously
// api document: https://help.aliyun.com/api/gpdb/describedbinstanceattribute.html
func (client *Client) DescribeDBInstanceAttribute(request *DescribeDBInstanceAttributeRequest) (response *DescribeDBInstanceAttributeResponse, err error) {
	response = CreateDescribeDBInstanceAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDBInstanceAttributeWithChan invokes the gpdb.DescribeDBInstanceAttribute API asynchronously
// api document: https://help.aliyun.com/api/gpdb/describedbinstanceattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDBInstanceAttributeWithChan(request *DescribeDBInstanceAttributeRequest) (<-chan *DescribeDBInstanceAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeDBInstanceAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBInstanceAttribute(request)
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

// DescribeDBInstanceAttributeWithCallback invokes the gpdb.DescribeDBInstanceAttribute API asynchronously
// api document: https://help.aliyun.com/api/gpdb/describedbinstanceattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDBInstanceAttributeWithCallback(request *DescribeDBInstanceAttributeRequest, callback func(response *DescribeDBInstanceAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBInstanceAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBInstanceAttribute(request)
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

// DescribeDBInstanceAttributeRequest is the request struct for api DescribeDBInstanceAttribute
type DescribeDBInstanceAttributeRequest struct {
	*requests.RpcRequest
	DBInstanceId string           `position:"Query" name:"DBInstanceId"`
	OwnerId      requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeDBInstanceAttributeResponse is the response struct for api DescribeDBInstanceAttribute
type DescribeDBInstanceAttributeResponse struct {
	*responses.BaseResponse
	RequestId string                             `json:"RequestId" xml:"RequestId"`
	Items     ItemsInDescribeDBInstanceAttribute `json:"Items" xml:"Items"`
}

// CreateDescribeDBInstanceAttributeRequest creates a request to invoke DescribeDBInstanceAttribute API
func CreateDescribeDBInstanceAttributeRequest() (request *DescribeDBInstanceAttributeRequest) {
	request = &DescribeDBInstanceAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("gpdb", "2016-05-03", "DescribeDBInstanceAttribute", "gpdb", "openAPI")
	return
}

// CreateDescribeDBInstanceAttributeResponse creates a response to parse from DescribeDBInstanceAttribute response
func CreateDescribeDBInstanceAttributeResponse() (response *DescribeDBInstanceAttributeResponse) {
	response = &DescribeDBInstanceAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
