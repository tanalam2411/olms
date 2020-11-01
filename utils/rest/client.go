package rest

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

type ClientContentConfig struct {
	// ContentType specifies the wire format used to communicate with the server.
	// This value will be set as the Accept header on requests made to the server.
	ContentType string
}

type RESTClient struct {
	// base is the root URL for all invocations of the client
	base *url.URL

	// content describes how a RESTClient encodes and decodes responses
	content ClientContentConfig

	// Set specific behaviour of the client.
	Client *http.Client
}

func NewRESTClient(baseURL *url.URL, config ClientContentConfig, client *http.Client) (*RESTClient, error) {

	if len(config.ContentType) == 0 {
		config.ContentType = "application/json"
	}

	base := *baseURL

	return &RESTClient{
		base:    &base,
		content: config,
		Client:  client,
	}, nil
}

// temp function, will be replaced by rest client interface
func HttpGET(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
