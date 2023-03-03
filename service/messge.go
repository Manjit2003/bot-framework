package service

import (
	"encoding/json"

	"github.com/Manjit2003/bot-framework/types"
)

func (s *WhatsAppService) SendTextMessage(req *types.TextMessageParams) ([]byte, error) {

	if req.PreviewUrl == "" {
		req.PreviewUrl = "false"
	}

	return s.makeRequest(&types.ReqParams{
		Params: map[string]string{
			"msg":        req.Text,
			"mobile":     req.To,
			"previewUrl": req.PreviewUrl,
			"header":     req.Header,
			"footer":     req.Footer,
			"msgType":    "text",
		},
	})
}

func (s *WhatsAppService) SendListMessage(req *types.ListMessageParams) ([]byte, error) {

	if req.PreviewUrl == "" {
		req.PreviewUrl = "false"
	}

	b, err := json.Marshal(req.ListData)
	if err != nil {
		return nil, err
	}

	return s.makeRequest(&types.ReqParams{
		Params: map[string]string{
			"msg":            req.Text,
			"mobile":         req.To,
			"previewUrl":     req.PreviewUrl,
			"header":         req.Header,
			"footer":         req.Footer,
			"msgType":        "list",
			"buttonsPayload": string(b),
		},
	})
}

func (s *WhatsAppService) SendButtonMessage(req *types.ButtonData) ([]byte, error) {

	if req.PreviewUrl == "" {
		req.PreviewUrl = "false"
	}

	payload := struct {
		Buttons []types.Button `json:"buttons"`
	}{
		Buttons: req.Buttons,
	}

	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	return s.makeRequest(&types.ReqParams{
		Params: map[string]string{
			"msg":            req.Text,
			"mobile":         req.To,
			"previewUrl":     req.PreviewUrl,
			"header":         req.Header,
			"footer":         req.Footer,
			"msgType":        "reply",
			"buttonsPayload": string(b),
		},
	})
}

func (s *WhatsAppService) SendDocumentMessage(req *types.DocumentMessageParams) ([]byte, error) {

	if req.PreviewUrl == "" {
		req.PreviewUrl = "false"
	}

	return s.makeRequest(&types.ReqParams{
		Params: map[string]string{
			"msg":          req.Text,
			"mobile":       req.To,
			"previewUrl":   req.PreviewUrl,
			"header":       req.Header,
			"footer":       req.Footer,
			"msgType":      "media",
			"mediaType":    "document",
			"caption":      req.Text,
			"mediaUrl":     req.DocumentData.DocumentUrl,
			"documentName": req.DocumentData.DocumentName,
		},
	})
}

func (s *WhatsAppService) SendMediaMessage(req *types.MediaMessageParams) ([]byte, error) {

	if req.PreviewUrl == "" {
		req.PreviewUrl = "false"
	}

	return s.makeRequest(&types.ReqParams{
		Params: map[string]string{
			"msg":        req.Text,
			"mobile":     req.To,
			"previewUrl": req.PreviewUrl,
			"header":     req.Header,
			"footer":     req.Footer,
			"msgType":    "media",
			"mediaType":  string(req.MediaData.MediaType),
			"caption":    req.Text,
			"mediaUrl":   req.MediaData.MediaUrl,
		},
	})

}
