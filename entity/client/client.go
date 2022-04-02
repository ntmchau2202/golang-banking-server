package client

import (
	"bankserver/entity/message"
	"bankserver/entity/message/request"
	"bankserver/entity/message/response"
	"bytes"
	"encoding/json"
	"net/http"
)

type Client struct {
	cl      http.Client
	baseUri string
	headers http.Header
}

func NewClient(base string) (c *Client) {
	return &Client{
		cl:      http.Client{},
		baseUri: base,
	}
}

func (c *Client) getBaseUri() string {
	return c.baseUri
}

func (c *Client) SetHeader(key string, value string) {
	c.headers.Add(key, value)
}

func (c *Client) POST(uri string, body request.Request) (resp response.Response, err error) {
	url := c.getBaseUri() + uri
	requestBody, err := json.Marshal(body)
	if err != nil {
		return
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return
	}
	for key, value := range c.headers {
		for _, item := range value {
			req.Header.Add(key, item)
		}
	}
	fullResp, err := c.cl.Do(req)
	if err != nil {
		return
	}

	defer fullResp.Body.Close()
	var respBody []byte
	_, err = fullResp.Body.Read(respBody)
	if err != nil {
		return
	}
	resp.Details = make(map[string]interface{})
	err = json.Unmarshal(respBody, &resp.Details)
	if err != nil {
		return
	}

	if fullResp.StatusCode == 200 {
		resp.Stat = message.SUCCESS
	} else {
		resp.Stat = message.ERROR
	}
	return
}

func (c *Client) GET(uri string, body request.Request) (resp response.Response, err error) {
	url := c.getBaseUri() + uri
	requestBody, err := json.Marshal(body)
	if err != nil {
		return
	}
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return
	}
	for key, value := range c.headers {
		for _, item := range value {
			req.Header.Add(key, item)
		}
	}
	fullResp, err := c.cl.Do(req)
	if err != nil {
		return
	}

	defer fullResp.Body.Close()
	var respBody []byte
	_, err = fullResp.Body.Read(respBody)
	if err != nil {
		return
	}
	resp.Details = make(map[string]interface{})
	err = json.Unmarshal(respBody, &resp.Details)
	if err != nil {
		return
	}

	if fullResp.StatusCode == 200 {
		resp.Stat = message.SUCCESS
	} else {
		resp.Stat = message.ERROR
	}
	return
}
