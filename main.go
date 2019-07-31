package beavergo

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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
	
	var req *http.Request
	var err error
	fullurl := this.Url + url
	if payload == "" {
		req, err = http.NewRequest(method, fullurl, nil)
	} else {
		payday := strings.NewReader(payload)
		req, err = http.NewRequest(method, fullurl, payday)
	}

	if err != nil {
		log.Print(err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-AUTH-TOKEN", this.Token)

	res, err := client.Do(req)
	if (err != nil) {
		log.Print(err)
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	return body, nil
}

func (this *ChatClient) HealthCheck() (*Status, error) {
	method := "GET"
	url := "/_healthcheck"
	body, err := this.command(method, url, "")

	if (err != nil) {
		return nil, err
	}
	var status *Status
	err = json.Unmarshal(body, &status)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return status, nil
}

func (this *ChatClient) CreateConfig(key string, value string) (bool, error) {
	method := "POST"
	url := "/api/config"
	payload := "{\n	\"key\" : \"" + key + "\",\n	\"value\" : \"" + value + "\"\n}"

	_, err := this.command(method, url, payload)

	if (err != nil) {
		return false, err
	}
	return true, nil
}

func (this *ChatClient) GetConfig(key string) (*Config, error) {
	method := "GET"
	url := "/api/config/"

	body, err := this.command(method, url+key, "")
	if (err != nil) {
		return nil, err
	}

	var config *Config
	err = json.Unmarshal(body, &config)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return config, nil
}

func (this *ChatClient) UpdateConfig(value string) (bool, error) {
	method := "PUT"
	url := "/api/config/"
	payload := "{\n	\"value\" : \"" + value + "\"\n}"
	_, err := this.command(method, url+value, payload)

	if (err != nil) {
		return false, err
	}
	return true, nil
}

func (this *ChatClient) DeleteConfig(key string) (bool, error) {
	method := "DELETE"
	url := "/api/config/"

	_, err := this.command(method, url+key, "")

	if (err != nil) {
		return false, err
	}
	return true, nil
}

func (this *ChatClient) GetChannel(channel string) (*Channel, error) {
	method := "GET"
	url := "/api/channel/"

	body, err := this.command(method, url+channel, "")

	if (err != nil) {
		return nil, err
	}
	var chanresp *Channel
	err = json.Unmarshal(body, &chanresp)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return chanresp, nil
}

func (this *ChatClient) CreateChannel(channel string, ctype string) (bool, error) {
	method := "POST"
	url := "/api/channel"
	payload := "{\n	\"name\" : \"" + channel + "\",\n	\"type\" : \"" + ctype + "\"\n}"
	_, err := this.command(method, url, payload)

	if (err != nil) {
		return false, err
	}
	return true, nil
}

func (this *ChatClient) UpdateChannel(channel string, ctype string) (bool, error) {
	method := "POST"
	url := "/api/channel/"
	payload := "{\n	\"type\" : \"" + ctype + "\"}"
	_, err := this.command(method, url+channel, payload)

	if (err != nil) {
		return false, err
	}
	return true, nil
}

// has to be json string
func (this *ChatClient) PublishChannel(channel string, data string) (bool, error) {
	method := "POST"
	url := "/api/publish"
	quoteddata := strconv.Quote(data)
	payload := "{\n	\"channel\" : \"" + channel + "\",\n	\"data\" : " + quoteddata + "\n}"
	_, err := this.command(method, url, payload)

	if (err != nil) {
		return false, err
	}
	return true, nil
}

// has to be json string
func (this *ChatClient) BroadcastChannel(channels []string, data string) (bool, error) {
	method := "POST"
	url := "/api/broadcast"
	urlsJson, _ := json.Marshal(channels)
	quoteddata := strconv.Quote(data)
	payload := "{\n	\"channels\" : " + string(urlsJson) + ",\n	\"data\" : " + quoteddata + "\n}"
	_, err := this.command(method, url, payload)

	if (err != nil) {
		return false, err
	}
	return true, nil
}

func (this *ChatClient) DeleteChannel(channel string) (bool, error) {
	method := "DELETE"
	url := "/api/channel/"

	_, err := this.command(method, url+channel, "")

	if (err != nil) {
		return false, err
	}
	return true, nil
}

func (this *ChatClient) CreateClient(channel []string) (*ClientResp, error) {
	method := "POST"
	url := "/api/client"
	urlsJson, _ := json.Marshal(channel)
	payload := "{\n	\"channels\" : " + string(urlsJson) + "\n}"
	body, err := this.command(method, url, payload)

	if (err != nil) {
		return nil, err
	}
	var clientresp *ClientResp
	err = json.Unmarshal(body, &clientresp)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return clientresp, nil
}

func (this *ChatClient) GetClient(id string) (*ClientResp, error) {
	method := "GET"
	url := "/api/client/"

	body, err := this.command(method, url+id, "")

	if (err != nil) {
		return nil, err
	}
	var clientresp *ClientResp
	err = json.Unmarshal(body, &clientresp)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return clientresp, nil
}

func (this *ChatClient) SubscribeClient(channel []string, id string) (bool, error) {
	method := "PUT"
	url := "/api/client/"
	urlsJson, _ := json.Marshal(channel)

	payload := "{\n	\"channels\" : " + string(urlsJson) + "\n}"
	_, err := this.command(method, url+id+"/subscribe", payload)

	if (err != nil) {
		return false, err
	}
	return true, nil
}
func (this *ChatClient) UnsubscribeClient(channel []string, id string) (bool, error) {
	method := "PUT"
	url := "/api/client/"
	urlsJson, _ := json.Marshal(channel)
	payload := "{\n	\"channels\" : " + string(urlsJson) + "\n}"
	_, err := this.command(method, url+id+"/unsubscribe", payload)

	if (err != nil) {
		return false, err
	}
	return true, err
}

func (this *ChatClient) DeleteClient(id string) (bool, error) {
	method := "DELETE"
	url := "/api/client/"

	_, err := this.command(method, url+id, "")

	if (err != nil) {
		return false, err
	}
	return true, nil
}

/* for future use */
func (this *ChatClient) Metrics() (bool, error) {
	method := "GET"
	url := "/api/metrics"
	_, err := this.command(method, url, "")
	if (err != nil) {
		return false, err
	}
	return true, nil
}

/* for future use */
func (this *ChatClient) Node() (bool, error) {
	method := "GET"
	url := "/api/node"
	_, err := this.command(method, url, "")
	if (err != nil) {
		return false, err
	}
	return true, nil
}
