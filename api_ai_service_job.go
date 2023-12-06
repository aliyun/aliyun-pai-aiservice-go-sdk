package aiservice

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/aliyun/aliyun-pai-aiservice-go-sdk/model"
)

// Linger please
var (
	_ context.Context
)

type AiServiceJobApiService service

func (a *AiServiceJobApiService) GetAsyncJobWithId(jobId int32) (*model.AsyncJobResponse, error) {
	response, httpResponse, err := a.GetAsyncJob(jobId)
	if err != nil {
		return response, err
	}

	if httpResponse.StatusCode != 200 {
		err = fmt.Errorf("http status code error:%d", httpResponse.StatusCode)
		return response, err
	}

	return response, nil
}

func (a *AiServiceJobApiService) GetBatchJobWithId(batchId int32) (*model.Response, error) {
	response, httpResponse, err := a.GetBatchJob(batchId)
	if err != nil {
		return response, err
	}

	if httpResponse.StatusCode != 200 {
		err = fmt.Errorf("http status code error:%d", httpResponse.StatusCode)
		return response, err
	}

	return response, nil
}

/*
AiServiceJobApiService 根据 jobid 查询异步任务
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param jobId job id

@return Response
*/
func (a *AiServiceJobApiService) GetAsyncJob(jobId int32) (*model.AsyncJobResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue = model.NewAsyncJobResponse()
	)

	ctx := context.Background()
	path := "/api/v1/async_jobs/" + fmt.Sprintf("%v", jobId)
	// create path and map variables
	localVarPath := a.client.cfg.BasePath + path

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"text/plain"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}

	sign := a.client.createSign(a.cfg, path, localVarHttpMethod, localVarPostBody, localVarQueryParams)
	localVarHeaderParams["Authorization"] = a.cfg.AppID + ":" + sign
	localVarHeaderParams["X-Request-Time"] = strconv.Itoa(int(time.Now().Unix()))

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	err = a.client.decode(localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
AiServiceJobApiService 根据 batch id 查询异步任务
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param batchId batch id

@return Response
*/
func (a *AiServiceJobApiService) GetBatchJob(batchId int32) (*model.Response, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue = model.NewResponse()
	)

	ctx := context.Background()
	path := "/api/v1/batch_jobs/" + fmt.Sprintf("%v", batchId)
	// create path and map variables
	localVarPath := a.client.cfg.BasePath + path

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"text/plain"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}

	sign := a.client.createSign(a.cfg, path, localVarHttpMethod, localVarPostBody, localVarQueryParams)
	localVarHeaderParams["Authorization"] = a.cfg.AppID + ":" + sign
	localVarHeaderParams["X-Request-Time"] = strconv.Itoa(int(time.Now().Unix()))

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	err = a.client.decode(localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
