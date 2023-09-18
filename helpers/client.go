package helpers

import (
	"fmt"
	"io"
	"net/http"
)

var baseURL string = "https://api.github.com"

func NewRequest(method string, path string, body io.Reader) (*http.Request, error) {

	withPath := fmt.Sprintf("%s/%s", baseURL, path)

	req, err := http.NewRequest(method, withPath, body)

	return req, err
}

func Client(r *http.Request) (*http.Response, error) {
	c := http.Client{}

	res, err := c.Do(r)

	return res, err
}
