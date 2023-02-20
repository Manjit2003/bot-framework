package service

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/Manjit2003/bot-framework/types"
)

const DefaultHost types.Host = "https://theultimate.io"

type WhatsAppService struct {
	Host       types.Host
	Token      string
	UserId     string
	Password   string
	WaBaNumber string
	httpClient *http.Client
}

func NewWhatsAppService(cfg *types.ServiceConfig) *WhatsAppService {
	return &WhatsAppService{
		Host:       cfg.Host,
		Token:      cfg.Token,
		UserId:     cfg.UserId,
		Password:   cfg.Password,
		WaBaNumber: cfg.WaBaNumber,

		httpClient: &http.Client{},
	}
}

func (s *WhatsAppService) makeRequest(reqData *types.ReqParams) ([]byte, error) {

	url := fmt.Sprintf("%s/WAApi/send", s.Host)
	method := "POST"

	fields := map[string]string{
		"userid":     s.UserId,
		"password":   s.Password,
		"wabaNumber": s.WaBaNumber,
		"output":     "json",
		"sendMethod": "quick",
	}

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	for key, val := range reqData.Params {
		if val == "" {
			continue
		}
		fields[key] = val
	}

	for key, val := range fields {

		_ = writer.WriteField(key, val)

	}

	err := writer.Close()
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return nil, err
	}
	req.Header.Add("apikey", s.Token)

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
