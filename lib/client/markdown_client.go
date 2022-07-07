package client

import (
	"time"
	"zph/config"
	"zph/lib/client/model"
)

type MarkdownClient interface {
	GetMarkdownHtml(markdown string) (*model.MarkdownHtml, error)
}

func NewMarkdownClient() MarkdownClient {
	return &MarkdownHttpClient{
		httpClient: NewHttpClient(10, 10*time.Second, 10*time.Second),
		host:       config.GetInternalServiceHost() + "/get_html",
	}
}

type MarkdownHttpClient struct {
	httpClient *HttpClient
	host       string
}

func (c *MarkdownHttpClient) GetMarkdownHtml(markdown string) (*model.MarkdownHtml, error) {
	retryTimes := 3
	var resp model.MarkdownHtml
	var body interface{} = map[string]string{
		"markdown": markdown,
	}
	header := map[string]string{
		"content-type": "application/json",
	}
	err := c.httpClient.PostResponseJson(c.host, &body, header, retryTimes, &resp)
	if err != nil {
		log.Errorf("GetMarkdownHtml failed, err is: %+v", err)
		return nil, err
	}
	return &resp, err
}
