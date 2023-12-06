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

type AiServiceImageApiService service

func (a *AiServiceImageApiService) OCRImage(image, modelName string, config map[string]interface{}) (*model.Response, error) {
	request := &model.ImageRequest{
		Action:    "OCRImage",
		Image:     image,
		ModelName: modelName,
		Configure: config,
	}

	response, httpResponse, err := a.AiModelImage(context.Background(), request)
	if err != nil {
		return response, err
	}

	if httpResponse.StatusCode != 200 {
		err = fmt.Errorf("http status code error:%d", httpResponse.StatusCode)
		return response, err
	}

	return response, nil
}

func (a *AiServiceImageApiService) ClassifyImage(image, modelName string, config map[string]interface{}) (*model.Response, error) {
	request := &model.ImageRequest{
		Action:    "ClassifyImage",
		Image:     image,
		ModelName: modelName,
		Configure: config,
	}

	response, httpResponse, err := a.AiModelImage(context.Background(), request)
	if err != nil {
		return response, err
	}

	if httpResponse.StatusCode != 200 {
		err = fmt.Errorf("http status code error:%d", httpResponse.StatusCode)
		return response, err
	}

	return response, nil
}

func (a *AiServiceImageApiService) MultiLabelImage(image, modelName string, config map[string]interface{}) (*model.Response, error) {
	request := &model.ImageRequest{
		Action:    "MultiLabelImage",
		Image:     image,
		ModelName: modelName,
		Configure: config,
	}

	response, httpResponse, err := a.AiModelImage(context.Background(), request)
	if err != nil {
		return response, err
	}

	if httpResponse.StatusCode != 200 {
		err = fmt.Errorf("http status code error:%d", httpResponse.StatusCode)
		return response, err
	}

	return response, nil
}

func (a *AiServiceImageApiService) MultiLabelImageV2(image, modelName string, config map[string]interface{}) (*model.Response, error) {
	request := &model.ImageRequest{
		Action:    "MultiLabelImageV2",
		Image:     image,
		ModelName: modelName,
		Configure: config,
	}

	response, httpResponse, err := a.AiModelImage(context.Background(), request)
	if err != nil {
		return response, err
	}

	if httpResponse.StatusCode != 200 {
		err = fmt.Errorf("http status code error:%d", httpResponse.StatusCode)
		return response, err
	}

	return response, nil
}

func (a *AiServiceImageApiService) DetectClassifyImage(image, modelName string, config map[string]interface{}) (*model.Response, error) {
	request := &model.ImageRequest{
		Action:    "DetectClassifyImage",
		Image:     image,
		ModelName: modelName,
		Configure: config,
	}

	response, httpResponse, err := a.AiModelImage(context.Background(), request)
	if err != nil {
		return response, err
	}

	if httpResponse.StatusCode != 200 {
		err = fmt.Errorf("http status code error:%d", httpResponse.StatusCode)
		return response, err
	}

	return response, nil
}

func (a *AiServiceImageApiService) FaceRecognitionImage(image, modelName string, config map[string]interface{}) (*model.Response, error) {
	request := &model.ImageRequest{
		Action:    "FaceRecognitionImage",
		Image:     image,
		ModelName: modelName,
		Configure: config,
	}

	response, httpResponse, err := a.AiModelImage(context.Background(), request)
	if err != nil {
		return response, err
	}

	if httpResponse.StatusCode != 200 {
		err = fmt.Errorf("http status code error:%d", httpResponse.StatusCode)
		return response, err
	}

	return response, nil
}

func (a *AiServiceImageApiService) ScanImage(image, modelName string, config map[string]interface{}) (*model.Response, error) {
	request := &model.ImageRequest{
		Action:    "ScanImage",
		Image:     image,
		ModelName: modelName,
		Configure: config,
	}

	response, httpResponse, err := a.AiModelImage(context.Background(), request)
	if err != nil {
		return response, err
	}

	if httpResponse.StatusCode != 200 {
		err = fmt.Errorf("http status code error:%d", httpResponse.StatusCode)
		return response, err
	}

	return response, nil
}

func (a *AiServiceImageApiService) ProductDetectImage(image, modelName string, config map[string]interface{}) (*model.Response, error) {
	request := &model.ImageRequest{
		Action:    "ProductDetectImage",
		Image:     image,
		ModelName: modelName,
		Configure: config,
	}

	response, httpResponse, err := a.AiModelImage(context.Background(), request)
	if err != nil {
		return response, err
	}

	if httpResponse.StatusCode != 200 {
		err = fmt.Errorf("http status code error:%d", httpResponse.StatusCode)
		return response, err
	}

	return response, nil
}

func (a *AiServiceImageApiService) FaceDetectImage(image, modelName string, config map[string]interface{}) (*model.Response, error) {
	request := &model.ImageRequest{
		Action:    "FaceDetectImage",
		Image:     image,
		ModelName: modelName,
		Configure: config,
	}

	response, httpResponse, err := a.AiModelImage(context.Background(), request)
	if err != nil {
		return response, err
	}

	if httpResponse.StatusCode != 200 {
		err = fmt.Errorf("http status code error:%d", httpResponse.StatusCode)
		return response, err
	}

	return response, nil
}

func (a *AiServiceImageApiService) FaceBeautyImage(image, modelName string, config map[string]interface{}) (*model.Response, error) {
	request := &model.ImageRequest{
		Action:    "FaceBeautyImage",
		Image:     image,
		ModelName: modelName,
		Configure: config,
	}

	response, httpResponse, err := a.AiModelImage(context.Background(), request)
	if err != nil {
		return response, err
	}

	if httpResponse.StatusCode != 200 {
		err = fmt.Errorf("http status code error:%d", httpResponse.StatusCode)
		return response, err
	}

	return response, nil
}

func (a *AiServiceImageApiService) FaceAttrImage(image, modelName string, config map[string]interface{}) (*model.Response, error) {
	request := &model.ImageRequest{
		Action:    "FaceAttrImage",
		Image:     image,
		ModelName: modelName,
		Configure: config,
	}

	response, httpResponse, err := a.AiModelImage(context.Background(), request)
	if err != nil {
		return response, err
	}

	if httpResponse.StatusCode != 200 {
		err = fmt.Errorf("http status code error:%d", httpResponse.StatusCode)
		return response, err
	}

	return response, nil
}

func (a *AiServiceImageApiService) ImageQA(image, modelName string, config map[string]interface{}) (*model.Response, error) {
	request := &model.ImageRequest{
		Action:    "ImageQA",
		Image:     image,
		ModelName: modelName,
		Configure: config,
	}

	response, httpResponse, err := a.AiModelImage(context.Background(), request)
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
AiServiceImageApiService 图片检测接口
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body 图片检测body

@return Response
*/
func (a *AiServiceImageApiService) AiModelImage(ctx context.Context, body *model.ImageRequest) (*model.Response, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue = model.NewResponse()
	)

	path := "/api/v1/images"
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
