package gpdb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeSlowLogRecords invokes the gpdb.DescribeSlowLogRecords API synchronously
// api document: https://help.aliyun.com/api/gpdb/describeslowlogrecords.html
func (client *Client) DescribeSlowLogRecords(request *DescribeSlowLogRecordsRequest) (response *DescribeSlowLogRecordsResponse, err error) {
	response = CreateDescribeSlowLogRecordsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSlowLogRecordsWithChan invokes the gpdb.DescribeSlowLogRecords API asynchronously
// api document: https://help.aliyun.com/api/gpdb/describeslowlogrecords.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSlowLogRecordsWithChan(request *DescribeSlowLogRecordsRequest) (<-chan *DescribeSlowLogRecordsResponse, <-chan error) {
	responseChan := make(chan *DescribeSlowLogRecordsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSlowLogRecords(request)
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

// DescribeSlowLogRecordsWithCallback invokes the gpdb.DescribeSlowLogRecords API asynchronously
// api document: https://help.aliyun.com/api/gpdb/describeslowlogrecords.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSlowLogRecordsWithCallback(request *DescribeSlowLogRecordsRequest, callback func(response *DescribeSlowLogRecordsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSlowLogRecordsResponse
		var err error
		defer close(result)
		response, err = client.DescribeSlowLogRecords(request)
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

// DescribeSlowLogRecordsRequest is the request struct for api DescribeSlowLogRecords
type DescribeSlowLogRecordsRequest struct {
	*requests.RpcRequest
	SQLId        requests.Integer `position:"Query" name:"SQLId"`
	DBName       string           `position:"Query" name:"DBName"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
	EndTime      string           `position:"Query" name:"EndTime"`
	DBInstanceId string           `position:"Query" name:"DBInstanceId"`
	StartTime    string           `position:"Query" name:"StartTime"`
	PageNumber   requests.Integer `position:"Query" name:"PageNumber"`
}

// DescribeSlowLogRecordsResponse is the response struct for api DescribeSlowLogRecords
type DescribeSlowLogRecordsResponse struct {
	*responses.BaseResponse
	RequestId        string                        `json:"RequestId" xml:"RequestId"`
	Engine           string                        `json:"Engine" xml:"Engine"`
	TotalRecordCount int                           `json:"TotalRecordCount" xml:"TotalRecordCount"`
	PageNumber       int                           `json:"PageNumber" xml:"PageNumber"`
	PageRecordCount  int                           `json:"PageRecordCount" xml:"PageRecordCount"`
	Items            ItemsInDescribeSlowLogRecords `json:"Items" xml:"Items"`
}

// CreateDescribeSlowLogRecordsRequest creates a request to invoke DescribeSlowLogRecords API
func CreateDescribeSlowLogRecordsRequest() (request *DescribeSlowLogRecordsRequest) {
	request = &DescribeSlowLogRecordsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("gpdb", "2016-05-03", "DescribeSlowLogRecords", "gpdb", "openAPI")
	return
}

// CreateDescribeSlowLogRecordsResponse creates a response to parse from DescribeSlowLogRecords response
func CreateDescribeSlowLogRecordsResponse() (response *DescribeSlowLogRecordsResponse) {
	response = &DescribeSlowLogRecordsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
