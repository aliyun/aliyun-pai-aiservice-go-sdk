package test

import (
	"testing"

	sdk "github.com/aliyun/aliyun-pai-aiservice-go-sdk"
)

func TestLLMEvalJudger(t *testing.T) {
	client := sdk.NewClient(host, appId, token)
	content := `[{"role":"user","content":[{"mode":"single","type":"json","json":{"question":"评价一下2021年的暴雪游戏公司","answer":"2021年对暴雪游戏公司来说是一个充满挑战和变化的一年。公司在这一年面临了一些负面的新闻和争议，比如员工抗议、性别歧视指控等， 这些事件对公司声誉造成了一定影响。然而，暴雪游戏公司也在2021年发布了一些备受期待的游戏作品，比如《暗黑破坏神2：复仇》和《炉石传说》的新资料片等，这些作品受到了玩家们的欢迎和好评。总的来说，2021年对暴雪游戏公司来说是一个充满挑战和机遇并存的一年，希望他们能够在未来继续努力，重塑公司形象，为玩家们带来更多优秀的游戏作品。"}}]}]
	`

	response, err := client.LLMApi.Create("", "Themis-1.0", content)

	if err != nil {
		t.Error(err)
	}
	if response.Code != "OK" {
		t.Errorf("response error:%v", response)
	}
	t.Log(response)
}

func TestLLMEvalJudger2(t *testing.T) {
	client := sdk.NewClient(host, appId, token)

	content := `[{"role":"user","content":[{"mode":"pairwise","type":"json","json":{"question":"请介绍下你自己","answer1":"我是一个人工智能助手，由 OpenAI 设计并开发，旨在帮助用户解答问题、提供信息、学习语言和执行各种语言处理任务。","answer2":"我是一个人工智能助手，设计用来回答问题、提供信息、助力解决问题，并在不同的领域与场景中辅助用户。"}}]}]
	`
	response, err := client.LLMApi.Create("", "Themis-1.0", content)

	if err != nil {
		t.Error(err)
	}
	if response.Code != "OK" {
		t.Errorf("response error:%v", response)
	}
	t.Log(response)

	t.Log(response.Data.Model, response.Data.Id)
	for _, choice := range response.Data.Choices {
		t.Log(choice.Index, choice.Text)
	}
}
