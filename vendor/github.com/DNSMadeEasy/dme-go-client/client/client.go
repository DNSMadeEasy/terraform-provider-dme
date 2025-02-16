package client

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"4d63.com/tz"

	"github.com/DNSMadeEasy/dme-go-client/container"
	"github.com/DNSMadeEasy/dme-go-client/models"
)

var mutex sync.Mutex

const defaultBaseURL = "https://api.dnsmadeeasy.com/V2.0"
const sleepDuration = 5

type Client struct {
	httpclient *http.Client
	apiKey     string // Required
	secretKey  string // Required
	insecure   bool   // Optional
	proxyURL   string // Optional
	baseURL    string // Optional
}

// Singleton implementation of a client
var clientImpl *Client

// Option for the client
type Option func(*Client)

func Insecure(insecure bool) Option {
	return func(client *Client) {
		client.insecure = insecure
	}
}

func ProxyURL(proxyURL string) Option {
	return func(client *Client) {
		client.proxyURL = proxyURL
	}
}

func BaseURL(baseURL string) Option {
	return func(client *Client) {
		client.baseURL = sanitizeURL(baseURL)
	}
}

func sanitizeURL(rawUrl string) string {
	// Remove leading/trailing white spaces.
	trimmedUrl := strings.TrimSpace(rawUrl)
	// Remove trailing slashes.
	trimmedUrl = strings.TrimRight(trimmedUrl, "/")
	if trimmedUrl == "" {
		return ""
	}

	// If no scheme is present, prepend "https://"
	if !strings.Contains(trimmedUrl, "://") {
		trimmedUrl = "https://" + trimmedUrl
	}
	return trimmedUrl
}

func initClient(apiKey, secretKey string, options ...Option) *Client {
	// Existing information about client
	client := &Client{
		apiKey:    apiKey,
		secretKey: secretKey,
	}
	for _, option := range options {
		option(client)
	}

	// Setting up the HTTP client for the API call
	var transport *http.Transport
	transport = client.useInsecureHTTPClient(client.insecure)
	if client.proxyURL != "" {
		transport = client.configProxy(transport)
	}
	client.httpclient = &http.Client{
		Transport: transport,
	}
	if client.baseURL == "" {
		client.baseURL = defaultBaseURL
	}
	return client
}

// GetClient Returns a singleton
func GetClient(apiKey, secretKey string, options ...Option) *Client {
	if clientImpl != nil {
		return clientImpl
	}
	clientImpl = initClient(apiKey, secretKey, options...)
	return clientImpl
}

func resetClient() {
	clientImpl = nil
}

func (c *Client) toAPIEndpoint(endpoint string) string {
	return fmt.Sprintf("%s/%s", c.baseURL, endpoint)
}

func (c *Client) useInsecureHTTPClient(insecure bool) *http.Transport {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
				tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
				tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			},
			PreferServerCipherSuites: true,
			InsecureSkipVerify:       insecure,
			MinVersion:               tls.VersionTLS11,
			MaxVersion:               tls.VersionTLS12,
		},
	}

	return transport
}

func (c *Client) configProxy(transport *http.Transport) *http.Transport {
	pUrl, err := url.Parse(c.proxyURL)
	if err != nil {
		log.Fatal(err)
	}
	transport.Proxy = http.ProxyURL(pUrl)
	return transport
}

func (c *Client) Save(obj models.Model, endpoint string) (*container.Container, error) {
	jsonPayload, err := c.PrepareModel(obj)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequestWithRateLimit("POST", c.toAPIEndpoint(endpoint), jsonPayload)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	respObj, err := container.ParseJSON(bodyBytes)
	if err != nil {
		return nil, err
	}
	log.Println("Response body is :", respObj)

	respErr := checkForErrors(resp, respObj)
	if respErr != nil {
		return nil, respErr
	}
	return respObj, nil
}

