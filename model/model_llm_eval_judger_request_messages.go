package model

import (
	"encoding/json"
	"errors"
)

var (
	ErrContentFieldsMisused = errors.New("can't use both Content and MultiContent properties simultaneously")
	ErrAnswerFieldsMissing  = errors.New("answer field missing")
	ErrAnswer1FieldsMissing = errors.New("answer1 field missing")
	ErrAnswer2FieldsMissing = errors.New("answer2 field missing")
)

type ChatMessagePartType string

const (
	ChatMessagePartTypeJson ChatMessagePartType = "json"
)

type ChatMessagePartMode string

const (
	ChatMessagePartModePair   ChatMessagePartMode = "pairwise"
	ChatMessagePartModeSingle ChatMessagePartMode = "single"
)

type ChatMessagePartJson struct {
	Question  string `json:"question"`
	Answer    string `json:"answer,omitempty"`
	Answer1   string `json:"answer1,omitempty"`
	Answer2   string `json:"answer2,omitempty"`
	Scene     string `json:"scene,omitempty"`
	SceneDesc string `json:"scene_desc,omitempty"`
	Metric    string `json:"metric,omitempty"`
	MaxScore  string `json:"max_score,omitempty"`
	ScoreDesc string `json:"score_desc,omitempty"`
	RefAnswer string `json:"ref_answer,omitempty"`
	Steps     string `json:"steps,omitempty"`
}

type ChatMessagePart struct {
	Type ChatMessagePartType `json:"type,omitempty"`
	Mode ChatMessagePartMode `json:"mode"`
	Json ChatMessagePartJson `json:"json"`
}

type LlmEvalJudgerRequestMessages struct {
	Role         string `json:"role,omitempty"`
	Content      string `json:"content,omitempty"`
	MultiContent []ChatMessagePart
}

func (m LlmEvalJudgerRequestMessages) MarshalJSON() ([]byte, error) {
	if m.Content != "" && m.MultiContent != nil {
		return nil, ErrContentFieldsMisused
	}

	if len(m.MultiContent) > 0 {
		for i := range m.MultiContent {
			if m.MultiContent[i].Type == "" {
				m.MultiContent[i].Type = ChatMessagePartTypeJson
			}

			if m.MultiContent[i].Json.Answer == "" {
				if m.MultiContent[i].Json.Answer1 == "" && m.MultiContent[i].Json.Answer2 == "" {
					return nil, ErrAnswerFieldsMissing
				} else if m.MultiContent[i].Json.Answer1 == "" {
					return nil, ErrAnswer1FieldsMissing
				} else if m.MultiContent[i].Json.Answer2 == "" {
					return nil, ErrAnswer2FieldsMissing
				}
			}

			if m.MultiContent[i].Mode == "" {
				if m.MultiContent[i].Json.Answer != "" {
					m.MultiContent[i].Mode = ChatMessagePartModeSingle
				} else if m.MultiContent[i].Json.Answer1 != "" && m.MultiContent[i].Json.Answer2 != "" {
					m.MultiContent[i].Mode = ChatMessagePartModePair
				}
			}
		}

		msg := struct {
			Role         string            `json:"role"`
			Content      string            `json:"-"`
			MultiContent []ChatMessagePart `json:"content,omitempty"`
		}(m)
		return json.Marshal(msg)
	}
	msg := struct {
		Role         string            `json:"role"`
		Content      string            `json:"content"`
		MultiContent []ChatMessagePart `json:"-"`
	}(m)
	return json.Marshal(msg)
}

func (m *LlmEvalJudgerRequestMessages) UnmarshalJSON(bs []byte) error {
	msg := struct {
		Role         string `json:"role"`
		Content      string `json:"content"`
		MultiContent []ChatMessagePart
	}{}
	if err := json.Unmarshal(bs, &msg); err == nil {
		*m = LlmEvalJudgerRequestMessages(msg)
		return nil
	}
	multiMsg := struct {
		Role         string `json:"role"`
		Content      string
		MultiContent []ChatMessagePart `json:"content"`
	}{}
	if err := json.Unmarshal(bs, &multiMsg); err != nil {
		return err
	}
	*m = LlmEvalJudgerRequestMessages(multiMsg)
	return nil
}
