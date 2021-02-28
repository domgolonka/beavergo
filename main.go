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
	URL   string
}

func NewConnect(token string, url string) *ChatClient {

	client := ChatClient{
		Token: token,
		URL:   url,
	}

	return &client
}

func (c *ChatClient) command(method string, url string, payload string) ([]byte, error) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	var req *http.Request
	var err error
	fullurl := c.URL + url
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
	req.Header.Add("X-AUTH-TOKEN", c.Token)

	res, err := client.Do(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	return body, err
}

func (c *ChatClient) HealthCheck() (*Status, error) {
	method := "GET"
	url := "/_healthcheck"
	body, err := c.command(method, url, "")

	if err != nil {
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

func (c *ChatClient) CreateConfig(key string, value string) (bool, error) {
	method := "POST"
	url := "/api/config"
	payload := "{\n	\"key\" : \"" + key + "\",\n	\"value\" : \"" + value + "\"\n}"

	_, err := c.command(method, url, payload)

	if err != nil {
		return false, err
	}
	return true, nil
}

func (c *ChatClient) GetConfig(key string) (*Config, error) {
	method := "GET"
	url := "/api/config/"

	body, err := c.command(method, url+key, "")
	if err != nil {
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

func (c *ChatClient) UpdateConfig(value string) (bool, error) {
	method := "PUT"
	url := "/api/config/"
	payload := "{\n	\"value\" : \"" + value + "\"\n}"
	_, err := c.command(method, url+value, payload)

	if err != nil {
		return false, err
	}
	return true, nil
}

func (c *ChatClient) DeleteConfig(key string) (bool, error) {
	method := "DELETE"
	url := "/api/config/"

	_, err := c.command(method, url+key, "")

	if err != nil {
		return false, err
	}
	return true, nil
}

func (c *ChatClient) GetChannel(channel string) (*Channel, error) {
	method := "GET"
	url := "/api/channel/"

	body, err := c.command(method, url+channel, "")

	if err != nil {
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

func (c *ChatClient) CreateChannel(channel string, ctype string) (bool, error) {
	method := "POST"
	url := "/api/channel"
	payload := "{\n	\"name\" : \"" + channel + "\",\n	\"type\" : \"" + ctype + "\"\n}"
	_, err := c.command(method, url, payload)

	if err != nil {
		return false, err
	}
	return true, nil
}

func (c *ChatClient) UpdateChannel(channel string, ctype string) (bool, error) {
	method := "POST"
	url := "/api/channel/"
	payload := "{\n	\"type\" : \"" + ctype + "\"}"
	_, err := c.command(method, url+channel, payload)

	if err != nil {
		return false, err
	}
	return true, nil
}

// has to be json string
func (c *ChatClient) PublishChannel(channel string, data string) (bool, error) {
	method := "POST"
	url := "/api/publish"
	quoteddata := strconv.Quote(data)
	payload := "{\n	\"channel\" : \"" + channel + "\",\n	\"data\" : " + quoteddata + "\n}"
	_, err := c.command(method, url, payload)

	if err != nil {
		return false, err
	}
	return true, nil
}

// has to be json string
func (c *ChatClient) BroadcastChannel(channels []string, data string) (bool, error) {
	method := "POST"
	url := "/api/broadcast"
	urlsJSON, _ := json.Marshal(channels)
	quoteddata := strconv.Quote(data)
	payload := "{\n	\"channels\" : " + string(urlsJSON) + ",\n	\"data\" : " + quoteddata + "\n}"
	_, err := c.command(method, url, payload)

	if err != nil {
		return false, err
	}
	return true, nil
}

func (c *ChatClient) DeleteChannel(channel string) (bool, error) {
	method := "DELETE"
	url := "/api/channel/"

	_, err := c.command(method, url+channel, "")

	if err != nil {
		return false, err
	}
	return true, nil
}

func (c *ChatClient) CreateClient(channel []string) (*ClientResp, error) {
	method := "POST"
	url := "/api/client"
	urlsJSON, _ := json.Marshal(channel)
	payload := "{\n	\"channels\" : " + string(urlsJSON) + "\n}"
	body, err := c.command(method, url, payload)

	if err != nil {
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

func (c *ChatClient) GetClient(id string) (*ClientResp, error) {
	method := "GET"
	url := "/api/client/"

	body, err := c.command(method, url+id, "")

	if err != nil {
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

func (c *ChatClient) SubscribeClient(channel []string, id string) (bool, error) {
	method := "PUT"
	url := "/api/client/"
	urlsJSON, _ := json.Marshal(channel)

	payload := "{\n	\"channels\" : " + string(urlsJSON) + "\n}"
	_, err := c.command(method, url+id+"/subscribe", payload)

	if err != nil {
		return false, err
	}
	return true, nil
}
func (c *ChatClient) UnsubscribeClient(channel []string, id string) (bool, error) {
	method := "PUT"
	url := "/api/client/"
	urlsJSON, _ := json.Marshal(channel)
	payload := "{\n	\"channels\" : " + string(urlsJSON) + "\n}"
	_, err := c.command(method, url+id+"/unsubscribe", payload)

	if err != nil {
		return false, err
	}
	return true, err
}

func (c *ChatClient) DeleteClient(id string) (bool, error) {
	method := "DELETE"
	url := "/api/client/"

	_, err := c.command(method, url+id, "")

	if err != nil {
		return false, err
	}
	return true, nil
}

/* for future use */
func (c *ChatClient) Metrics() (bool, error) {
	method := "GET"
	url := "/api/metrics"
	_, err := c.command(method, url, "")
	if err != nil {
		return false, err
	}
	return true, nil
}

/* for future use */
func (c *ChatClient) Node() (bool, error) {
	method := "GET"
	url := "/api/node"
	_, err := c.command(method, url, "")
	if err != nil {
		return false, err
	}
	return true, nil
}