func (c *Client) GetbyId(endpoint string) (*container.Container, error) {
	resp, err := c.doRequestWithRateLimit("GET", c.toAPIEndpoint(endpoint), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	respObj, err := container.ParseJSON(bodyBytes)
	if err != nil {
		return nil, err
	}
	log.Println("Response body is :", respObj)

	respErr := checkForErrors(resp, respObj)
	if respErr != nil {
		return nil, respErr
	}
	return respObj, nil
}

func (c *Client) Update(obj models.Model, endpoint string) (*container.Container, error) {
	jsonPayload, err := c.PrepareModel(obj)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequestWithRateLimit("PUT", c.toAPIEndpoint(endpoint), jsonPayload)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return nil, nil
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	respObj, err := container.ParseJSON(bodyBytes)
	if err != nil {
		return nil, err
	}
	log.Println("Response body is :", respObj)

	respErr := checkForErrors(resp, respObj)
	if respErr != nil {
		return nil, respErr
	}
	return respObj, nil
}

func (c *Client) Delete(endpoint string) error {
	var resp *http.Response

	resp, err := c.doRequestWithRateLimit("DELETE", c.toAPIEndpoint(endpoint), nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return nil
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	respObj, err := container.ParseJSON(bodyBytes)
	if err != nil {
		return err
	}
	log.Println("Response body is :", respObj)

	respErr := checkForErrors(resp, respObj)
	if respErr != nil {
		return respErr
	}
	return nil
}

func checkForErrors(resp *http.Response, obj *container.Container) error {
	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		log.Println(" Into the check for errors ")
		if resp.StatusCode == 404 {
			// Not fixing capitalization issue on next line,
			// as it might break some upstream code.
			return fmt.Errorf("Particular item not found")
		}

		errs := obj.S("error").Data().([]interface{})

		var allErrors string
		for _, tp := range errs {
			allErrors = allErrors + tp.(string)
		}
		return fmt.Errorf("%s", allErrors)
	}
	return nil
}

func (c *Client) PrepareModel(obj models.Model) (*container.Container, error) {
	con := obj.ToMap()

	payload := &container.Container{}

	for key, value := range con {
		payload.Set(value, key)
	}
	return payload, nil
}

func getToken(secretKey string) string {
	// Epoch time in milliseconds
	loc, _ := tz.LoadLocation("GMT")
	now := time.Now().In(loc)
	timestamp := now.Format("Mon, 2 Jan 2006 15:04:05 MST")

	// Generates hmac from secret key
	h := hmac.New(sha1.New, []byte(secretKey))
	h.Write([]byte(timestamp))
	return hex.EncodeToString(h.Sum(nil))
}

func (c *Client) doRequestWithRateLimit(method, endpoint string, con *container.Container) (*http.Response, error) {
	var resp *http.Response
	mutex.Lock()
	defer mutex.Unlock()
	for {
		req, err := c.makeRequest(method, endpoint, con)
		if err != nil {
			return nil, err
		}

		reqDump, err := httputil.DumpRequestOut(req, true)
		if err != nil {
			log.Println(err)
		}
		log.Printf("[DEBUG] \n--[ HTTP Request Sent]------------------------------------ \n %s\n---------------------------------------------\n", string(reqDump))

		// DO
		resp, err = c.httpclient.Do(req)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		logRequest(req, resp)

		// Retry If Rate Limit is reached
		requestsRemaining, _ := strconv.Atoi(resp.Header.Get("x-dnsme-requestsRemaining"))
		if (resp.StatusCode == 400 || resp.StatusCode == 404) && requestsRemaining == 0 {
			log.Println("waiting until more API calls can be done")
			time.Sleep(time.Duration(sleepDuration) * time.Second)
			continue
		}
		time.Sleep(time.Duration(100) * time.Millisecond)
		break
	}
	return resp, nil
}

func (c *Client) makeRequest(method, endpoint string, con *container.Container) (*http.Request, error) {
	var req *http.Request
	var err error
	if method == "POST" || method == "PUT" {
		req, err = http.NewRequest(method, endpoint, bytes.NewBuffer(con.Bytes()))
	} else {
		req, err = http.NewRequest(method, endpoint, nil)
	}
	if err != nil {
		return nil, err
	}

	hash := getToken(c.secretKey)
	loc, _ := tz.LoadLocation("GMT")
	now := time.Now().In(loc)
	timestamp := now.Format("Mon, 2 Jan 2006 15:04:05 MST")

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-dnsme-hmac", hash)
	req.Header.Set("x-dnsme-apiKey", c.apiKey)
	req.Header.Set("x-dnsme-requestDate", timestamp)

	return req, nil
}

func logRequest(req *http.Request, resp *http.Response) {
	reqDump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		log.Println(err)
	}
	respDump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Println(err)
	}
	log.Printf("[DEBUG] \n--[ HTTP Request ]------------------------------------ \n %s\n---------------------------------------------\n", string(reqDump))
	log.Printf("[DEBUG] \n--[ HTTP Response ]----------------------------------- \n %s\n---------------------------------------------\n", string(respDump))
}
