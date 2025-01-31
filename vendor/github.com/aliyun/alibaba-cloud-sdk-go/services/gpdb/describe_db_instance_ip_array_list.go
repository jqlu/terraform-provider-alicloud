package gpdb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeDBInstanceIPArrayList invokes the gpdb.DescribeDBInstanceIPArrayList API synchronously
// api document: https://help.aliyun.com/api/gpdb/describedbinstanceiparraylist.html
func (client *Client) DescribeDBInstanceIPArrayList(request *DescribeDBInstanceIPArrayListRequest) (response *DescribeDBInstanceIPArrayListResponse, err error) {
	response = CreateDescribeDBInstanceIPArrayListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDBInstanceIPArrayListWithChan invokes the gpdb.DescribeDBInstanceIPArrayList API asynchronously
// api document: https://help.aliyun.com/api/gpdb/describedbinstanceiparraylist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDBInstanceIPArrayListWithChan(request *DescribeDBInstanceIPArrayListRequest) (<-chan *DescribeDBInstanceIPArrayListResponse, <-chan error) {
	responseChan := make(chan *DescribeDBInstanceIPArrayListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBInstanceIPArrayList(request)
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

// DescribeDBInstanceIPArrayListWithCallback invokes the gpdb.DescribeDBInstanceIPArrayList API asynchronously
// api document: https://help.aliyun.com/api/gpdb/describedbinstanceiparraylist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDBInstanceIPArrayListWithCallback(request *DescribeDBInstanceIPArrayListRequest, callback func(response *DescribeDBInstanceIPArrayListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBInstanceIPArrayListResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBInstanceIPArrayList(request)
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

// DescribeDBInstanceIPArrayListRequest is the request struct for api DescribeDBInstanceIPArrayList
type DescribeDBInstanceIPArrayListRequest struct {
	*requests.RpcRequest
	DBInstanceId string `position:"Query" name:"DBInstanceId"`
}

// DescribeDBInstanceIPArrayListResponse is the response struct for api DescribeDBInstanceIPArrayList
type DescribeDBInstanceIPArrayListResponse struct {
	*responses.BaseResponse
	RequestId string                               `json:"RequestId" xml:"RequestId"`
	Items     ItemsInDescribeDBInstanceIPArrayList `json:"Items" xml:"Items"`
}

// CreateDescribeDBInstanceIPArrayListRequest creates a request to invoke DescribeDBInstanceIPArrayList API
func CreateDescribeDBInstanceIPArrayListRequest() (request *DescribeDBInstanceIPArrayListRequest) {
	request = &DescribeDBInstanceIPArrayListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("gpdb", "2016-05-03", "DescribeDBInstanceIPArrayList", "gpdb", "openAPI")
	return
}

// CreateDescribeDBInstanceIPArrayListResponse creates a response to parse from DescribeDBInstanceIPArrayList response
func CreateDescribeDBInstanceIPArrayListResponse() (response *DescribeDBInstanceIPArrayListResponse) {
	response = &DescribeDBInstanceIPArrayListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
