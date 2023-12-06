package test

import (
	"testing"

	sdk "github.com/aliyun/aliyun-pai-aiservice-go-sdk"
)

func TestAICreation(t *testing.T) {
	client := sdk.NewClient(host, appId, token)

	response, err := client.AiCreationApi.AICreation("枯藤老树昏鸦", "", nil)

	if err != nil {
		t.Error(err)
	}

	if response.Code != "OK" {
		t.Errorf("response error:%v", response)
	}
	t.Log(response)
}
