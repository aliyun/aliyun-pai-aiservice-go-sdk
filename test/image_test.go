package test

import (
	"testing"

	sdk "github.com/aliyun/aliyun-pai-aiservice-go-sdk"
)

func TestOCRImage(t *testing.T) {
	client := sdk.NewClient(host, appId, token)

	response, err := client.ImageApi.OCRImage("https://pai-rec-public.oss-cn-shanghai.aliyuncs.com/ai-service/demo/lADPJvDqUu318wPNAWDNAjw_572_352.jpg", "", nil)

	if err != nil {
		t.Error(err)
	}

	if response.Code != "OK" {
		t.Errorf("response error:%v", response)
	}
	t.Log(response)
}

func TestClassifyImage(t *testing.T) {
	client := sdk.NewClient(host, appId, token)

	response, err := client.ImageApi.ClassifyImage("https://pai-rec-public.oss-cn-shanghai.aliyuncs.com/ai-service/demo/lADPJvDqUu318wPNAWDNAjw_572_352.jpg", "", nil)

	if err != nil {
		t.Error(err)
	}
	if response.Code != "OK" {
		t.Errorf("response error:%v", response)
	}
	t.Log(response)
}

func TestMultiLabelImage(t *testing.T) {
	client := sdk.NewClient(host, appId, token)

	response, err := client.ImageApi.MultiLabelImage("http://47.95.234.175/assets/229ca549d18e28b025aa28a933eb99d5.jpg", "", nil)

	if err != nil {
		t.Error(err)
	}
	if response.Code != "OK" {
		t.Errorf("response error:%v", response)
	}
	t.Log(response)
}

func TestMultiLabelImageV2(t *testing.T) {
	client := sdk.NewClient(host, appId, token)

	response, err := client.ImageApi.MultiLabelImageV2("http://47.95.234.175/assets/229ca549d18e28b025aa28a933eb99d5.jpg", "", nil)

	if err != nil {
		t.Error(err)
	}
	if response.Code != "OK" {
		t.Errorf("response error:%v", response)
	}
	t.Log(response)
}

func TestFaceRecognitionImage(t *testing.T) {
	client := sdk.NewClient(host, appId, token)

	response, err := client.ImageApi.FaceRecognitionImage("https://pai-rec-public.oss-cn-shanghai.aliyuncs.com/ai-service/demo/lADPJvDqUu318wPNAWDNAjw_572_352.jpg", "", nil)

	if err != nil {
		t.Error(err)
	}

	if response.Code != "OK" {
		t.Errorf("response error:%v", response)
	}
	t.Log(response)
}

func TestDetectClassifyImage(t *testing.T) {
	client := sdk.NewClient(host, appId, token)

	response, err := client.ImageApi.DetectClassifyImage("https://pai-rec-public.oss-cn-shanghai.aliyuncs.com/ai-service/demo/lADPJvDqUu318wPNAWDNAjw_572_352.jpg", "", nil)

	if err != nil {
		t.Error(err)
	}

	if response.Code != "OK" {
		t.Errorf("response error:%v", response)
	}
	t.Log(response)
}

func TestScanImage(t *testing.T) {
	client := sdk.NewClient(host, appId, token)

	response, err := client.ImageApi.ScanImage("https://pai-rec-public.oss-cn-shanghai.aliyuncs.com/ai-service/demo/lADPJvDqUu318wPNAWDNAjw_572_352.jpg", "", nil)

	if err != nil {
		t.Error(err)
	}

	if response.Code != "OK" {
		t.Errorf("response error:%v", response)
	}
	t.Log(response)
}

func TestProductDetectImage(t *testing.T) {
	client := sdk.NewClient(host, appId, token)

	response, err := client.ImageApi.ProductDetectImage("https://pai-rec-public.oss-cn-shanghai.aliyuncs.com/ai-service/demo/lADPJvDqUu318wPNAWDNAjw_572_352.jpg", "", nil)

	if err != nil {
		t.Error(err)
	}

	if response.Code != "OK" {
		t.Errorf("response error:%v", response)
	}
	t.Log(response)
}

func TestFaceDetectImage(t *testing.T) {
	client := sdk.NewClient(host, appId, token)

	response, err := client.ImageApi.FaceDetectImage("https://pai-rec-public.oss-cn-shanghai.aliyuncs.com/ai-service/demo/face_chenglong.png", "", nil)

	if err != nil {
		t.Error(err)
	}

	if response.Code != "OK" {
		t.Errorf("response error:%v", response)
	}
	t.Log(response)
}

func TestFaceBeautyImage(t *testing.T) {
	client := sdk.NewClient(host, appId, token)

	response, err := client.ImageApi.FaceBeautyImage("https://pai-rec-public.oss-cn-shanghai.aliyuncs.com/ai-service/demo/face_chenglong.png", "", nil)

	if err != nil {
		t.Error(err)
	}

	if response.Code != "OK" {
		t.Errorf("response error:%v", response)
	}
	t.Log(response)
}

func TestFaceAttrImage(t *testing.T) {
	client := sdk.NewClient(host, appId, token)

	response, err := client.ImageApi.FaceAttrImage("https://pai-rec-public.oss-cn-shanghai.aliyuncs.com/ai-service/demo/face_chenglong.png", "", nil)

	if err != nil {
		t.Error(err)
	}

	if response.Code != "OK" {
		t.Errorf("response error:%v", response)
	}
	t.Log(response)
}

func TestImageTextTag(t *testing.T) {
	client := sdk.NewClient(host, appId, token)

	images := []string{"https://pai-rec-public.oss-cn-shanghai.aliyuncs.com/ai-service/demo/face_chenglong.png"}

	response, err := client.ImageTextApi.ImageTextTag(images, "肖像", "", nil)

	if err != nil {
		t.Error(err)
	}

	if response.Code != "OK" {
		t.Errorf("response error:%v", response)
	}
	t.Log(response)
}

func TestImageQA(t *testing.T) {
	client := sdk.NewClient(host, appId, token)

	response, err := client.ImageApi.ImageQA("https://pai-rec-public.oss-cn-shanghai.aliyuncs.com/ai-service/demo/face_chenglong.png", "", nil)

	if err != nil {
		t.Error(err)
	}

	if response.Code != "OK" {
		t.Errorf("response error:%v", response)
	}
	t.Log(response)
}
