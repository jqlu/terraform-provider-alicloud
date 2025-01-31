package gpdb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeResourceUsage invokes the gpdb.DescribeResourceUsage API synchronously
// api document: https://help.aliyun.com/api/gpdb/describeresourceusage.html
func (client *Client) DescribeResourceUsage(request *DescribeResourceUsageRequest) (response *DescribeResourceUsageResponse, err error) {
	response = CreateDescribeResourceUsageResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeResourceUsageWithChan invokes the gpdb.DescribeResourceUsage API asynchronously
// api document: https://help.aliyun.com/api/gpdb/describeresourceusage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeResourceUsageWithChan(request *DescribeResourceUsageRequest) (<-chan *DescribeResourceUsageResponse, <-chan error) {
	responseChan := make(chan *DescribeResourceUsageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeResourceUsage(request)
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

// DescribeResourceUsageWithCallback invokes the gpdb.DescribeResourceUsage API asynchronously
// api document: https://help.aliyun.com/api/gpdb/describeresourceusage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeResourceUsageWithCallback(request *DescribeResourceUsageRequest, callback func(response *DescribeResourceUsageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeResourceUsageResponse
		var err error
		defer close(result)
		response, err = client.DescribeResourceUsage(request)
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

// DescribeResourceUsageRequest is the request struct for api DescribeResourceUsage
type DescribeResourceUsageRequest struct {
	*requests.RpcRequest
	DBInstanceId string `position:"Query" name:"DBInstanceId"`
}

// DescribeResourceUsageResponse is the response struct for api DescribeResourceUsage
type DescribeResourceUsageResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	DBInstanceId string `json:"DBInstanceId" xml:"DBInstanceId"`
	Engine       string `json:"Engine" xml:"Engine"`
	DiskUsed     int    `json:"DiskUsed" xml:"DiskUsed"`
	DataSize     int    `json:"DataSize" xml:"DataSize"`
	LogSize      int    `json:"LogSize" xml:"LogSize"`
	BackupSize   int    `json:"BackupSize" xml:"BackupSize"`
}

// CreateDescribeResourceUsageRequest creates a request to invoke DescribeResourceUsage API
func CreateDescribeResourceUsageRequest() (request *DescribeResourceUsageRequest) {
	request = &DescribeResourceUsageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("gpdb", "2016-05-03", "DescribeResourceUsage", "gpdb", "openAPI")
	return
}

// CreateDescribeResourceUsageResponse creates a response to parse from DescribeResourceUsage response
func CreateDescribeResourceUsageResponse() (response *DescribeResourceUsageResponse) {
	response = &DescribeResourceUsageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
