package requests

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"time"

	"gitlab.hamkorbank.uz/libs/remote"
)

type OpenAPIConfig struct {
	Host           string
	CustomerKey    string
	CustomerSecret string
	AuthURL        string
}
type requestMiddleware int

func (r *requestMiddleware) BeforeRequest(ctx context.Context, req *http.Request) error {
	return nil
}
func (r *requestMiddleware) AfterResponse(res *http.Response) error {
	if res.StatusCode < http.StatusMultipleChoices {
		return nil
	}
	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("status: %s, %+v", res.Status, err)
	}
	return fmt.Errorf("status code %s, response body:%s", res.Status, string(respBody))
}

func NewOpenAPI(cfg OpenAPIConfig) remote.HTTPClientI {
	//nolint:gosec
	httpClient := remote.New(cfg.Host, &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // In case of missing certificates
		IdleConnTimeout: time.Second * 1,
	})
	cred := remote.WSOAMCredentials{
		Key:    cfg.CustomerKey,
		Secret: cfg.CustomerSecret,
		URL:    cfg.AuthURL,
	}
	//nolint:gosec
	m := remote.NewWSO2AMMiddleware(cred, &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		IdleConnTimeout: time.Second * 1, // In case of missing certificates
	})

	httpClient.SetMiddleware(m)
	httpClient.SetMiddleware(new(requestMiddleware))
	return httpClient

}

func CreateMultipartPostRequest(ctx context.Context, url string, data, header map[string]string) (*http.Request, error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	for k, v := range data {
		err := w.WriteField(k, v)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, &b)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-type", "multipart/form-data; boundary=<calculated when request is sent>")
	//req.Header.Set("Accept", "encoding/json")
	for k, v := range header {
		req.Header.Set(k, v)
	}
	return req, nil
}
