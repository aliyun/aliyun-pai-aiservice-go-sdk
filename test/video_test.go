package test

import (
	"encoding/json"
	"fmt"
	"testing"

	sdk "github.com/aliyun/aliyun-pai-aiservice-go-sdk"
	"github.com/aliyun/aliyun-pai-aiservice-go-sdk/model"
)

func TestClassifyLabelVideo(t *testing.T) {
	client := sdk.NewClient(host, appId, token)

	response, err := client.VideoApi.ClassifyLabelVideo("https://ocr-alt.oss-cn-zhangjiakou.aliyuncs.com/uc_video/2881191243992322456.mp4", "2881191243992322456", "", nil)

	if err != nil {
		t.Error(err)
	}
	j, _ := json.Marshal(response)
	if response.Code != "OK" {
		t.Fatal(response)
	}
	t.Log(string(j))
	jobId := response.Data["job_id"].(float64)

	jobResponse, err := client.JobApi.GetAsyncJobWithId(int32(jobId))

	if err != nil {
		t.Error(err)
	}
	fmt.Println(jobResponse)
}

func TestBatchClassifyLabelVideo(t *testing.T) {
	client := sdk.NewClient(host, appId, token)

	videos := []model.VideoMeta{
		{
			Url:   "https://ocr-alt.oss-cn-zhangjiakou.aliyuncs.com/uc_video/2881191243992322456.mp4",
			Title: "test video",
		},
		{
			Url: "https://ocr-alt.oss-cn-zhangjiakou.aliyuncs.com/uc_video/2881191243992322456.mp4",
		},
	}
	response, err := client.VideoApi.BatchClassifyLabelVideo(videos, "", nil)

	if err != nil {
		t.Error(err)
	}
	j, _ := json.Marshal(response)
	if response.Code != "OK" {
		t.Fatal(response)
	}
	t.Log(string(j))
	batchId := response.Data["batch_id"].(float64)

	response, err = client.JobApi.GetBatchJobWithId(int32(batchId))

	if err != nil {
		t.Error(err)
	}
	j, _ = json.Marshal(response)
	t.Log(string(j))
}

func TestVideoQA(t *testing.T) {
	client := sdk.NewClient(host, appId, token)

	response, err := client.VideoApi.VideoQA("https://ocr-alt.oss-cn-zhangjiakou.aliyuncs.com/uc_video/2881191243992322456.mp4", "2881191243992322456", "", nil)

	if err != nil {
		t.Error(err)
	}
	j, _ := json.Marshal(response)
	if response.Code != "OK" {
		t.Fatal(response)
	}
	t.Log(string(j))
	jobId := response.Data["job_id"].(float64)

	jobResponse, err := client.JobApi.GetAsyncJobWithId(int32(jobId))

	if err != nil {
		t.Error(err)
	}
	j, _ = json.Marshal(jobResponse)
	t.Log(string(j))
}

func TestBatchVideoQA(t *testing.T) {
	client := sdk.NewClient(host, appId, token)

	videos := []model.VideoMeta{
		{
			Url:   "https://ocr-alt.oss-cn-zhangjiakou.aliyuncs.com/uc_video/2881191243992322456.mp4",
			Title: "test video",
		},
		{
			Url: "https://ocr-alt.oss-cn-zhangjiakou.aliyuncs.com/uc_video/2881191243992322456.mp4",
		},
	}
	response, err := client.VideoApi.BatchVideoQA(videos, "", nil)

	if err != nil {
		t.Error(err)
	}
	j, _ := json.Marshal(response)
	if response.Code != "OK" {
		t.Fatal(response)
	}
	t.Log(string(j))
	batchId := response.Data["batch_id"].(float64)

	response, err = client.JobApi.GetBatchJobWithId(int32(batchId))

	if err != nil {
		t.Error(err)
	}
	j, _ = json.Marshal(response)
	t.Log(string(j))
}

func TestMultiModalClassify(t *testing.T) {
	client := sdk.NewClient(host, appId, token)

	video := "https://v1.keepcdn.com/tweet-entry-video/5f24184e952174187bcc05b9/2023/11/22/7f78db08-cd06-478d-94c7-90d5476a1756/1700620615040655d69474d2e24000186c6e3.mp4"
	images := []string{"https://static1.keepcdn.com/tweet-entry/58022a3d4933036dfecd1d11/2023/11/22/0cfef2c2-dd10-410c-b698-819e6a4022f3/1f270cf122c612c557e50af21d24c963_1080x1440.jpg"}
	text := "标题:晨间跑步 5.04 km 。内容:坚持坚持再坚持"

	response, err := client.VideoApi.MultiModalClassify(video, images, text, "keep_ugc_level", nil)

	if err != nil {
		t.Error(err)
	}
	j, _ := json.Marshal(response)
	if response.Code != "OK" {
		t.Fatal(response)
	}
	t.Log(string(j))
	jobId := response.Data["job_id"].(float64)
	t.Log(jobId)
}
