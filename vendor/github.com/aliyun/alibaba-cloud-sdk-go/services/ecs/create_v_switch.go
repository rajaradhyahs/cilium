package ecs

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// CreateVSwitch invokes the ecs.CreateVSwitch API synchronously
func (client *Client) CreateVSwitch(request *CreateVSwitchRequest) (response *CreateVSwitchResponse, err error) {
	response = CreateCreateVSwitchResponse()
	err = client.DoAction(request, response)
	return
}

// CreateVSwitchWithChan invokes the ecs.CreateVSwitch API asynchronously
func (client *Client) CreateVSwitchWithChan(request *CreateVSwitchRequest) (<-chan *CreateVSwitchResponse, <-chan error) {
	responseChan := make(chan *CreateVSwitchResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateVSwitch(request)
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

// CreateVSwitchWithCallback invokes the ecs.CreateVSwitch API asynchronously
func (client *Client) CreateVSwitchWithCallback(request *CreateVSwitchRequest, callback func(response *CreateVSwitchResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateVSwitchResponse
		var err error
		defer close(result)
		response, err = client.CreateVSwitch(request)
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

// CreateVSwitchRequest is the request struct for api CreateVSwitch
type CreateVSwitchRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	Description          string           `position:"Query" name:"Description"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	VpcId                string           `position:"Query" name:"VpcId"`
	VSwitchName          string           `position:"Query" name:"VSwitchName"`
	CidrBlock            string           `position:"Query" name:"CidrBlock"`
	ZoneId               string           `position:"Query" name:"ZoneId"`
}

// CreateVSwitchResponse is the response struct for api CreateVSwitch
type CreateVSwitchResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	VSwitchId string `json:"VSwitchId" xml:"VSwitchId"`
}

// CreateCreateVSwitchRequest creates a request to invoke CreateVSwitch API
func CreateCreateVSwitchRequest() (request *CreateVSwitchRequest) {
	request = &CreateVSwitchRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "CreateVSwitch", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateVSwitchResponse creates a response to parse from CreateVSwitch response
func CreateCreateVSwitchResponse() (response *CreateVSwitchResponse) {
	response = &CreateVSwitchResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}