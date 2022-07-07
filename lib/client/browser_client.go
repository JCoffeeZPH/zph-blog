package client

import (
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"zph/config"
	"zph/lib/client/model"
	"zph/lib/common"
)

type BrowserClient interface {
	GetBrowserDetails(userAgent string) (*model.BrowserDetail, error)
}

func NewBrowserClient() BrowserClient {
	return &BrowserHttpClient{
		httpClient: NewHttpClient(10, 10*time.Second, 10*time.Second),
		host:       config.GetBrowserDetailUrl(),
	}
}

type BrowserHttpClient struct {
	httpClient *HttpClient
	host       string
}

func (c *BrowserHttpClient) GetBrowserDetails(userAgent string) (*model.BrowserDetail, error) {
	reqUrl := c.host
	var resp model.BrowserDetail
	if strings.Contains(strings.ToLower(userAgent), "postman") {
		return nil, nil
	}
	params := map[string]string{
		"ua": userAgent,
	}
	err := c.getResponseJson(reqUrl, nil, params, &resp)
	if err != nil {
		log.Errorf("GetBrowserDetails failed, user-agent: %s, err is: %+v", userAgent, err)
		return nil, err
	}
	return &resp, nil
}

func (c *BrowserHttpClient) getResponseJson(requestUrl string, headers map[string]string, parameters map[string]string, responseBody *model.BrowserDetail) error {
	retryTimes := 3
	responseHandler := func(response *http.Response) error {
		responseBodyBytes, _ := ioutil.ReadAll(response.Body)
		if response.StatusCode != http.StatusOK {
			return &common.HttpError{
				ErrorCode:    response.StatusCode,
				ResponseBody: string(responseBodyBytes),
				URL:          requestUrl,
			}
		}
		err := jsoniter.Unmarshal(responseBodyBytes, responseBody)
		if err != nil {
			log.Errorf("Unmarshal response error, response: %v, %v", string(responseBodyBytes), err)
		}
		return err
	}
	err := c.httpClient.Get(requestUrl, parameters, headers, retryTimes, responseHandler)
	if err != nil {
		log.Errorf("getResponseJson url: %s, params: %+v, err: %+v", requestUrl, parameters, err.Error())
		return err
	}
	return nil
}
