package cloopen

import (
	"net"
	"net/http"
	"time"
)

type Config struct {
	UseSSL     bool
	HttpClient *http.Client
	HttpConf   *HttpConf

	SmsHost     string
	APIAccount  string
	APIToken    string
	ContentType string
}

type HttpConf struct {
	Timeout             time.Duration
	KeepAlive           time.Duration
	MaxIdleConns        int
	IdleConnTimeout     time.Duration
	TLSHandshakeTimeout time.Duration
}

func DefaultConfig() *Config {
	cfg := &Config{
		SmsHost: "app.cloopen.com:8883",
		UseSSL:  true,
		HttpConf: &HttpConf{
			Timeout:             300,
			KeepAlive:           30000,
			MaxIdleConns:        100,
			IdleConnTimeout:     30000,
			TLSHandshakeTimeout: 300,
		},
	}
	return cfg.WithHTTPClient(defaultHTTPClient(cfg))
}

func defaultHTTPClient(conf *Config) *http.Client {
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   conf.HttpConf.Timeout * time.Millisecond,
			KeepAlive: conf.HttpConf.KeepAlive * time.Millisecond,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:        conf.HttpConf.MaxIdleConns,
		IdleConnTimeout:     conf.HttpConf.IdleConnTimeout * time.Millisecond,
		TLSHandshakeTimeout: conf.HttpConf.TLSHandshakeTimeout * time.Millisecond,
	}
	return &http.Client{Transport: transport}
}

func (c *Config) WithHTTPClient(client *http.Client) *Config {
	if client != nil {
		c.HttpClient = client
	}
	return c
}

func (c *Config) WithUseSSL(use bool) *Config {
	c.UseSSL = use
	return c
}

func (c *Config) WithAPIAccount(account string) *Config {
	c.APIAccount = account
	return c
}

func (c *Config) WithAPIToken(token string) *Config {
	c.APIToken = token
	return c
}

func (c *Config) WithHttpConf(conf *HttpConf) *Config {
	if conf != nil {
		c.HttpConf = conf
	}
	return c.WithHTTPClient(defaultHTTPClient(c))
}

func (c *Config) WithSmsHost(smsHost string) *Config {
	c.SmsHost = smsHost
	return c
}

func (c *Config) withContentType(contentType string) *Config {
	c.ContentType = contentType
	return c
}
