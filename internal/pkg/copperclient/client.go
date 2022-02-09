package copperclient

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// CopperClient copper API client
type CopperClient struct {
	client       http.Client
	secretKey    string
	apiKey       string
	baseURL      string
	portfolioID  string
	mainCurrency string
}

// NewCopperClient create new copper client
func NewCopperClient() *CopperClient {
	return &CopperClient{
		client: http.Client{
			Timeout: cfg.ClientTimeout(),
		},
		secretKey:    cfg.SecretKey(),
		apiKey:       cfg.APIKey(),
		baseURL:      cfg.BaseURL(),
		portfolioID:  cfg.PortfolioID(),
		mainCurrency: cfg.MainCurrency(),
	}
}

func (cc *CopperClient) getSign(method, path, body string, requestTime time.Time) string {
	h := hmac.New(sha256.New, []byte(cc.secretKey))
	h.Write([]byte(cc.getTimestamp(requestTime) + method + path + body))
	return hex.EncodeToString(h.Sum(nil))
}

func (cc *CopperClient) getTimestamp(t time.Time) string {
	return fmt.Sprintf("%d", t.UnixMilli())
}

func (cc *CopperClient) getAuthorizationHeader() string {
	return fmt.Sprintf("%s %s", HeaderPrefixAuthorization, cc.apiKey)
}

func (cc *CopperClient) getTimestampHeader(t time.Time) string {
	return cc.getTimestamp(t)
}

func (cc *CopperClient) addHeaders(req *http.Request, path, body string) {
	requestTime := time.Now()

	req.Header.Add(HeaderKeyContentType, HeaderValueContentType)
	req.Header.Add(HeaderKeyAuthorization, cc.getAuthorizationHeader())
	req.Header.Add(HeaderKeyTimestamp, cc.getTimestampHeader(requestTime))

	sign := cc.getSign(strings.ToUpper(req.Method), path, body, requestTime)
	req.Header.Add(HeaderKeySignature, sign)
}

func (cc *CopperClient) composeRequestURL(path string, params map[string]string) (string, error) {
	parsedBaseURL, err := url.Parse(cc.baseURL)
	if err != nil {
		return "", err
	}

	parsedBaseURL.Path = path
	if params == nil || len(params) > 0 {
		queryParams := parsedBaseURL.Query()

		for key, value := range params {
			queryParams.Add(key, value)
		}

		parsedBaseURL.RawQuery = queryParams.Encode()
	}

	return parsedBaseURL.String(), nil
}

func (cc *CopperClient) newRequest(method, path string, params map[string]string, body interface{}) (*http.Request, []byte, error) {
	fullURL, err := cc.composeRequestURL(path, params)
	if err != nil {
		return nil, nil, err
	}

	var req *http.Request
	var requestBodyRaw []byte
	if body != nil {
		requestBodyRaw, err = marshal(body)
		if err != nil {
			return nil, nil, err
		}

		req, err = http.NewRequest(method, fullURL, bytes.NewReader(requestBodyRaw))
	} else {
		requestBodyRaw = []byte{}
		req, err = http.NewRequest(method, fullURL, nil)
	}

	return req, requestBodyRaw, err
}

func (cc *CopperClient) request(method, path string, params map[string]string, body, responseBody interface{}) error {
	req, requestBodyRaw, err := cc.newRequest(method, path, params, body)
	if err != nil {
		return err
	}

	cc.addHeaders(req, path, string(requestBodyRaw))
	resp, err := cc.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBodyRaw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = cc.responsePostProcessor(resp.StatusCode, respBodyRaw, responseBody)
	if err != nil {
		return err
	}

	return nil
}

func (cc *CopperClient) responsePostProcessor(status int, respBodyRaw []byte, responseBody interface{}) error {
	if status < 300 {
		if responseBody != nil {
			err := unmarshal(respBodyRaw, responseBody)
			if err != nil {
				return err
			}
		}
	} else {
		respErr := ErrorResponse{}

		err := unmarshal(respBodyRaw, &respErr)
		if err != nil {
			return err
		} else {
			return respErr.GetError()
		}
	}

	return nil
}
