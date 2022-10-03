package client

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"sync"
	"time"

	"4d63.com/tz"

	"github.com/DNSMadeEasy/dme-go-client/container"
	"github.com/DNSMadeEasy/dme-go-client/models"
)

var mutex sync.Mutex

const BaseURL = "https://api.dnsmadeeasy.com/V2.0/"
const SleepDuration = 5

type Client struct {
	httpclient *http.Client
	apiKey     string //Required
	secretKey  string //Required
	insecure   bool   //Optional
	proxyurl   string //Optional
}

//singleton implementation of a client
var clientImpl *Client

//get first
type Option func(*Client)

func Insecure(insecure bool) Option {
	return func(client *Client) {
		client.insecure = insecure
	}
}

func ProxyUrl(pUrl string) Option {
	return func(client *Client) {
		client.proxyurl = pUrl
	}
}

func initClient(apiKey, secretKey string, options ...Option) *Client {
	//existing information about client
	client := &Client{
		apiKey:    apiKey,
		secretKey: secretKey,
	}
	for _, option := range options {
		option(client)
	}

	//Setting up the HTTP client for the API call
	var transport *http.Transport
	transport = client.useInsecureHTTPClient(client.insecure)
	if client.proxyurl != "" {
		transport = client.configProxy(transport)
	}
	client.httpclient = &http.Client{
		Transport: transport,
	}
	return client
}

//Returns a singleton
func GetClient(apiKey, secretKey string, options ...Option) *Client {
	if clientImpl != nil {
		return clientImpl
	}
	clientImpl = initClient(apiKey, secretKey, options...)
	return clientImpl
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
	pUrl, err := url.Parse(c.proxyurl)
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

	url := fmt.Sprintf("%s%s", BaseURL, endpoint)

	resp, err := c.doRequestWithRateLimit("POST", url, jsonPayload)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	respObj, err := container.ParseJSON(bodyBytes)
	if err != nil {
		return nil, err
	}
	log.Println("Respons body is :", respObj)

	respErr := checkForErrors(resp, respObj)
	if respErr != nil {
		return nil, respErr
	}
	return respObj, nil
}

func (c *Client) GetbyId(endpoint string) (*container.Container, error) {

	url := fmt.Sprintf("%s%s", BaseURL, endpoint)
	resp, err := c.doRequestWithRateLimit("GET", url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
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
	url := fmt.Sprintf("%s%s", BaseURL, endpoint)

	resp, err := c.doRequestWithRateLimit("PUT", url, jsonPayload)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return nil, nil
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
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
	url := fmt.Sprintf("%s%s", BaseURL, endpoint)
	var resp *http.Response

	resp, err := c.doRequestWithRateLimit("DELETE", url, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return nil
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	respObj, err := container.ParseJSON(bodyBytes)
	if err != nil {
		return err
	}
	log.Println("Respons body is :", respObj)

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

func getToken(apikey, secretkey string) string {
	//epoch time in milliseconds
	loc, _ := tz.LoadLocation("GMT")
	now := time.Now().In(loc)
	time := now.Format("Mon, 2 Jan 2006 15:04:05 MST")

	//generates hmac from secret key
	h := hmac.New(sha1.New, []byte(secretkey))
	h.Write([]byte(time))
	sha := hex.EncodeToString(h.Sum(nil))

	return string(sha)
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
		logRequest(req, resp)

		// Retry If Rate Limit is reached
		requestsRemaining, _ := strconv.Atoi(resp.Header.Get("x-dnsme-requestsRemaining"))
		if (err != nil) || (resp.StatusCode == 400 || resp.StatusCode == 404) && requestsRemaining == 0 {
			log.Println("waiting until more API calls can be done")
			time.Sleep(time.Duration(SleepDuration) * time.Second)
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

	hmac := getToken(c.apiKey, c.secretKey)
	loc, _ := tz.LoadLocation("GMT")
	now := time.Now().In(loc)
	time := now.Format("Mon, 2 Jan 2006 15:04:05 MST")

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-dnsme-hmac", hmac)
	req.Header.Set("x-dnsme-apiKey", c.apiKey)
	req.Header.Set("x-dnsme-requestDate", time)

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
