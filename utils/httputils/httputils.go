package httputils

import (
	"io"
	"net/http"
	"strings"
)

func Get(url string, body string, headers map[string]string) (string, error) {
	return doRequestWithJsonString(http.MethodGet, url, body, headers)
}

func Post(url string, body string, headers map[string]string) (string, error) {
	return doRequestWithJsonString(http.MethodPost, url, body, headers)
}

func Put(url string, body string, headers map[string]string) (string, error) {
	return doRequestWithJsonString(http.MethodPut, url, body, headers)
}

func Delete(url string, body string, headers map[string]string) (string, error) {
	return doRequestWithJsonString(http.MethodDelete, url, body, headers)
}

func doRequestWithJsonString(method string, url string, body string, headers map[string]string) (string, error) {
	headers["Content-Type"] = "application/json"
	bodyReader := strings.NewReader(body)
	return doRequest(method, url, bodyReader, headers)
}

func doRequest(method string, url string, body io.Reader, headers map[string]string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return "", err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	response, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	responseStr := string(response)
	return responseStr, nil
}
