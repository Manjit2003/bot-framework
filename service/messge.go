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

	payload := struct {
		Buttons []types.Button `json:"buttons"`
	}{
		Buttons: req.Buttons,
	}

	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	reqParams := &types.ReqParams{
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
	}

	if len(req.Buttons) > 0 {
		reqParams.Params["buttonsPayload"] = string(b)
		reqParams.Params["msgType"] = "reply"
	}

	return s.makeRequest(reqParams)

}

func (s *WhatsAppService) SendContactMessage(req *types.ContactMessageParams) ([]byte, error) {

	if req.PreviewUrl == "" {
		req.PreviewUrl = "false"
	}

	type phoneData struct {
		Phone string `json:"phone"`
		Type  string `json:"type"`
		WaId  string `json:"wa_id"`
	}

	var payload struct {
		Name struct {
			FirstName     string `json:"first_name"`
			LastName      string `json:"last_name"`
			FormattedName string `json:"formatted_name"`
		} `json:"name"`
		Phones    []phoneData `json:"phones"`
		Addresses []struct{}  `json:"addresses"`
	}

	payload.Name.FormattedName = req.ContactData.ContactName
	payload.Name.FirstName = req.ContactData.ContactName
	payload.Name.LastName = req.ContactData.ContactName
	payload.Phones = append(payload.Phones, phoneData{
		Phone: req.ContactData.ContactPhone,
		Type:  "WORK",
		WaId:  req.ContactData.ContactPhone,
	})

	b, err := json.Marshal(payload)

	if err != nil {
		return nil, err
	}

	return s.makeRequest(&types.ReqParams{
		Params: map[string]string{
			"msg":        string(b),
			"mobile":     req.To,
			"previewUrl": req.PreviewUrl,
			"header":     req.Header,
			"footer":     req.Footer,
			"msgType":    "contact",
		},
	})
}
