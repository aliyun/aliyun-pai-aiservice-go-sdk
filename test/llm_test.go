package test

import (
	"testing"

	sdk "github.com/aliyun/aliyun-pai-aiservice-go-sdk"
)

func TestLLMEvalJudger(t *testing.T) {
	client := sdk.NewClient(host, appId, token)

	response, err := client.LLMApi.LlmEvalJudger(sdk.ChatCompletionRequest{
		Model: "Themis-1.0",
		Messages: []sdk.ChatCompletionMessage{
			{
				Role:    "user",
				Content: "请介绍下你自己",
			},
			{
				Role:    "assistant",
				Content: "我是一个人工智能助手，由 OpenAI 设计并开发，旨在帮助用户解答问题、提供信息、学习语言和执行各种语言处理任务。作为 AI，我可以理解和生成自然语言文本，但与真实的人类助手或生物智能不同，我不拥有个人经历、情感或主观意识。我的知识是基于数据和预先训练的模型，并且我遵循特定的编程协议和道德指南来确保提供有用、准确和公正的信息。\\n我的功能包括但不限于：\\n回答一般知识问题\\n提供解释和定义\\n协助学习和教育\\n翻译语言\\n辅助编程和技术问题\\n提供建议和推荐\\n我在处理请求时不能进行个人化决策或提供个人建议，也不储存任何个人信息。我的目标是成为一个有益的资源，让用户更容易获取信息和解答。",
			},
		},
	}, "")

	if err != nil {
		t.Error(err)
	}
	if response.Code != "OK" {
		t.Errorf("response error:%v", response)
	}
	t.Log(response)
}
