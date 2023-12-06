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

type AiServiceTextApiService service

func (a *AiServiceTextApiService) KeyWordText(text, modelName string, config map[string]interface{}) (*model.Response, error) {
	textRequest := &model.TextRequest{
		Action:    "KeyWordText",
		Text:      text,
		ModelName: modelName,
		Configure: config,
	}

	response, httpResponse, err := a.AiModelText(context.Background(), textRequest)
	if err != nil {
		return response, err
	}

	if httpResponse.StatusCode != 200 {
		err = fmt.Errorf("http status code error:%d", httpResponse.StatusCode)
		return response, err
	}

	return response, nil
}

func (a *AiServiceTextApiService) NERText(text, modelName string, config map[string]interface{}) (*model.Response, error) {
	textRequest := &model.TextRequest{
		Action:    "NERText",
		Text:      text,
		ModelName: modelName,
		Configure: config,
	}

	response, httpResponse, err := a.AiModelText(context.Background(), textRequest)
	if err != nil {
		return response, err
	}

	if httpResponse.StatusCode != 200 {
		err = fmt.Errorf("http status code error:%d", httpResponse.StatusCode)
		return response, err
	}

	return response, nil
}

func (a *AiServiceTextApiService) FineGrainedLabelText(text, modelName string, config map[string]interface{}) (*model.Response, error) {
	textRequest := &model.TextRequest{
		Action:    "FineGrainedLabelText",
		Text:      text,
		ModelName: modelName,
		Configure: config,
	}

	response, httpResponse, err := a.AiModelText(context.Background(), textRequest)
	if err != nil {
		return response, err
	}

	if httpResponse.StatusCode != 200 {
		err = fmt.Errorf("http status code error:%d", httpResponse.StatusCode)
		return response, err
	}

	return response, nil
}

func (a *AiServiceTextApiService) VectorizationText(text, modelName string, config map[string]interface{}) (*model.Response, error) {
	textRequest := &model.TextRequest{
		Action:    "VectorizationText",
		Text:      text,
		ModelName: modelName,
		Configure: config,
	}

	response, httpResponse, err := a.AiModelText(context.Background(), textRequest)
	if err != nil {
		return response, err
	}

	if httpResponse.StatusCode != 200 {
		err = fmt.Errorf("http status code error:%d", httpResponse.StatusCode)
		return response, err
	}

	return response, nil
}

func (a *AiServiceTextApiService) AntiSpamText(text, modelName string, config map[string]interface{}) (*model.Response, error) {
	textRequest := &model.TextRequest{
		Action:    "AntiSpamText",
		Text:      text,
		ModelName: modelName,
		Configure: config,
	}

	response, httpResponse, err := a.AiModelText(context.Background(), textRequest)
	if err != nil {
		return response, err
	}

	if httpResponse.StatusCode != 200 {
		err = fmt.Errorf("http status code error:%d", httpResponse.StatusCode)
		return response, err
	}

	return response, nil
}

func (a *AiServiceTextApiService) SentimentAnalysisText(text, modelName string, config map[string]interface{}) (*model.Response, error) {
	textRequest := &model.TextRequest{
		Action:    "SentimentAnalysisText",
		Text:      text,
		ModelName: modelName,
		Configure: config,
	}

	response, httpResponse, err := a.AiModelText(context.Background(), textRequest)
	if err != nil {
		return response, err
	}

	if httpResponse.StatusCode != 200 {
		err = fmt.Errorf("http status code error:%d", httpResponse.StatusCode)
		return response, err
	}

	return response, nil
}

func (a *AiServiceTextApiService) EcomerceReviewSentimentAnalysisText(text, modelName string, config map[string]interface{}) (*model.Response, error) {
	textRequest := &model.TextRequest{
		Action:    "EcomerceReviewSentimentAnalysisText",
		Text:      text,
		ModelName: modelName,
		Configure: config,
	}

	response, httpResponse, err := a.AiModelText(context.Background(), textRequest)
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
AiServiceTextApiService 文本检测接口
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body 文本检测body

@return Response
*/
func (a *AiServiceTextApiService) AiModelText(ctx context.Context, body *model.TextRequest) (*model.Response, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue = model.NewResponse()
	)

	path := "/api/v1/text"
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
