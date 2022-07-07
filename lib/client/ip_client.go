package client

import (
	"encoding/json"
	"github.com/WillVi/ipgo"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"zph/config"
	"zph/constants"
	"zph/lib/client/model"
	"zph/lib/common"
)

type IPClient interface {
	GetIPAndAttribution() (*model.IPDetails, error)
}

func NewIPClient() IPClient {
	return &IpHttpClient{
		httpClient: NewHttpClient(10, 10*time.Second, 10*time.Second),
		host:       config.GetParseIpUrl(),
	}
}

type IpHttpClient struct {
	httpClient *HttpClient
	host       string
}

func (c *IpHttpClient) GetIPAndAttribution() (*model.IPDetails, error) {
	queryUrl := c.host
	resp := &model.IPDetails{}
	err := c.getResponseJson(queryUrl, nil, nil, resp)
	if err != nil {
		data, e := c.getIp2()
		if e != nil {
			log.Errorf("get ip failed, err is: %+v", e)
			resp.IP = config.GetDefaultIp()
		} else {
			res := strings.Split(data, "\r\n")
			resp.IP = res[0][strings.Index(res[0], ":")+1:]
		}
	}
	if len(resp.IP) == 0 {
		return resp, nil
	}
	ipgo.New("./ipdb/ip2region.db")
	search, _ := ipgo.BtreeSearch(resp.IP)
	ipSource := strings.ReplaceAll(strings.ReplaceAll(search.Region, "|0", ""), "0|", "")
	if strings.Contains(ipSource, constants.InternalNetIp) {
		responseBody, e := c.GetIpRegionMsg(resp.IP)
		if e != nil || responseBody.Info != "OK" {
			log.Errorf("GetIpRegionMsg failed, responseBody is: %+v, err is: %+v", responseBody, err)
			panic(e)
		} else {
			ipSource = responseBody.Province + "|" + responseBody.City
		}
	}
	resp.IpSource = ipSource

	return resp, nil
}

func (c *IpHttpClient) getResponseJson(requestUrl string, headers map[string]string, parameters map[string]string, responseBody *model.IPDetails) error {
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
		res := string(responseBodyBytes)
		responseBody.IP = res
		return nil
	}
	err := c.httpClient.Get(requestUrl, parameters, headers, retryTimes, responseHandler)
	if err != nil {
		log.Errorf("getResponseJson url: %s, params: %+v, err: %+v", requestUrl, parameters, err.Error())
		return err
	}
	return nil
}

func (c *IpHttpClient) getIp2() (string, error) {
	requestUrl := config.GetParseIpUrl2()
	retryTimes := 3
	var ip string
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
		ip = string(responseBodyBytes)
		return nil
	}
	err := c.httpClient.Get(requestUrl, nil, nil, retryTimes, responseHandler)
	if err != nil {
		log.Errorf("getResponseJson url: %s, err: %+v", requestUrl, err.Error())
		return "", err
	}
	return ip, nil
}
func (c *IpHttpClient) GetIpRegionMsg(ip string) (model.RegionData, error) {
	requestUrl := config.GetRegionMsgUrl()
	params := map[string]string{
		"ip":  ip,
		"key": "f4cf14aca974dfbb0501c582ce3fce77",
	}
	responseBody := model.RegionData{}
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
		json.Unmarshal(responseBodyBytes, &responseBody)
		return nil
	}
	err := c.httpClient.Get(requestUrl, params, nil, 3, responseHandler)
	log.Infof("GetIpRegionMsg, responseBody is: %+v, err is: %+v", responseBody, err)
	return responseBody, err
}
