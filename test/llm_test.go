package test

import (
	"testing"

	sdk "github.com/aliyun/aliyun-pai-aiservice-go-sdk"
	"github.com/aliyun/aliyun-pai-aiservice-go-sdk/model"
)

func TestLLMEvalJudger(t *testing.T) {
	client := sdk.NewClient(host, appId, token)
	message := []model.LlmEvalJudgerRequestMessages{{
		Role:    "user",
		Content: "[\n{\n\"mode\": \"single\",\n\"type\": \"json\",\n\"json\": {\n\"question\": \"请介绍下你自己\",\n\"answer\": \"我是一个人工智能助手，由 OpenAI 设计并开发，旨在帮助用户解答问题、提供信息、学习语言和执行各种语言处理任务。\"\n}\n}\n]",
	}}
	response, err := client.LLMApi.Create("", "Themis-1.0", message)

	if err != nil {
		t.Error(err)
	}
	if response.Code != "OK" {
		t.Errorf("response error:%v", response)
	}
	t.Log(response)
}
