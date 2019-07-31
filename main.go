package beavergo

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type ChatClient struct {
	Token string
	Url   string
}

func NewConnect(token string, url string) *ChatClient {

	client := ChatClient{
		Token: token,
		Url:   url,
	}

	return &client
}

func (this *ChatClient) command(method string, url string, payload string) ([]byte, error) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	logs.Error(payload)
	var req *http.Request
	var err error
	if payload == "" {
		req, err = http.NewRequest(method, url, nil)
	} else {
		payday := strings.NewReader(payload)
		req, err = http.NewRequest(method, url, payday)
	}

	if err != nil {
		log.Print(err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-AUTH-TOKEN", this.Token)

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	return body, nil
}

func (this *ChatClient) HealthCheck() (*Status, error) {
	method := "GET"
	url := "/_healthcheck"
	body, err := this.command(method, url, "")
	var status *Status
	err = json.Unmarshal(body, &status)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return status, nil
}

func (this *ChatClient) CreateConfig(appname string) (*http.ConnState, error) {
	method := "POST"
	url := "/api/config"
	payload := "{\n	\"key\" : \"app_name\",\n	\"value\" : \"" + appname + "\"\n}"

	body, err := this.command(method, url, payload)

	var config *http.ConnState
	err = json.Unmarshal(body, &config)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return config, nil
}

func (this *ChatClient) GetConfig(appname string) (*Config, error) {
	method := "GET"
	url := "/api/config/"

	body, err := this.command(method, url+appname, "")

	var config *Config
	err = json.Unmarshal(body, &config)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return config, nil
}

func (this *ChatClient) UpdateConfig(appname string) (*http.ConnState, error) {
	method := "PUT"
	url := "/api/config/"
	payload := "{\n	\"value\" : \"" + appname + "\"\n}"
	body, err := this.command(method, url+appname, payload)

	var config *http.ConnState
	err = json.Unmarshal(body, &config)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return config, nil
}

func (this *ChatClient) DeleteConfig(appname string) (*http.ConnState, error) {
	method := "DELETE"
	url := "/api/config/"

	body, err := this.command(method, url+appname, "")

	var config *http.ConnState
	err = json.Unmarshal(body, &config)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return config, nil
}

func (this *ChatClient) GetChannel(channel string) (*Channel, error) {
	method := "GET"
	url := "/api/channel/"

	body, err := this.command(method, url+channel, "")

	var chanresp *Channel
	err = json.Unmarshal(body, &chanresp)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return chanresp, nil
}

func (this *ChatClient) CreateChannel(channel string, ctype string) (*http.ConnState, error) {
	method := "POST"
	url := "/api/channel"
	payload := "{\n	\"name\" : \"" + channel + "\",\n	\"value\" : \"" + ctype + "\"\n}"
	body, err := this.command(method, url, payload)

	var chanresp *http.ConnState
	err = json.Unmarshal(body, &chanresp)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return chanresp, nil
}

func (this *ChatClient) UpdateChannel(channel string, ctype string) (*http.ConnState, error) {
	method := "POST"
	url := "/api/channel/"
	payload := "{\n	\"type\" : \"" + ctype + "\"}"
	body, err := this.command(method, url+channel, payload)

	var chanresp *http.ConnState
	err = json.Unmarshal(body, &chanresp)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return chanresp, nil
}

// has to be json string
func (this *ChatClient) PublishChannel(channel string, data string) (*http.ConnState, error) {
	method := "POST"
	url := "/api/publish"
	payload := "{\n	\"channel\" : \"" + channel + "\",\n	\"data\" : \"" + data + "\"\n}"
	body, err := this.command(method, url, payload)

	var chanresp *http.ConnState
	err = json.Unmarshal(body, &chanresp)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return chanresp, nil
}

// has to be json string
func (this *ChatClient) BroadcastChannel(channels []string, data string) (*http.ConnState, error) {
	method := "POST"
	url := "/api/publish"
	urlsJson, _ := json.Marshal(channels)
	payload := "{\n	\"channels\" : \"" + string(urlsJson) + "\",\n	\"data\" : \"" + data + "\"\n}"
	body, err := this.command(method, url, payload)

	var chanresp *http.ConnState
	err = json.Unmarshal(body, &chanresp)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return chanresp, nil
}

func (this *ChatClient) DeleteChannel(channel string) (*http.ConnState, error) {
	method := "DELETE"
	url := "/api/channel/"

	body, err := this.command(method, url+channel, "")

	var chanresp *http.ConnState
	err = json.Unmarshal(body, &chanresp)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return chanresp, nil
}

func (this *ChatClient) CreateClient(channel []string) (*ClientResp, error) {
	method := "POST"
	url := "/api/client"
	urlsJson, _ := json.Marshal(channel)
	payload := "{\n	\"channels\" : \"" + string(urlsJson) + "\"\n}"
	body, err := this.command(method, url, payload)

	var clientresp *ClientResp
	err = json.Unmarshal(body, &clientresp)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return clientresp, nil
}

func (this *ChatClient) GetClient(token string) (*ClientResp, error) {
	method := "GET"
	url := "/api/client/"

	body, err := this.command(method, url+token, "")

	var clientresp *ClientResp
	err = json.Unmarshal(body, &clientresp)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return clientresp, nil
}

func (this *ChatClient) SubscribeClient(channel []string, token string) (*http.ConnState, error) {
	method := "POST"
	url := "/api/client/"
	urlsJson, _ := json.Marshal(channel)
	payload := "{\n	\"type\" : \"" + string(urlsJson) + "\"}"
	body, err := this.command(method, url+token+"/subscribe", payload)

	var clientresp *http.ConnState
	err = json.Unmarshal(body, &clientresp)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return clientresp, nil
}
func (this *ChatClient) UnsubscribeClient(channel []string, token string) (*http.ConnState, error) {
	method := "POST"
	url := "/api/client/"
	urlsJson, _ := json.Marshal(channel)
	payload := "{\n	\"type\" : \"" + string(urlsJson) + "\"}"
	body, err := this.command(method, url+token+"/unsubscribe", payload)

	var clientresp *http.ConnState
	err = json.Unmarshal(body, &clientresp)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return clientresp, nil
}

func (this *ChatClient) DeleteClient(token string) (*http.ConnState, error) {
	method := "DELETE"
	url := "/api/client/"

	body, err := this.command(method, url+token, "")

	var clientresp *http.ConnState
	err = json.Unmarshal(body, &clientresp)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return clientresp, nil
}

/* for future use */
func (this *ChatClient) Metrics() (*http.ConnState, error) {
	method := "GET"
	url := "/api/metrics"
	body, err := this.command(method, url, "")
	var clientresp *http.ConnState
	err = json.Unmarshal(body, &clientresp)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return clientresp, nil
}

/* for future use */
func (this *ChatClient) Node() (*http.ConnState, error) {
	method := "GET"
	url := "/api/node"
	body, err := this.command(method, url, "")
	var clientresp *http.ConnState
	err = json.Unmarshal(body, &clientresp)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return clientresp, nil
}