package api

import (
	"bytes"
	"github.com/hashicorp/go-cleanhttp"
	"io"
	"k8s.io/apimachinery/pkg/util/json"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const (
	HTTPAddrEnvName      = "LAUNCHBOX_HTTP_ADDR"
	HTTPAuthEnvName      = "LAUNCHBOX_HTTP_AUTH"
	HTTPSSLEnvName       = "LAUNCHBOX_HTTP_SSL"
	HTTPTLSServerName    = "LAUNCHBOX_TLS_SERVER_NAME"
	HTTPCAFile           = "LAUNCHBOX_CACERT"
	HTTPCAPath           = "LAUNCHBOX_CAPATH"
	HTTPClientCert       = "LAUNCHBOX_CLIENT_CERT"
	HTTPClientKey        = "LAUNCHBOX_CLIENT_KEY"
	HTTPSSLVerifyEnvName = "LAUNCHBOX_HTTP_SSL_VERIFY"
)

type Client struct {
	config Config
}

type Config struct {
	Address    string
	Scheme     string
	Transport  *http.Transport
	HttpClient *http.Client
	TLSConfig  TLSConfig
}

type TLSConfig struct {
	Address  string
	CAFile   string
	CAPath   string
	CertFile string
	KeyFile  string
	Insecure bool
}

type PaginationOptions struct {
	Page    uint
	PerPage uint
}

func defaultConfig(transportFn func() *http.Transport) *Config {
	config := &Config{
		Address:   "http://127.0.0.1:8080",
		Transport: transportFn(),
	}
	if addr := os.Getenv(HTTPAddrEnvName); addr != "" {
		config.Address = addr
	}
	if v := os.Getenv(HTTPTLSServerName); v != "" {
		config.TLSConfig.Address = v
	}
	if v := os.Getenv(HTTPCAFile); v != "" {
		config.TLSConfig.CAFile = v
	}
	if v := os.Getenv(HTTPCAPath); v != "" {
		config.TLSConfig.CAPath = v
	}
	if v := os.Getenv(HTTPClientCert); v != "" {
		config.TLSConfig.CertFile = v
	}
	if v := os.Getenv(HTTPClientKey); v != "" {
		config.TLSConfig.KeyFile = v
	}
	if v := os.Getenv(HTTPSSLVerifyEnvName); v != "" {
		verify, err := strconv.ParseBool(v)
		if err == nil {
			if !verify {
				config.TLSConfig.Insecure = true
			}
		}

	}
	return config
}

func DefaultConfig() *Config {
	return defaultConfig(cleanhttp.DefaultPooledTransport)
}

func New() (*Client, error) {
	conf := DefaultConfig()
	client := &http.Client{
		Transport: conf.Transport,
	}
	conf.HttpClient = client
	return &Client{config: *conf}, nil
}

func (c *Client) get(path string, query map[string]string, out interface{}) error {
	url := strings.Join([]string{c.config.Address, path}, "")
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	q := req.URL.Query()
	if query != nil {
		for key, value := range query {
			q.Add(key, value)
		}
	}

	// assign encoded query string to http request
	req.URL.RawQuery = q.Encode()

	resp, err := c.config.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return decodeResponse(resp.Body, out)
}

func (c *Client) post(path string, in interface{}) error {
	url := strings.Join([]string{c.config.Address, path}, "")
	input, err := json.Marshal(in)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(input))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return decodeResponse(resp.Body, in)
}

func (c *Client) put(path string, in interface{}) error {
	url := strings.Join([]string{c.config.Address, path}, "")
	input, err := json.Marshal(in)
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(input))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.config.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return decodeResponse(resp.Body, in)
}

func (c *Client) delete(path string) error {
	url := strings.Join([]string{c.config.Address, path}, "")
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}
	_, err = c.config.HttpClient.Do(req)
	if err != nil {
		return err
	}
	return nil
}

func decodeResponse(resBody io.ReadCloser, out interface{}) error {
	body, err := io.ReadAll(resBody)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, out)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	return nil
}
