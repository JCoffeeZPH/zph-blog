package client

import (
	"bytes"
	"fmt"
	"net"
	"zph/logger"

	"io/ioutil"
	"net/http"
	"time"

	"github.com/json-iterator/go"
)
var log = logger.Log

var retryStatusCodes = []int{
	http.StatusBadGateway,
	http.StatusServiceUnavailable,
	http.StatusGatewayTimeout,
	http.StatusInternalServerError,
}

type HttpError struct {
	ErrorCode    int
	ResponseBody string
	URL          string
}

func (httpError HttpError) Error() string {
	return fmt.Sprintf("ErrorCode: %d, ResponseBody: %s, URL: %s", httpError.ErrorCode, httpError.ResponseBody, httpError.URL)
}

type HttpClient struct {
	client *http.Client
}

func (h *HttpClient) GetHttpClient() *http.Client {
	return h.client
}

func NewHttpClient(maxConnectionNum int, dialTimeout time.Duration, requestTimeout time.Duration) *HttpClient {
	defaultTransport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   dialTimeout,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	defaultTransport.MaxIdleConns = maxConnectionNum
	defaultTransport.MaxIdleConnsPerHost = maxConnectionNum
	httpClient := &HttpClient{&http.Client{Transport: defaultTransport}}
	httpClient.client.Timeout = requestTimeout
	return httpClient
}

func (httpClient *HttpClient) Post(
	url string,
	body *interface{},
	headers map[string]string,
	retryTimes int,
	responseHandler func(response *http.Response) error) error {
	bodyBytes, err := jsoniter.Marshal(body)
	if err != nil {
		return err
	}
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return err
	}
	log.Debugf(request.URL.String())
	response, err := httpClient.httpRequest(request, headers, retryTimes)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	return responseHandler(response)
}

func (httpClient *HttpClient) PostForQQ(
	url string,
	body *interface{},
	headers map[string]string,
	retryTimes int,
	responseHandler func(response *http.Response) error) error {
	bodyBytes, err := jsoniter.Marshal(body)
	if err != nil {
		return err
	}
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return err
	}
	log.Debugf(request.URL.String())
	response, err := httpClient.httpRequest(request, headers, retryTimes)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	return responseHandler(response)
}

func (httpClient *HttpClient) PostResponseJson(
	url string,
	body *interface{},
	headers map[string]string,
	retryTimes int,
	responseBody interface{}) error {
	responseHandler := func(response *http.Response) error {
		responseBodyBytes, _ := ioutil.ReadAll(response.Body)
		if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusNoContent {
			return &HttpError{response.StatusCode, string(responseBodyBytes), url}
		}
		if responseBody != nil && responseBodyBytes != nil && len(responseBodyBytes) > 0 {
			return jsoniter.Unmarshal(responseBodyBytes, responseBody)
		}
		return nil
	}
	err := httpClient.Post(url, body, headers, retryTimes, responseHandler)
	return err
}

func (httpClient *HttpClient) Get(
	url string,
	parameters map[string]string,
	headers map[string]string,
	retryTimes int,
	responseHandler func(response *http.Response) error) error {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	query := request.URL.Query()
	for key, value := range parameters {
		query.Add(key, value)
	}
	request.URL.RawQuery = query.Encode()
	response, err := httpClient.httpRequest(request, headers, retryTimes)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	return responseHandler(response)
}

func (httpClient *HttpClient) Put(
	url string,
	body *interface{},
	headers map[string]string,
	retryTimes int,
	responseHandler func(response *http.Response) error) error {
	bodyBytes, err := jsoniter.Marshal(body)
	if err != nil {
		return err
	}

	request, err := http.NewRequest("PUT", url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return err
	}
	log.Debugf(request.URL.String())
	response, err := httpClient.httpRequest(request, headers, retryTimes)
	defer response.Body.Close()
	if err != nil {
		return err
	}

	return responseHandler(response)
}

func (httpClient *HttpClient) GetResponseJson(
	url string,
	parameters map[string]string,
	headers map[string]string,
	retryTimes int,
	responseBody interface{}) error {
	responseHandler := func(response *http.Response) error {
		responseBodyBytes, _ := ioutil.ReadAll(response.Body)
		if response.StatusCode != http.StatusOK {
			return &HttpError{response.StatusCode, string(responseBodyBytes), url}
		}
		err := jsoniter.Unmarshal(responseBodyBytes, responseBody)
		if err != nil {
			log.Errorf("Unmarshal response error, response: %v, %v", string(responseBodyBytes), err)
		}
		return err
	}
	return httpClient.Get(url, parameters, headers, retryTimes, responseHandler)
}

func (httpClient *HttpClient) httpRequest(request *http.Request, headers map[string]string, retryTimes int) (*http.Response, error) {
	runTimes := retryTimes
	if runTimes <= 0 {
		runTimes = 1
	}
	for key, value := range headers {
		request.Header.Set(key, value)
	}
	var response *http.Response = nil
	var err error = nil
	var body []byte
	if request.Body != nil {
		body, _ = ioutil.ReadAll(request.Body)
	}
	for i := 0; i < runTimes; i += 1 {
		if len(body) > 0 {
			request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		}
		response, err = httpClient.client.Do(request)
		if err != nil {
			continue
		}
		if !isRetry(response.StatusCode) {
			break
		}
		if response != nil {
			response.Body.Close()
		}
	}
	return response, err
}

func isRetry(statusCode int) bool {
	for _, s := range retryStatusCodes {
		if statusCode == s {
			return true
		}
	}
	return false
}

func (httpClient *HttpClient) EnsureHttpServer(url string) {
	err := httpClient.GetResponseJson(url, nil, nil, 0, nil)
	for err != nil {
		log.Errorf("ensuring http server: url: %s, err: %+v", url, err)
		if httpError, ok := err.(*HttpError); ok && httpError.ErrorCode == http.StatusBadGateway {
			err = httpClient.GetResponseJson(url, nil, nil, 0, nil)
		} else {
			break
		}
	}
}
