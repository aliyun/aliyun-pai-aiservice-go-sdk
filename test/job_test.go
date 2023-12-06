package test

import (
	"encoding/json"
	"fmt"
	"testing"

	sdk "github.com/aliyun/aliyun-pai-aiservice-go-sdk"
)

func TestGetAsyncJob(t *testing.T) {
	client := sdk.NewClient(host, appId, token)

	response, err := client.JobApi.GetAsyncJobWithId(13890)

	if err != nil {
		t.Error(err)
	}
	j, _ := json.Marshal(response)
	fmt.Println(string(j))
}

func TestGetBatchJobWithId(t *testing.T) {
	client := sdk.NewClient(host, appId, token)

	response, err := client.JobApi.GetBatchJobWithId(3)

	if err != nil {
		t.Error(err)
	}
	j, _ := json.Marshal(response)
	fmt.Println(string(j))
}
