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

const (
	VIDEO_PATH       = "/api/v1/videos"
	BATCH_VIDEO_PATH = "/api/v1/batch_videos"
)

type AiServiceVideoApiService service

func (a *AiServiceVideoApiService) MultiModalClassify(videoUrl string, images []string, text string, modelName string, config map[string]interface{}) (*model.Response, error) {
	request := &model.VideoRequest{
		Action:    "MultiModalClassify",
		VideoUrl:  videoUrl,
		Images:    images,
		Text:      text,
		ModelName: modelName,
		Configure: config,
	}

	response, httpResponse, err := a.AiModelVideo(context.Background(), request)
	if err != nil {
		return response, err
	}

	if httpResponse.StatusCode != 200 {
		err = fmt.Errorf("http status code error:%d", httpResponse.StatusCode)
		return response, err
	}

	return response, nil
}

func (a *AiServiceVideoApiService) ClassifyLabelVideo(videoUrl, videoTitle, modelName string, config map[string]interface{}) (*model.Response, error) {
	request := &model.VideoRequest{
		Action:     "ClassifyLabelVideo",
		VideoUrl:   videoUrl,
		VideoTitle: videoTitle,
		ModelName:  modelName,
		Configure:  config,
	}

	response, httpResponse, err := a.AiModelVideo(context.Background(), request)
	if err != nil {
		return response, err
	}

	if httpResponse.StatusCode != 200 {
		err = fmt.Errorf("http status code error:%d", httpResponse.StatusCode)
		return response, err
	}

	return response, nil
}

func (a *AiServiceVideoApiService) BatchClassifyLabelVideo(videos []model.VideoMeta, modelName string, config map[string]interface{}) (*model.Response, error) {
	request := &model.BatchVideoRequest{
		Action:    "ClassifyLabelVideo",
		Videos:    videos,
		ModelName: modelName,
		Configure: config,
	}

	response, httpResponse, err := a.AiModelBatchVideo(context.Background(), request)
	if err != nil {
		return response, err
	}

	if httpResponse.StatusCode != 200 {
		err = fmt.Errorf("http status code error:%d", httpResponse.StatusCode)
		return response, err
	}

	return response, nil
}

func (a *AiServiceVideoApiService) VideoQA(videoUrl, videoTitle, modelName string, config map[string]interface{}) (*model.Response, error) {
	request := &model.VideoRequest{
		Action:     "VideoQA",
		VideoUrl:   videoUrl,
		VideoTitle: videoTitle,
		ModelName:  modelName,
		Configure:  config,
	}

	response, httpResponse, err := a.AiModelVideo(context.Background(), request)
	if err != nil {
		return response, err
	}

	if httpResponse.StatusCode != 200 {
		err = fmt.Errorf("http status code error:%d", httpResponse.StatusCode)
		return response, err
	}

	return response, nil
}

func (a *AiServiceVideoApiService) BatchVideoQA(videos []model.VideoMeta, modelName string, config map[string]interface{}) (*model.Response, error) {
	request := &model.BatchVideoRequest{
		Action:    "VideoQA",
		Videos:    videos,
		ModelName: modelName,
		Configure: config,
	}

	response, httpResponse, err := a.AiModelBatchVideo(context.Background(), request)
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
AiServiceVideoApiService 视频检测接口
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body 视频检测body

@return Response
*/
func (a *AiServiceVideoApiService) AiModelVideo(ctx context.Context, body *model.VideoRequest) (*model.Response, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue = model.NewResponse()
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + VIDEO_PATH

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

	sign := a.client.createSign(a.cfg, VIDEO_PATH, localVarHttpMethod, localVarPostBody, localVarQueryParams)
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

func (a *AiServiceVideoApiService) AiModelBatchVideo(ctx context.Context, body *model.BatchVideoRequest) (*model.Response, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue = model.NewResponse()
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + BATCH_VIDEO_PATH

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
	sign := a.client.createSign(a.cfg, BATCH_VIDEO_PATH, localVarHttpMethod, localVarPostBody, localVarQueryParams)
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
