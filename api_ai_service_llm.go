package aiservice

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	aimodel "github.com/aliyun/aliyun-pai-aiservice-go-sdk/model"
)

// Linger please
var (
	_ context.Context
)

type AiServiceLlmApiService service

func (a *AiServiceLlmApiService) Create(model string, messages string) (*aimodel.LLMEvalJudgerResponse, error) {
	return a.CreateWithModelName(model, messages, "")
}
func (a *AiServiceLlmApiService) CreateWithModelName(model string, messages string, modelName string) (*aimodel.LLMEvalJudgerResponse, error) {
	var messageList []aimodel.LlmEvalJudgerRequestMessages
	if err := json.Unmarshal([]byte(messages), &messageList); err != nil {
		return nil, fmt.Errorf("messages is not valid json, %v", err)
	}

	request := &aimodel.LlmEvalJudgerRequest{
		Action:    "LLMEvalJudger",
		Model:     model,
		Messages:  messageList,
		ModelName: modelName,
	}

	response, httpResponse, err := a.LlmEvalJudgerPost(context.Background(), request)
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
AiServiceLlmApiService 裁判员模型服务
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body llm eval judger body

@return Response
*/
func (a *AiServiceLlmApiService) LlmEvalJudgerPost(ctx context.Context, body *aimodel.LlmEvalJudgerRequest) (*aimodel.LLMEvalJudgerResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue = aimodel.NewLLMEvalJudgerResponse()
	)

	path := "/api/v1/llm/eval_judger"

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + path

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body

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
