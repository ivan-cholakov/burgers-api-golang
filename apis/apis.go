package apis

import (
	"io/ioutil"
	"net/http"
)

type HttpInterface interface {
	Get(string) ([]byte, error)
}

type HttpClient struct{}

func (c *HttpClient) Get(url string) ([]byte, error) {
	return httpGet(url)
}

func httpGet(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	return ioutil.ReadAll((res.Body))
}
