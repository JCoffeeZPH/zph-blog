package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"zph/config"
	"zph/lib/client/model"
	"zph/lib/common"
)

type QQHttpClient struct {
	httpClient *HttpClient
	host       string
}

func NewQQHttpClient() *QQHttpClient {
	return &QQHttpClient{
		httpClient: NewHttpClient(10, 10*time.Second, 10*time.Second),
		host:       config.GetInternalServiceHost(),
	}
}

func (c *QQHttpClient) GetQQImageCDNUrl(qq string) (*model.QQDetails, error) {
	resMap := make(map[string]string)
	err := c.postResponseJson(qq, resMap)
	if err != nil {
		return nil, err
	}

	return &model.QQDetails{
		QQAvatarCDNUrl: resMap[qq],
		Nickname:       resMap["nickname"],
	}, nil
}

func (c *QQHttpClient) postResponseJson(qq string, responseBody map[string]string) error {
	requestUrl := c.host + "/image/upload"
	retryTimes := 3
	responseHandler := func(response *http.Response) error {
		responseBodyBytes, _ := ioutil.ReadAll(response.Body)
		log.Debugf("responseBodyBytes: %+v", string(responseBodyBytes))
		if response.StatusCode != http.StatusOK {
			return &common.HttpError{
				ErrorCode:    response.StatusCode,
				ResponseBody: string(responseBodyBytes),
				URL:          requestUrl,
			}
		}
		err := json.Unmarshal(responseBodyBytes, &responseBody)
		return err
	}
	header := map[string]string{
		"Content-Type": "application/json",
	}
	var body interface{} = map[string]string{
		"qq": qq,
	}
	err := c.httpClient.Post(requestUrl, &body, header, retryTimes, responseHandler)
	if err != nil {
		log.Errorf("postResponseJson failed, err is: %+v, qq is: %s", err, qq)
		return err
	}
	return nil
}

func (c *QQHttpClient) getResponseJson(qq string, headers map[string]string, parameters map[string]string, responseBody *model.QQImageModel) error {
	retryTimes := 3
	requestUrl := fmt.Sprintf(c.host, qq)
	responseHandler := func(response *http.Response) error {
		responseBodyBytes, _ := ioutil.ReadAll(response.Body)
		log.Debugf("responseBodyBytes: %+v", string(responseBodyBytes))
		if response.StatusCode != http.StatusOK {
			return &common.HttpError{
				ErrorCode:    response.StatusCode,
				ResponseBody: string(responseBodyBytes),
				URL:          requestUrl,
			}
		}
		contentType := response.Header.Get("content-type")
		if !strings.Contains(contentType, "image") {
			return fmt.Errorf("response contentType unlike image, qq is: %s", qq)
		}
		responseBody.Type = strings.Split(contentType, "/")[1]
		responseBody.Data = responseBodyBytes
		return nil
	}
	err := c.httpClient.Get(requestUrl, parameters, headers, retryTimes, responseHandler)
	if err != nil {
		log.Errorf("getResponseJson url: %s, params: %+v, err: %+v", requestUrl, parameters, err.Error())
		return err
	}
	return nil
}
