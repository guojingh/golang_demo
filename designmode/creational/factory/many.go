package factory

import (
	"net/http"
	"net/http/httptest"
)

type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}

// NewHTTPClient 返回一个 net/http包提供的HTTP客户端
func NewHTTPClient() Doer {
	return &http.Client{}
}

type mockHTTPClient struct{}

func (*mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	res := httptest.NewRecorder()
	return res.Result(), nil
}

// NewMockHTTPClient 是返回一个模拟的 HTTP 客户端，该客户端接收任何请求，并返回一个空的响应
func NewMockHTTPClient() Doer {
	return &mockHTTPClient{}
}

func QueryUser(doer Doer) error {
	req, err := http.NewRequest("GET", "http://iam.api.marmotdeu.com:8080/v1/secrets", nil)
	if err != nil {
		return err
	}

	_, err = doer.Do(req)
	if err != nil {
		return err
	}

	return nil
}
