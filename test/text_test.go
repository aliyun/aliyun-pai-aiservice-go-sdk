package test

import (
	"testing"

	sdk "github.com/aliyun/aliyun-pai-aiservice-go-sdk"
)

func TestKeyWordText(t *testing.T) {

	client := sdk.NewClient(host, appId, token)

	response, err := client.TextApi.KeyWordText("日料一般指日本料理，起源于日本列岛，逐渐发展成为具有日本特色的菜肴。主食以米饭、面条为主，因为靠近大海所以副食多为新鲜鱼虾等海产，常配以日本酒。", "", nil)

	if err != nil {
		t.Error(err)
	}
	if response.Code != "OK" {
		t.Errorf("response error:%v", response)
	}
	t.Log(response)
}

func TestNERText(t *testing.T) {

	client := sdk.NewClient(host, appId, token)

	response, err := client.TextApi.NERText("日料一般指日本料理，起源于日本列岛，逐渐发展成为具有日本特色的菜肴。主食以米饭、面条为主，因为靠近大海所以副食多为新鲜鱼虾等海产，常配以日本酒。", "", nil)

	if err != nil {
		t.Error(err)
	}

	if response.Code != "OK" {
		t.Errorf("response error:%v", response)
	}
	t.Log(response)
}

func TestFineGrainedLabelText(t *testing.T) {

	client := sdk.NewClient(host, appId, token)

	response, err := client.TextApi.FineGrainedLabelText("日料一般指日本料理，起源于日本列岛，逐渐发展成为具有日本特色的菜肴。主食以米饭、面条为主，因为靠近大海所以副食多为新鲜鱼虾等海产，常配以日本酒。", "", nil)

	if err != nil {
		t.Error(err)
	}
	if response.Code != "OK" {
		t.Errorf("response error:%v", response)
	}
	t.Log(response)
}

func TestVectorizationText(t *testing.T) {

	client := sdk.NewClient(host, appId, token)

	response, err := client.TextApi.VectorizationText("日料一般指日本料理，起源于日本列岛，逐渐发展成为具有日本特色的菜肴。主食以米饭、面条为主，因为靠近大海所以副食多为新鲜鱼虾等海产，常配以日本酒。", "", nil)

	if err != nil {
		t.Error(err)
	}
	if response.Code != "OK" {
		t.Errorf("response error:%v", response)
	}
	t.Log(response)
}

func TestAntiSpamText(t *testing.T) {

	client := sdk.NewClient(host, appId, token)

	response, err := client.TextApi.AntiSpamText("聘手机小时工,180一天现结有手机就可以", "", nil)

	if err != nil {
		t.Error(err)
	}
	if response.Code != "OK" {
		t.Errorf("response error:%v", response)
	}
	t.Log(response)
}

func TestSentimentAnalysisText(t *testing.T) {

	client := sdk.NewClient(host, appId, token)

	response, err := client.TextApi.SentimentAnalysisText("裤裤面料 做工都很不错 裤型也很漂亮 很满意", "", nil)

	if err != nil {
		t.Error(err)
	}
	if response.Code != "OK" {
		t.Errorf("response error:%v", response)
	}
	t.Log(response)
}

func TestEcomerceReviewSentimentAnalysisText(t *testing.T) {

	client := sdk.NewClient(host, appId, token)

	response, err := client.TextApi.EcomerceReviewSentimentAnalysisText("质量很不错，穿上大小合适，款式很好看，和卖家描述一致", "", nil)

	if err != nil {
		t.Error(err)
	}
	if response.Code != "OK" {
		t.Errorf("response error:%v", response)
	}
	t.Log(response)
}
