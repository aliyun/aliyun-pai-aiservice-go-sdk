package test

import (
	"fmt"
	"log"
	"testing"
	"time"

	sdk "github.com/aliyun/aliyun-pai-aiservice-go-sdk"
)

func TestAIGCImagesCheck(t *testing.T) {
	client := sdk.NewClient(host, appId, token)

	images := []string{"https://pai-aigc-photog-bj.oss-cn-beijing.aliyuncs.com/photog/user_images/foto/train/1.jpg", "https://pai-aigc-photog-bj.oss-cn-beijing.aliyuncs.com/photog/user_images/foto/train/0.jpg"}
	response, err := client.AIGCApi.AigcImageCheck(images, "", nil)

	if err != nil {
		t.Error(err)
	}

	if response.Code != "OK" {
		t.Errorf("response error:%v", response)
	}

	for _, result := range response.Data.CheckResults {
		fmt.Printf("code=%d\tfrontal=%v\turl=%s\n", result.Code, result.Frontal, result.Url)
	}
	fmt.Println(response)
}

func TestAIGCImagesTrain(t *testing.T) {
	client := sdk.NewClient(host, appId, token)

	images := []string{"https://pai-aigc-photog-bj.oss-cn-beijing.aliyuncs.com/photog/user_images/foto/train/1.jpg", "https://pai-aigc-photog-bj.oss-cn-beijing.aliyuncs.com/photog/user_images/foto/train/0.jpg"}
	response, err := client.AIGCApi.AigcImagesTrain(images, "", nil)

	if err != nil {
		t.Error(err)
	}

	if response.Code != "OK" {
		t.Errorf("response error:%v", response)
	}
	t.Log(response)
}

func TestAIGCImagesCreate(t *testing.T) {
	client := sdk.NewClient(host, appId, token)

	response, err := client.AIGCApi.AigcImagesCreate("77be0dff-91a0-4fa6-b6a6-5cf94f5194fc", "https://pai-aigc-photog-bj.oss-cn-beijing.aliyuncs.com/photog/user_weights/foto10/validation/global_step_Blue_1_100_0.jpg", "", nil)

	if err != nil {
		t.Error(err)
	}
	if response.Code != "OK" {
		t.Errorf("response error:%v", response)
	}
	t.Log(response)
	fmt.Println(response.Data.Image)
}

func TestAIGCImagesCreateByMultiModelIds(t *testing.T) {
	client := sdk.NewClient(host, appId, token)
	modelIds := []string{"115aaf09-47a0-4aef-98cf-3582ceffb675", "115aaf09-47a0-4aef-98cf-3582ceffb675"}
	response, err := client.AIGCApi.AigcImagesCreateByMultiModelIds(modelIds, "https://pai-aigc-photog-bj.oss-cn-beijing.aliyuncs.com/photog/user_weights/foto10/validation/global_step_Blue_1_100_0.jpg", "", nil)

	if err != nil {
		t.Error(err)
	}
	if response.Code != "OK" {
		t.Errorf("response error:%v", response)
	}
	t.Log(response)
}

func TestAIGCTrainAndCreate(t *testing.T) {
	client := sdk.NewClient(host, appId, token)

	images := []string{"https://pai-aigc-photog-bj.oss-cn-beijing.aliyuncs.com/photog/user_images/foto/train/1.jpg", "https://pai-aigc-photog-bj.oss-cn-beijing.aliyuncs.com/photog/user_images/foto/train/0.jpg"}
	response, err := client.AIGCApi.AigcImagesTrain(images, "", nil)
	if err != nil {
		log.Fatal(err)
	}

	// 任务ID
	jobId := response.Data.JobId
	modelId := response.Data.ModelId

	for {
		jobResponse, err := client.JobApi.GetAsyncJobWithId(jobId)
		if err != nil {
			log.Fatal("get job fail", err)
		}

		if jobResponse.Data.Job.State == sdk.JOB_STATE_WAIT {
			fmt.Println("job running")
		} else if jobResponse.Data.Job.State == sdk.JOB_STATE_SUCCESS { // job success
			fmt.Println(jobResponse)
			fmt.Println("job success")
			break
		} else if jobResponse.Data.Job.State == sdk.JOB_STATE_FAILED {
			log.Fatal("job fail", err)
			return
		}

		time.Sleep(30 * time.Second)
	}

	createResponse, err := client.AIGCApi.AigcImagesCreate(modelId, "https://pai-aigc-photog-bj.oss-cn-beijing.aliyuncs.com/photog/user_weights/foto10/validation/global_step_Blue_1_100_0.jpg", "", nil)
	if err != nil {
		log.Fatal(err)
	}

	// 生成的图片 base64 编码
	fmt.Println(createResponse.Data.Image)

}
