package tesla

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	httpClient   *http.Client
	Token        *Token
}

var (
	BaseURL      = "https://owner-api.teslamotors.com/api/1"
)

func NewClient() *Client {
	c := &Client{}
	c.httpClient = &http.Client{}

	return c
}

func (c *Client) post(url string, body []byte) ([]byte, error) {
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))

	if c.Token != nil {
		req.Header.Set("Authorization", "Bearer " + c.Token.AccessToken)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid status code: %d", res.StatusCode)
	}
	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}

func (c *Client) get(url string) ([]byte, error) {
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	if c.Token != nil {
		req.Header.Set("Authorization", "Bearer " + c.Token.AccessToken)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid status code: %d", res.StatusCode)
	}
	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}


