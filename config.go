package openai

import (
	"net/http"
	"net/url"
)

const (
	apiURLv1                       = "https://api.openai.com/v1"
	defaultEmptyMessagesLimit uint = 300
)

// ClientConfig is a configuration of a client.
type ClientConfig struct {
	authToken string

	HTTPClient *http.Client

	BaseURL string
	OrgID   string

	EmptyMessagesLimit uint
}

func DefaultConfig(authToken string) ClientConfig {
	// 创建代理请求
	proxyUrl := "http://127.0.0.1:7890"
	proxyUrlParsed, _ := url.Parse(proxyUrl)
	//client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrlParsed)}}
	return ClientConfig{
		HTTPClient: &http.Client{
			Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrlParsed)},
		},
		BaseURL:    apiURLv1,
		OrgID:      "",
		authToken:  authToken,

		EmptyMessagesLimit: defaultEmptyMessagesLimit,
	}
}
